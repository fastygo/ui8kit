package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestFullModeCopiesAssetsAndWritesManifest(t *testing.T) {
	appDir := t.TempDir()
	ui8kitDir := filepath.Join(t.TempDir(), "ui8kit")
	frameworkDir := filepath.Join(t.TempDir(), "framework")
	ariaDir := filepath.Join(appDir, "node_modules", "@ui8kit", "aria")

	writeFile(t, filepath.Join(appDir, "package.json"), `{
  "name": "example",
  "devDependencies": { "@ui8kit/aria": "0.1.0" },
  "ui8kit": { "aria": { "mode": "full", "patterns": ["dialog"] } }
}`)
	writeFile(t, filepath.Join(ui8kitDir, "styles", "base.css"), "base")
	writeFile(t, filepath.Join(ui8kitDir, "styles", "shell.css"), "shell")
	writeFile(t, filepath.Join(ui8kitDir, "styles", "components.css"), "components")
	writeFile(t, filepath.Join(ui8kitDir, "styles", "latty.css"), "latty")
	writeFile(t, filepath.Join(ui8kitDir, "styles", "prose.css"), "prose")
	writeFile(t, filepath.Join(ui8kitDir, "js", "theme.js"), "(function(){window.themeLoaded=true})();")
	writeFile(t, filepath.Join(frameworkDir, "pkg", "fonts", "outfit.css"), "@font-face{}")
	writeFile(t, filepath.Join(frameworkDir, "pkg", "fonts", "outfit", "latin", "font.woff2"), "font")
	writeFile(t, filepath.Join(ariaDir, "dist", "all.iife.min.js"), "(function(){window.ariaLoaded=true})();")
	writeFile(t, filepath.Join(ariaDir, "dist", "sri.json"), `{"all.iife.min.js":"sha384-upstream"}`)

	setEnv(t, "UI8KIT_SYNC_UI8KIT_DIR", ui8kitDir)
	setEnv(t, "UI8KIT_SYNC_FRAMEWORK_DIR", frameworkDir)

	oldwd, _ := os.Getwd()
	t.Cleanup(func() { _ = os.Chdir(oldwd) })
	if err := os.Chdir(appDir); err != nil {
		t.Fatalf("chdir: %v", err)
	}

	cfg := syncConfig{
		StaticRoot:   "web/static",
		EnableCSS:    true,
		EnableFonts:  true,
		EnableTheme:  true,
		EnableAria:   true,
		EnableLocale: true,
		HashNames:    true,
		AriaMode:     "full",
		Report:       false,
	}
	if err := run(cfg); err != nil {
		t.Fatalf("run() error = %v", err)
	}

	jsDir := filepath.Join(appDir, "web", "static", "js")
	assertExists(t, filepath.Join(jsDir, "theme.js"))
	assertExists(t, filepath.Join(jsDir, "ui8kit.js"))
	manifestPath := filepath.Join(jsDir, "manifest.json")
	assertExists(t, manifestPath)

	data, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Fatalf("read manifest: %v", err)
	}
	var manifest syncManifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		t.Fatalf("unmarshal manifest: %v", err)
	}
	if manifest.Theme == nil || manifest.UI8Kit == nil {
		t.Fatalf("manifest entries missing: %+v", manifest)
	}
	if !strings.HasPrefix(manifest.Theme.File, "theme.") {
		t.Fatalf("unexpected theme file: %q", manifest.Theme.File)
	}
	if !strings.HasPrefix(manifest.UI8Kit.File, "ui8kit.") {
		t.Fatalf("unexpected ui8kit file: %q", manifest.UI8Kit.File)
	}
	if manifest.AriaUpstreamSRI != "sha384-upstream" {
		t.Fatalf("unexpected upstream sri: %q", manifest.AriaUpstreamSRI)
	}
}

func TestBuildAriaEntryUsesSelectedPatterns(t *testing.T) {
	entry := buildAriaEntry([]string{"dialog", "tabs"}, true)
	if !strings.Contains(entry, "registerPattern(dialog)") {
		t.Fatalf("missing dialog registration:\n%s", entry)
	}
	if !strings.Contains(entry, "registerPattern(tabs)") {
		t.Fatalf("missing tabs registration:\n%s", entry)
	}
	if !strings.Contains(entry, `window.__UI8KIT_ARIA_AUTO_INIT__ !== false`) {
		t.Fatalf("missing auto-init guard:\n%s", entry)
	}
}

func TestSubsetModeGeneratesBundleWithBun(t *testing.T) {
	if _, err := exec.LookPath("bun"); err != nil {
		t.Skip("bun not installed")
	}

	appDir := t.TempDir()
	ui8kitDir := filepath.Join(t.TempDir(), "ui8kit")
	frameworkDir := filepath.Join(t.TempDir(), "framework")
	ariaPkg := filepath.Join(appDir, "node_modules", "@ui8kit", "aria")

	writeFile(t, filepath.Join(appDir, "package.json"), `{
  "name": "example",
  "devDependencies": { "@ui8kit/aria": "0.1.0" },
  "ui8kit": { "aria": { "mode": "subset", "patterns": ["dialog"] } }
}`)
	writeFile(t, filepath.Join(ui8kitDir, "styles", "base.css"), "base")
	writeFile(t, filepath.Join(ui8kitDir, "styles", "shell.css"), "shell")
	writeFile(t, filepath.Join(ui8kitDir, "styles", "components.css"), "components")
	writeFile(t, filepath.Join(ui8kitDir, "styles", "latty.css"), "latty")
	writeFile(t, filepath.Join(ui8kitDir, "styles", "prose.css"), "prose")
	writeFile(t, filepath.Join(ui8kitDir, "js", "theme.js"), "(function(){window.themeLoaded=true})();")
	writeFile(t, filepath.Join(frameworkDir, "pkg", "fonts", "outfit.css"), "@font-face{}")
	writeFile(t, filepath.Join(frameworkDir, "pkg", "fonts", "outfit", "latin", "font.woff2"), "font")
	writeFile(t, filepath.Join(ariaPkg, "package.json"), `{
  "name": "@ui8kit/aria",
  "type": "module",
  "main": "./index.js"
}`)
	writeFile(t, filepath.Join(ariaPkg, "index.js"), `export function registerPattern() {}
export function getNamespace() { return { init() {} } }
export const dialog = { name: "dialog" }
`)

	setEnv(t, "UI8KIT_SYNC_UI8KIT_DIR", ui8kitDir)
	setEnv(t, "UI8KIT_SYNC_FRAMEWORK_DIR", frameworkDir)

	oldwd, _ := os.Getwd()
	t.Cleanup(func() { _ = os.Chdir(oldwd) })
	if err := os.Chdir(appDir); err != nil {
		t.Fatalf("chdir: %v", err)
	}

	cfg := syncConfig{
		StaticRoot:   "web/static",
		EnableCSS:    false,
		EnableFonts:  false,
		EnableTheme:  true,
		EnableAria:   true,
		EnableLocale: true,
		HashNames:    true,
		AriaMode:     "subset",
		Report:       false,
	}
	if err := run(cfg); err != nil {
		t.Fatalf("run() error = %v", err)
	}
	assertExists(t, filepath.Join(appDir, "web", "static", "js", "ui8kit.js"))
}

func setEnv(t *testing.T, key, value string) {
	t.Helper()
	old, had := os.LookupEnv(key)
	if err := os.Setenv(key, value); err != nil {
		t.Fatalf("setenv %s: %v", key, err)
	}
	t.Cleanup(func() {
		if had {
			_ = os.Setenv(key, old)
		} else {
			_ = os.Unsetenv(key)
		}
	})
}

func writeFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir %s: %v", filepath.Dir(path), err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}

func assertExists(t *testing.T, path string) {
	t.Helper()
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("expected %s to exist: %v", path, err)
	}
}
