package styleaudit

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type Config struct {
	CSSPath      string
	Root         string
	Extensions   []string
	TemplateDirs []string
}

type Result struct {
	Classes     []string
	Unused      []string
	SourceFiles []string
}

type UnusedClassesError struct {
	Count int
}

func (e UnusedClassesError) Error() string {
	return fmt.Sprintf("%d unused CSS classes found", e.Count)
}

func Run(cfg Config) (Result, error) {
	cssPath, err := filepath.Abs(cfg.CSSPath)
	if err != nil {
		return Result{}, fmt.Errorf("resolve css path: %w", err)
	}
	root, err := filepath.Abs(cfg.Root)
	if err != nil {
		return Result{}, fmt.Errorf("resolve root path: %w", err)
	}

	css, err := os.ReadFile(cssPath)
	if err != nil {
		return Result{}, fmt.Errorf("read css: %w", err)
	}
	classes := ExtractSelectorClasses(string(css))
	if len(classes) == 0 {
		return Result{}, fmt.Errorf("no CSS selector classes found in %s", cssPath)
	}

	sourceFiles, err := CollectSourceFiles(root, cfg.TemplateDirs, cfg.Extensions)
	if err != nil {
		return Result{}, err
	}
	if len(sourceFiles) == 0 {
		return Result{}, fmt.Errorf("no source files found under %s for extensions %s", root, strings.Join(cfg.Extensions, ", "))
	}

	usage, err := FindClassUsage(classes, sourceFiles)
	if err != nil {
		return Result{}, err
	}

	result := Result{
		Classes:     classes,
		SourceFiles: sourceFiles,
	}
	for _, className := range classes {
		if !usage[className] {
			result.Unused = append(result.Unused, className)
		}
	}
	if len(result.Unused) > 0 {
		return result, UnusedClassesError{Count: len(result.Unused)}
	}
	return result, nil
}

func CleanFile(cfg Config) (Result, int, error) {
	result, err := Run(cfg)
	var auditErr UnusedClassesError
	if err != nil && !errors.As(err, &auditErr) {
		return result, 0, err
	}
	if len(result.Unused) == 0 {
		return result, 0, nil
	}

	css, err := os.ReadFile(cfg.CSSPath)
	if err != nil {
		return result, 0, fmt.Errorf("read css: %w", err)
	}
	cleaned, removed := RemoveUnusedRules(string(css), result.Unused)
	if removed == 0 {
		return result, 0, nil
	}
	if err := os.WriteFile(cfg.CSSPath, []byte(cleaned), 0o644); err != nil {
		return result, 0, fmt.Errorf("write cleaned css: %w", err)
	}
	return result, removed, nil
}

func ExtractSelectorClasses(css string) []string {
	css = StripComments(css)

	selectorBlock := regexp.MustCompile(`(?s)([^{}]+)\{`)
	classPattern := regexp.MustCompile(`\.([A-Za-z_][A-Za-z0-9_-]*)`)
	seen := map[string]bool{}

	for _, match := range selectorBlock.FindAllStringSubmatch(css, -1) {
		selector := strings.TrimSpace(match[1])
		if selector == "" || strings.HasPrefix(selector, "@") {
			continue
		}
		for _, classMatch := range classPattern.FindAllStringSubmatch(selector, -1) {
			seen[classMatch[1]] = true
		}
	}

	classes := make([]string, 0, len(seen))
	for className := range seen {
		classes = append(classes, className)
	}
	sort.Strings(classes)
	return classes
}

func StripComments(css string) string {
	commentPattern := regexp.MustCompile(`(?s)/\*.*?\*/`)
	return commentPattern.ReplaceAllString(css, "")
}

func CollectSourceFiles(root string, templateDirs []string, extensions []string) ([]string, error) {
	allowed := make(map[string]bool, len(extensions))
	for _, ext := range extensions {
		allowed[ext] = true
	}

	roots := []string{root}
	if len(templateDirs) > 0 {
		roots = roots[:0]
		for _, dir := range templateDirs {
			if strings.TrimSpace(dir) == "" {
				continue
			}
			if filepath.IsAbs(dir) {
				roots = append(roots, dir)
			} else {
				roots = append(roots, filepath.Join(root, dir))
			}
		}
	}

	var files []string
	for _, sourceRoot := range roots {
		err := filepath.WalkDir(sourceRoot, func(path string, entry os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if entry.IsDir() {
				switch entry.Name() {
				case ".git", "node_modules", "dist", "bin":
					return filepath.SkipDir
				default:
					return nil
				}
			}
			if allowed[filepath.Ext(path)] {
				files = append(files, path)
			}
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("scan source files: %w", err)
		}
	}
	sort.Strings(files)
	return files, nil
}

func FindClassUsage(classes []string, sourceFiles []string) (map[string]bool, error) {
	usage := make(map[string]bool, len(classes))
	for _, className := range classes {
		usage[className] = false
	}

	patterns := make(map[string]*regexp.Regexp, len(classes))
	for _, className := range classes {
		patterns[className] = regexp.MustCompile(`(^|[^A-Za-z0-9_-])` + regexp.QuoteMeta(className) + `([^A-Za-z0-9_-]|$)`)
	}

	for _, sourceFile := range sourceFiles {
		data, err := os.ReadFile(sourceFile)
		if err != nil {
			return nil, fmt.Errorf("read source file %s: %w", sourceFile, err)
		}
		content := string(data)
		for className, pattern := range patterns {
			if !usage[className] && pattern.MatchString(content) {
				usage[className] = true
			}
		}
	}
	return usage, nil
}

func SplitExtensions(value string) []string {
	parts := strings.Split(value, ",")
	extensions := make([]string, 0, len(parts))
	for _, part := range parts {
		ext := strings.TrimSpace(part)
		if ext == "" {
			continue
		}
		if !strings.HasPrefix(ext, ".") {
			ext = "." + ext
		}
		extensions = append(extensions, ext)
	}
	return extensions
}

func RemoveUnusedRules(css string, unusedClasses []string) (string, int) {
	unused := make(map[string]bool, len(unusedClasses))
	for _, className := range unusedClasses {
		unused[className] = true
	}
	if len(unused) == 0 {
		return css, 0
	}

	rulePattern := regexp.MustCompile(`(?ms)(\n\s*)([^{}\n][^{}]*\.[A-Za-z_][A-Za-z0-9_-][^{}]*)\{\s*([^{}]*)\}`)
	classPattern := regexp.MustCompile(`\.([A-Za-z_][A-Za-z0-9_-]*)`)
	removed := 0

	cleaned := rulePattern.ReplaceAllStringFunc(css, func(rule string) string {
		matches := rulePattern.FindStringSubmatch(rule)
		if len(matches) != 4 {
			return rule
		}

		prefix := matches[1]
		selectorText := strings.TrimSpace(matches[2])
		body := matches[3]
		selectors := splitSelectorList(selectorText)
		kept := make([]string, 0, len(selectors))

		for _, selector := range selectors {
			classMatches := classPattern.FindAllStringSubmatch(selector, -1)
			if len(classMatches) == 0 {
				kept = append(kept, selector)
				continue
			}

			dead := false
			for _, classMatch := range classMatches {
				if unused[classMatch[1]] {
					dead = true
					break
				}
			}
			if !dead {
				kept = append(kept, selector)
			}
		}

		if len(kept) == 0 {
			removed++
			return ""
		}
		if len(kept) != len(selectors) {
			removed++
		}
		return prefix + strings.Join(kept, ",\n  ") + " {\n" + body + "}"
	})

	cleaned = removeEmptyAtRules(cleaned)
	return cleaned, removed
}

func splitSelectorList(selectorText string) []string {
	parts := strings.Split(selectorText, ",")
	selectors := make([]string, 0, len(parts))
	for _, part := range parts {
		selector := strings.TrimSpace(part)
		if selector != "" {
			selectors = append(selectors, selector)
		}
	}
	return selectors
}

func removeEmptyAtRules(css string) string {
	emptyAtRulePattern := regexp.MustCompile(`(?ms)\n\s*@(?:media|supports)[^{]+\{\s*\}`)
	for {
		next := emptyAtRulePattern.ReplaceAllString(css, "")
		if next == css {
			return css
		}
		css = next
	}
}
