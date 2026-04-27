package main

import (
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"embed"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
)

//go:embed assets/locale.js
var assetFS embed.FS

type ui8kitPackageConfig struct {
	UI8Kit struct {
		Aria struct {
			Mode     string   `json:"mode"`
			Patterns []string `json:"patterns"`
			AutoInit *bool    `json:"autoInit"`
		} `json:"aria"`
	} `json:"ui8kit"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}

type syncConfig struct {
	StaticRoot   string
	EnableCSS    bool
	EnableFonts  bool
	EnableTheme  bool
	EnableAria   bool
	EnableLocale bool
	EnableUI8px  bool
	HashNames    bool
	AriaMode     string
	AriaPatterns []string
	AriaVersion  string
	AutoInit     bool
	ManifestPath string
	Report       bool
}

type manifestEntry struct {
	File string `json:"file"`
	SRI  string `json:"sri"`
}

type syncManifest struct {
	Theme           *manifestEntry `json:"theme,omitempty"`
	UI8Kit          *manifestEntry `json:"ui8kit,omitempty"`
	AriaMode        string         `json:"ariaMode,omitempty"`
	AriaVersion     string         `json:"ariaVersion,omitempty"`
	AriaUpstreamSRI string         `json:"ariaUpstreamSRI,omitempty"`
	Patterns        []string       `json:"patterns,omitempty"`
}

func main() {
	cfg, err := parseFlags(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	if err := run(cfg); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func parseFlags(args []string) (syncConfig, error) {
	var cfg syncConfig
	fs := flag.NewFlagSet("sync-assets", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	fs.BoolVar(&cfg.EnableCSS, "css", true, "copy UI8Kit CSS assets")
	fs.BoolVar(&cfg.EnableFonts, "fonts", true, "copy font assets")
	fs.BoolVar(&cfg.EnableTheme, "theme", true, "emit theme.js")
	fs.BoolVar(&cfg.EnableAria, "aria", true, "emit ui8kit.js with aria bundle")
	fs.BoolVar(&cfg.EnableLocale, "locale", true, "include locale.js in ui8kit.js")
	fs.BoolVar(&cfg.EnableUI8px, "ui8px-policy", true, "write missing app-local ui8px policy files")
	fs.BoolVar(&cfg.HashNames, "hash", true, "emit hashed JS filenames in addition to stable aliases")
	fs.StringVar(&cfg.AriaMode, "aria-mode", "subset", "aria mode: subset or full")
	fs.StringVar(&cfg.AriaVersion, "aria-version", "", "override aria version")
	fs.StringVar(&cfg.ManifestPath, "manifest", "", "manifest path (defaults to <static_root>/js/manifest.json)")
	fs.BoolVar(&cfg.Report, "report", true, "print size report")
	patterns := fs.String("aria-patterns", "", "comma-separated aria patterns")
	if err := fs.Parse(args); err != nil {
		return cfg, err
	}
	if fs.NArg() != 1 {
		return cfg, errors.New("usage: sync-assets [flags] <static_root>")
	}
	cfg.StaticRoot = fs.Arg(0)
	if *patterns != "" {
		cfg.AriaPatterns = splitCSV(*patterns)
	}
	cfg.AriaMode = strings.ToLower(strings.TrimSpace(cfg.AriaMode))
	if cfg.AriaMode != "subset" && cfg.AriaMode != "full" {
		return cfg, fmt.Errorf("unsupported --aria-mode %q", cfg.AriaMode)
	}
	return cfg, nil
}

func run(cfg syncConfig) error {
	startDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("resolve working directory: %w", err)
	}
	staticRoot, err := filepath.Abs(cfg.StaticRoot)
	if err != nil {
		return fmt.Errorf("resolve static root: %w", err)
	}
	cfg.StaticRoot = staticRoot
	if cfg.ManifestPath == "" {
		cfg.ManifestPath = filepath.Join(staticRoot, "js", "manifest.json")
	}

	localPkgDir, localPkg, err := findPackageConfig(startDir)
	if err != nil {
		return err
	}
	cfg = applyPackageConfig(cfg, localPkg)

	ui8kitDir, err := resolveModuleDir("github.com/fastygo/ui8kit")
	if err != nil {
		return err
	}
	frameworkDir, err := resolveModuleDir("github.com/fastygo/framework")
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Join(staticRoot, "css", "ui8kit"), 0o755); err != nil {
		return fmt.Errorf("create css target: %w", err)
	}
	if err := os.MkdirAll(filepath.Join(staticRoot, "js"), 0o755); err != nil {
		return fmt.Errorf("create js target: %w", err)
	}

	if cfg.EnableCSS {
		if err := syncCSS(staticRoot, ui8kitDir); err != nil {
			return err
		}
		if cfg.EnableUI8px {
			if err := syncUI8pxPolicy(localPkgDir, ui8kitDir); err != nil {
				return err
			}
		}
	}
	if cfg.EnableFonts {
		if err := syncFonts(staticRoot, frameworkDir); err != nil {
			return err
		}
	}

	manifest := syncManifest{
		AriaMode:    cfg.AriaMode,
		AriaVersion: cfg.AriaVersion,
		Patterns:    slices.Clone(cfg.AriaPatterns),
	}
	if err := cleanupJSOutputs(filepath.Join(staticRoot, "js")); err != nil {
		return err
	}

	themePath := filepath.Join(ui8kitDir, "js", "theme.js")
	if cfg.EnableTheme {
		themeBytes, err := os.ReadFile(themePath)
		if err != nil {
			return fmt.Errorf("read theme.js: %w", err)
		}
		entry, err := writeJSArtifact(filepath.Join(staticRoot, "js"), "theme", themeBytes, cfg.HashNames, true)
		if err != nil {
			return err
		}
		manifest.Theme = &entry
	}

	if cfg.EnableAria {
		ariaDir, err := resolveAriaDir(localPkgDir)
		if err != nil {
			return err
		}
		if cfg.AriaVersion == "" {
			cfg.AriaVersion = dependencyVersion(localPkg, "@ui8kit/aria")
		}

		tempDir := filepath.Join(staticRoot, ".ui8kit-build")
		if err := os.RemoveAll(tempDir); err != nil {
			return fmt.Errorf("clean temp build dir: %w", err)
		}
		if err := os.MkdirAll(tempDir, 0o755); err != nil {
			return fmt.Errorf("create temp build dir: %w", err)
		}
		defer os.RemoveAll(tempDir)

		ariaBytes, upstreamSRI, err := buildAriaBundle(tempDir, ariaDir, cfg)
		if err != nil {
			return err
		}
		manifest.AriaUpstreamSRI = upstreamSRI

		bundleParts := [][]byte{ariaBytes}
		if cfg.EnableLocale {
			localeBytes, err := assetFS.ReadFile("assets/locale.js")
			if err != nil {
				return fmt.Errorf("read embedded locale.js: %w", err)
			}
			bundleParts = append(bundleParts, localeBytes)
		}
		ui8kitBytes := bytes.Join(bundleParts, []byte(";\n"))
		entry, err := writeJSArtifact(filepath.Join(staticRoot, "js"), "ui8kit", ui8kitBytes, cfg.HashNames, true)
		if err != nil {
			return err
		}
		manifest.UI8Kit = &entry
		if cfg.Report {
			fmt.Printf("ui8kit bundle: %d bytes (%s, patterns: %s)\n", len(ui8kitBytes), cfg.AriaMode, strings.Join(cfg.AriaPatterns, ", "))
		}
	}

	if err := writeManifest(cfg.ManifestPath, manifest); err != nil {
		return err
	}

	fmt.Printf("ui8kit assets synced into %s\n", staticRoot)
	fmt.Printf(" - css:   %s\n", filepath.Join(staticRoot, "css", "ui8kit"))
	if cfg.EnableFonts {
		fmt.Printf(" - fonts: %s\n", filepath.Join(staticRoot, "fonts", "outfit"))
	}
	if cfg.EnableTheme && manifest.Theme != nil {
		fmt.Printf(" - theme: %s\n", filepath.Join(staticRoot, "js", manifest.Theme.File))
	}
	if cfg.EnableAria && manifest.UI8Kit != nil {
		fmt.Printf(" - js:    %s\n", filepath.Join(staticRoot, "js", manifest.UI8Kit.File))
	}
	fmt.Printf(" - manifest: %s\n", cfg.ManifestPath)
	return nil
}

func applyPackageConfig(cfg syncConfig, pkg ui8kitPackageConfig) syncConfig {
	if pkg.UI8Kit.Aria.Mode != "" {
		cfg.AriaMode = strings.ToLower(pkg.UI8Kit.Aria.Mode)
	}
	if len(cfg.AriaPatterns) == 0 {
		cfg.AriaPatterns = slices.Clone(pkg.UI8Kit.Aria.Patterns)
	}
	if cfg.AriaVersion == "" {
		cfg.AriaVersion = dependencyVersion(pkg, "@ui8kit/aria")
	}
	if pkg.UI8Kit.Aria.AutoInit != nil {
		cfg.AutoInit = *pkg.UI8Kit.Aria.AutoInit
	} else {
		cfg.AutoInit = true
	}
	if len(cfg.AriaPatterns) == 0 {
		cfg.AriaPatterns = []string{"dialog"}
	}
	slices.Sort(cfg.AriaPatterns)
	cfg.AriaPatterns = slices.Compact(cfg.AriaPatterns)
	return cfg
}

func findPackageConfig(startDir string) (string, ui8kitPackageConfig, error) {
	var cfg ui8kitPackageConfig
	dir := startDir
	for {
		path := filepath.Join(dir, "package.json")
		data, err := os.ReadFile(path)
		if err == nil {
			if err := json.Unmarshal(data, &cfg); err != nil {
				return "", cfg, fmt.Errorf("parse %s: %w", path, err)
			}
			return dir, cfg, nil
		}
		if !errors.Is(err, os.ErrNotExist) {
			return "", cfg, fmt.Errorf("read %s: %w", path, err)
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "", cfg, fmt.Errorf("package.json not found from %s upward", startDir)
}

func dependencyVersion(pkg ui8kitPackageConfig, name string) string {
	if pkg.DevDependencies != nil {
		if version, ok := pkg.DevDependencies[name]; ok {
			return version
		}
	}
	if pkg.Dependencies != nil {
		if version, ok := pkg.Dependencies[name]; ok {
			return version
		}
	}
	return ""
}

func resolveModuleDir(module string) (string, error) {
	if module == "github.com/fastygo/ui8kit" {
		if env := strings.TrimSpace(os.Getenv("UI8KIT_SYNC_UI8KIT_DIR")); env != "" {
			return env, nil
		}
	}
	if module == "github.com/fastygo/framework" {
		if env := strings.TrimSpace(os.Getenv("UI8KIT_SYNC_FRAMEWORK_DIR")); env != "" {
			return env, nil
		}
	}
	dir, err := runGoModuleDir(module)
	if err == nil && dir != "" {
		return dir, nil
	}
	return "", fmt.Errorf("resolve %s: %w", module, err)
}

func runGoModuleDir(module string) (string, error) {
	cmd := exec.Command("go", "list", "-m", "-f", "{{.Dir}}", module)
	out, err := cmd.CombinedOutput()
	if err == nil {
		dir := strings.TrimSpace(string(out))
		if dir != "" {
			return dir, nil
		}
	}
	cmd = exec.Command("go", "mod", "download", "-json", module)
	out, err = cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s", strings.TrimSpace(string(out)))
	}
	type downloadInfo struct {
		Dir string `json:"Dir"`
	}
	var info downloadInfo
	if err := json.Unmarshal(out, &info); err != nil {
		return "", err
	}
	if info.Dir == "" {
		return "", errors.New("module dir missing from go mod download output")
	}
	return info.Dir, nil
}

func syncCSS(staticRoot, ui8kitDir string) error {
	names := []string{"base.css", "shell.css", "components.css", "latty.css", "prose.css"}
	for _, name := range names {
		src := filepath.Join(ui8kitDir, "styles", name)
		dst := filepath.Join(staticRoot, "css", "ui8kit", name)
		if err := copyFile(src, dst); err != nil {
			return fmt.Errorf("copy %s: %w", name, err)
		}
	}
	return nil
}

func syncUI8pxPolicy(appDir, ui8kitDir string) error {
	policyDir := filepath.Join(appDir, ".ui8px", "policy")
	if err := os.MkdirAll(policyDir, 0o755); err != nil {
		return fmt.Errorf("create ui8px policy dir: %w", err)
	}

	if err := syncUI8pxAllowed(filepath.Join(ui8kitDir, ".ui8px", "policy", "allowed.json"), filepath.Join(policyDir, "allowed.json")); err != nil {
		return err
	}
	for _, name := range []string{"denied.json", "groups.json"} {
		dst := filepath.Join(policyDir, name)
		if _, err := os.Stat(dst); err == nil {
			continue
		} else if !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("stat %s: %w", dst, err)
		}

		src := filepath.Join(ui8kitDir, ".ui8px", "policy", name)
		if err := copyFile(src, dst); err != nil {
			return fmt.Errorf("copy ui8px %s: %w", name, err)
		}
	}
	if err := syncUI8pxPatterns(filepath.Join(ui8kitDir, ".ui8px", "policy", "patterns.json"), filepath.Join(policyDir, "patterns.json")); err != nil {
		return err
	}

	scopesPath := filepath.Join(policyDir, "scopes.json")
	if _, err := os.Stat(scopesPath); err == nil {
		return nil
	} else if !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("stat %s: %w", scopesPath, err)
	}

	if err := os.WriteFile(scopesPath, []byte(appUI8pxScopesJSON), 0o644); err != nil {
		return fmt.Errorf("write ui8px scopes: %w", err)
	}
	return nil
}

type ui8pxAllowedPolicy struct {
	Spacing   map[string][]string `json:"spacing"`
	Utilities []string            `json:"utilities"`
}

func syncUI8pxAllowed(src, dst string) error {
	srcPolicy, err := readUI8pxAllowed(src)
	if err != nil {
		return fmt.Errorf("read ui8px allowed source: %w", err)
	}
	if _, err := os.Stat(dst); errors.Is(err, os.ErrNotExist) {
		if err := copyFile(src, dst); err != nil {
			return fmt.Errorf("copy ui8px allowed.json: %w", err)
		}
		return nil
	} else if err != nil {
		return fmt.Errorf("stat %s: %w", dst, err)
	}

	dstPolicy, err := readUI8pxAllowed(dst)
	if err != nil {
		return fmt.Errorf("read ui8px allowed target: %w", err)
	}
	dstPolicy.Utilities = mergeStringLists(dstPolicy.Utilities, srcPolicy.Utilities)
	if dstPolicy.Spacing == nil {
		dstPolicy.Spacing = srcPolicy.Spacing
	}

	data, err := json.MarshalIndent(dstPolicy, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal ui8px allowed: %w", err)
	}
	data = append(data, '\n')
	if err := os.WriteFile(dst, data, 0o644); err != nil {
		return fmt.Errorf("write ui8px allowed: %w", err)
	}
	return nil
}

func readUI8pxAllowed(path string) (ui8pxAllowedPolicy, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return ui8pxAllowedPolicy{}, err
	}
	var policy ui8pxAllowedPolicy
	if err := json.Unmarshal(data, &policy); err != nil {
		return ui8pxAllowedPolicy{}, err
	}
	return policy, nil
}

type ui8pxPatternsPolicy struct {
	Patterns map[string][]string `json:"patterns"`
}

func syncUI8pxPatterns(src, dst string) error {
	srcPolicy, err := readUI8pxPatterns(src)
	if err != nil {
		return fmt.Errorf("read ui8px patterns source: %w", err)
	}
	dstPolicy := ui8pxPatternsPolicy{Patterns: map[string][]string{}}
	if _, err := os.Stat(dst); err == nil {
		dstPolicy, err = readUI8pxPatterns(dst)
		if err != nil {
			return fmt.Errorf("read ui8px patterns target: %w", err)
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("stat %s: %w", dst, err)
	}
	if dstPolicy.Patterns == nil {
		dstPolicy.Patterns = map[string][]string{}
	}

	for name, tokens := range srcPolicy.Patterns {
		if strings.HasPrefix(name, "ui-") {
			dstPolicy.Patterns[name] = slices.Clone(tokens)
		} else if _, ok := dstPolicy.Patterns[name]; !ok {
			dstPolicy.Patterns[name] = slices.Clone(tokens)
		}
	}

	data, err := json.MarshalIndent(dstPolicy, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal ui8px patterns: %w", err)
	}
	data = append(data, '\n')
	if err := os.WriteFile(dst, data, 0o644); err != nil {
		return fmt.Errorf("write ui8px patterns: %w", err)
	}
	return nil
}

func readUI8pxPatterns(path string) (ui8pxPatternsPolicy, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return ui8pxPatternsPolicy{}, err
	}
	var policy ui8pxPatternsPolicy
	if err := json.Unmarshal(data, &policy); err != nil {
		return ui8pxPatternsPolicy{}, err
	}
	return policy, nil
}

func mergeStringLists(existing, incoming []string) []string {
	seen := make(map[string]bool, len(existing)+len(incoming))
	merged := make([]string, 0, len(existing)+len(incoming))
	for _, value := range existing {
		if value == "" || seen[value] {
			continue
		}
		seen[value] = true
		merged = append(merged, value)
	}
	for _, value := range incoming {
		if value == "" || seen[value] {
			continue
		}
		seen[value] = true
		merged = append(merged, value)
	}
	slices.Sort(merged)
	return merged
}

const appUI8pxScopesJSON = `{
  "defaultScope": "layout",
  "scopes": [
    {
      "name": "views",
      "files": [
        "**/internal/site/views/**"
      ],
      "spacing": "layout"
    },
    {
      "name": "app-components",
      "files": [
        "web/static/css/*-components.css",
        "**/web/static/css/*-components.css"
      ],
      "spacing": "control"
    },
    {
      "name": "app-theme",
      "files": [
        "web/static/css/shadcn.css",
        "web/static/css/theme.css",
        "web/static/css/tokens.css",
        "**/web/static/css/shadcn.css",
        "**/web/static/css/theme.css",
        "**/web/static/css/tokens.css"
      ],
      "spacing": "control"
    },
    {
      "name": "ui8kit-assets",
      "files": [
        "web/static/css/ui8kit/**",
        "**/web/static/css/ui8kit/**"
      ],
      "spacing": "control"
    }
  ]
}
`

func syncFonts(staticRoot, frameworkDir string) error {
	srcCSS := filepath.Join(frameworkDir, "pkg", "fonts", "outfit.css")
	if err := copyFile(srcCSS, filepath.Join(staticRoot, "css", "fonts.css")); err != nil {
		return fmt.Errorf("copy fonts.css: %w", err)
	}
	srcFonts := filepath.Join(frameworkDir, "pkg", "fonts", "outfit")
	dstFonts := filepath.Join(staticRoot, "fonts", "outfit")
	if err := os.RemoveAll(dstFonts); err != nil {
		return fmt.Errorf("remove old fonts: %w", err)
	}
	if err := copyDir(srcFonts, dstFonts); err != nil {
		return fmt.Errorf("copy font directory: %w", err)
	}
	return nil
}

func buildAriaBundle(tempDir, ariaDir string, cfg syncConfig) ([]byte, string, error) {
	if cfg.AriaMode == "full" {
		src := filepath.Join(ariaDir, "dist", "all.iife.min.js")
		data, err := os.ReadFile(src)
		if err != nil {
			return nil, "", fmt.Errorf("read full aria bundle: %w", err)
		}
		upstreamSRI := ""
		if sriData, err := os.ReadFile(filepath.Join(ariaDir, "dist", "sri.json")); err == nil {
			var sriMap map[string]string
			if json.Unmarshal(sriData, &sriMap) == nil {
				upstreamSRI = sriMap["all.iife.min.js"]
			}
		}
		return data, upstreamSRI, nil
	}

	entryPath := filepath.Join(tempDir, "aria.entry.mjs")
	outPath := filepath.Join(tempDir, "aria.iife.js")
	if err := os.WriteFile(entryPath, []byte(buildAriaEntry(cfg.AriaPatterns, cfg.AutoInit)), 0o644); err != nil {
		return nil, "", fmt.Errorf("write aria entry: %w", err)
	}

	bunBinary := strings.TrimSpace(os.Getenv("UI8KIT_SYNC_BUN"))
	if bunBinary == "" {
		bunBinary = "bun"
	}
	cmd := exec.Command(bunBinary, "build", entryPath, "--target=browser", "--format=iife", "--minify", "--outfile", outPath)
	cmd.Dir = tempDir
	cmd.Env = append(os.Environ(), "NODE_PATH="+filepath.Join(filepath.Dir(ariaDir), ".."))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, "", fmt.Errorf("bun build subset aria bundle: %w\n%s", err, strings.TrimSpace(string(output)))
	}
	data, err := os.ReadFile(outPath)
	if err != nil {
		return nil, "", fmt.Errorf("read built aria bundle: %w", err)
	}
	return data, "", nil
}

func buildAriaEntry(patterns []string, autoInit bool) string {
	imports := []string{"getNamespace", "registerPattern"}
	registerLines := make([]string, 0, len(patterns))
	for _, pattern := range patterns {
		imports = append(imports, pattern)
		registerLines = append(registerLines, fmt.Sprintf("registerPattern(%s)", pattern))
	}
	autoInitLine := ""
	if autoInit {
		autoInitLine = `if (typeof window !== "undefined" && window.__UI8KIT_ARIA_AUTO_INIT__ !== false) {` + "\n  getNamespace().init()\n}"
	}
	return fmt.Sprintf(`import { %s } from "@ui8kit/aria"

%s
%s
`, strings.Join(imports, ", "), strings.Join(registerLines, "\n"), autoInitLine)
}

func resolveAriaDir(startDir string) (string, error) {
	if env := strings.TrimSpace(os.Getenv("UI8KIT_SYNC_ARIA_DIR")); env != "" {
		return env, nil
	}
	dir := startDir
	for {
		candidate := filepath.Join(dir, "node_modules", "@ui8kit", "aria")
		if stat, err := os.Stat(candidate); err == nil && stat.IsDir() {
			return candidate, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "", fmt.Errorf("@ui8kit/aria not found from %s upward; run bun install first", startDir)
}

func cleanupJSOutputs(jsDir string) error {
	entries, err := os.ReadDir(jsDir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return os.MkdirAll(jsDir, 0o755)
		}
		return fmt.Errorf("read %s: %w", jsDir, err)
	}
	for _, entry := range entries {
		name := entry.Name()
		if name == "manifest.json" || strings.HasPrefix(name, "theme") || strings.HasPrefix(name, "ui8kit") || strings.HasPrefix(name, "aria.iife") {
			if err := os.RemoveAll(filepath.Join(jsDir, name)); err != nil {
				return fmt.Errorf("remove %s: %w", name, err)
			}
		}
	}
	return nil
}

func writeManifest(path string, manifest syncManifest) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("create manifest dir: %w", err)
	}
	data, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal manifest: %w", err)
	}
	data = append(data, '\n')
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("write manifest: %w", err)
	}
	return nil
}

func writeJSArtifact(targetDir, baseName string, data []byte, hashed, stableAlias bool) (manifestEntry, error) {
	if err := os.MkdirAll(targetDir, 0o755); err != nil {
		return manifestEntry{}, fmt.Errorf("create js target dir: %w", err)
	}
	hash := shortHash(data)
	fileName := baseName + ".js"
	if hashed {
		fileName = fmt.Sprintf("%s.%s.js", baseName, hash)
	}
	fullPath := filepath.Join(targetDir, fileName)
	if err := os.WriteFile(fullPath, data, 0o644); err != nil {
		return manifestEntry{}, fmt.Errorf("write %s: %w", fileName, err)
	}
	if stableAlias {
		if err := os.WriteFile(filepath.Join(targetDir, baseName+".js"), data, 0o644); err != nil {
			return manifestEntry{}, fmt.Errorf("write stable alias for %s: %w", baseName, err)
		}
	}
	return manifestEntry{File: fileName, SRI: sriSHA384(data)}, nil
}

func shortHash(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:4])
}

func sriSHA384(data []byte) string {
	sum := sha512.Sum384(data)
	return "sha384-" + base64.StdEncoding.EncodeToString(sum[:])
}

func splitCSV(value string) []string {
	if value == "" {
		return nil
	}
	parts := strings.Split(value, ",")
	out := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			out = append(out, part)
		}
	}
	return out
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return err
	}
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	if _, err := io.Copy(out, in); err != nil {
		return err
	}
	return out.Close()
}

func copyDir(src, dst string) error {
	return filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)
		if d.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		return copyFile(path, target)
	})
}

func init() {
	// Ensure Windows path separators do not leak into the emitted JS.
	if runtime.GOOS == "windows" {
		os.Setenv("BUN_INSTALL_CACHE_DIR", filepath.Join(os.TempDir(), "bun-install-cache"))
	}
}
