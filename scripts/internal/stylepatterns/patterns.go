package stylepatterns

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type Config struct {
	CSSFiles    []string
	Output      string
	SourceRoots []string
	Extensions  []string
}

type Policy struct {
	Patterns map[string][]string `json:"patterns"`
}

func Generate(cfg Config) (Policy, error) {
	policy := Policy{Patterns: map[string][]string{}}
	for _, cssFile := range cfg.CSSFiles {
		data, err := os.ReadFile(cssFile)
		if err != nil {
			return Policy{}, fmt.Errorf("read css %s: %w", cssFile, err)
		}
		for className, tokens := range ExtractPatterns(string(data)) {
			policy.Patterns[className] = mergeTokens(policy.Patterns[className], tokens)
		}
	}
	sourceClasses, err := collectSourceClasses(cfg.SourceRoots, cfg.Extensions)
	if err != nil {
		return Policy{}, err
	}
	for _, className := range sourceClasses {
		if _, ok := policy.Patterns[className]; !ok {
			policy.Patterns[className] = []string{}
		}
	}
	return policy, nil
}

func Write(cfg Config) (Policy, error) {
	policy, err := Generate(cfg)
	if err != nil {
		return Policy{}, err
	}
	if cfg.Output == "" {
		return Policy{}, fmt.Errorf("output path is required")
	}
	if err := os.MkdirAll(filepath.Dir(cfg.Output), 0o755); err != nil {
		return Policy{}, fmt.Errorf("create output directory: %w", err)
	}
	if err := os.WriteFile(cfg.Output, Format(policy), 0o644); err != nil {
		return Policy{}, fmt.Errorf("write patterns policy: %w", err)
	}
	return policy, nil
}

func ExtractPatterns(css string) map[string][]string {
	css = stripComments(css)
	css = removeNestedAtRules(css, []string{"@media", "@supports"})

	rulePattern := regexp.MustCompile(`(?ms)([^{}]+)\{([^{}]*)\}`)
	selectorPattern := regexp.MustCompile(`^\.(ui-[A-Za-z0-9_-]+)$`)
	applyPattern := regexp.MustCompile(`@apply\s+([^;]+);`)
	patterns := map[string][]string{}

	for _, match := range rulePattern.FindAllStringSubmatch(css, -1) {
		selector := strings.TrimSpace(match[1])
		body := match[2]
		if strings.Contains(selector, ",") {
			continue
		}
		selectorMatch := selectorPattern.FindStringSubmatch(selector)
		if len(selectorMatch) != 2 {
			continue
		}
		className := selectorMatch[1]
		var tokens []string
		for _, applyMatch := range applyPattern.FindAllStringSubmatch(body, -1) {
			tokens = append(tokens, strings.Fields(applyMatch[1])...)
		}
		if len(tokens) > 0 {
			patterns[className] = mergeTokens(patterns[className], tokens)
		}
	}
	return patterns
}

func Format(policy Policy) []byte {
	keys := make([]string, 0, len(policy.Patterns))
	for key := range policy.Patterns {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var out bytes.Buffer
	out.WriteString("{\n")
	out.WriteString("  \"patterns\": {\n")
	for i, key := range keys {
		tokens, _ := json.Marshal(policy.Patterns[key])
		keyJSON, _ := json.Marshal(key)
		out.WriteString("    ")
		out.Write(keyJSON)
		out.WriteString(": ")
		out.Write(tokens)
		if i < len(keys)-1 {
			out.WriteString(",")
		}
		out.WriteString("\n")
	}
	out.WriteString("  }\n")
	out.WriteString("}\n")
	return out.Bytes()
}

func collectSourceClasses(sourceRoots []string, extensions []string) ([]string, error) {
	if len(sourceRoots) == 0 {
		return nil, nil
	}
	if len(extensions) == 0 {
		extensions = []string{".templ", ".go"}
	}
	allowed := map[string]bool{}
	for _, ext := range extensions {
		if strings.TrimSpace(ext) == "" {
			continue
		}
		if !strings.HasPrefix(ext, ".") {
			ext = "." + ext
		}
		allowed[ext] = true
	}

	classPattern := regexp.MustCompile(`ui-[A-Za-z0-9_-]+`)
	seen := map[string]bool{}
	for _, root := range sourceRoots {
		if strings.TrimSpace(root) == "" {
			continue
		}
		err := filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error {
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
			if !allowed[filepath.Ext(path)] {
				return nil
			}
			data, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("read source %s: %w", path, err)
			}
			for _, match := range classPattern.FindAllString(string(data), -1) {
				seen[match] = true
			}
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("scan source classes: %w", err)
		}
	}

	classes := make([]string, 0, len(seen))
	for className := range seen {
		classes = append(classes, className)
	}
	sort.Strings(classes)
	return classes, nil
}

func mergeTokens(existing []string, next []string) []string {
	seen := make(map[string]bool, len(existing)+len(next))
	merged := make([]string, 0, len(existing)+len(next))
	for _, token := range append(existing, next...) {
		if strings.TrimSpace(token) == "" || seen[token] {
			continue
		}
		seen[token] = true
		merged = append(merged, token)
	}
	return merged
}

func stripComments(css string) string {
	commentPattern := regexp.MustCompile(`(?s)/\*.*?\*/`)
	return commentPattern.ReplaceAllString(css, "")
}

func removeNestedAtRules(css string, atRules []string) string {
	var out strings.Builder
	for i := 0; i < len(css); {
		atRule := matchingAtRule(css[i:], atRules)
		if atRule == "" {
			out.WriteByte(css[i])
			i++
			continue
		}

		start := i
		open := strings.IndexByte(css[start:], '{')
		if open < 0 {
			out.WriteString(css[start:])
			break
		}
		open += start
		end := matchingBrace(css, open)
		if end < 0 {
			out.WriteString(css[start:])
			break
		}
		i = end + 1
	}
	return out.String()
}

func matchingAtRule(value string, atRules []string) string {
	trimmed := strings.TrimLeft(value, " \t\r\n")
	if len(trimmed) == len(value) {
		for _, atRule := range atRules {
			if strings.HasPrefix(value, atRule) {
				return atRule
			}
		}
		return ""
	}
	leading := len(value) - len(trimmed)
	for _, atRule := range atRules {
		if strings.HasPrefix(trimmed, atRule) {
			return strings.Repeat(" ", leading) + atRule
		}
	}
	return ""
}

func matchingBrace(value string, open int) int {
	depth := 0
	for i := open; i < len(value); i++ {
		switch value[i] {
		case '{':
			depth++
		case '}':
			depth--
			if depth == 0 {
				return i
			}
		}
	}
	return -1
}
