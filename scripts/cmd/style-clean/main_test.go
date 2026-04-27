package main

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

func TestParseFlagsWithTemplateDirs(t *testing.T) {
	cfg, err := parseFlags([]string{"--css", "styles/components.css", "--templates", "components", "ui", "layout"})
	if err != nil {
		t.Fatalf("parseFlags() error = %v", err)
	}
	want := []string{"components", "ui", "layout"}
	if !slices.Equal(cfg.TemplateDirs, want) {
		t.Fatalf("TemplateDirs = %#v, want %#v", cfg.TemplateDirs, want)
	}
}

func TestRunCleansUnusedCSS(t *testing.T) {
	root := t.TempDir()
	cssPath := filepath.Join(root, "styles", "components.css")
	writeFile(t, cssPath, `
@layer components {
  .ui-used {
    @apply block;
  }

  .ui-unused {
    @apply hidden;
  }
}
`)
	writeFile(t, filepath.Join(root, "components", "card.templ"), `
package components

templ Card() {
  <div class="ui-used"></div>
}
`)

	_, removed, err := run(config{
		CSSPath:      cssPath,
		Root:         root,
		Extensions:   []string{".templ"},
		TemplateDirs: []string{"components"},
	})
	if err != nil {
		t.Fatalf("run() error = %v", err)
	}
	if removed != 1 {
		t.Fatalf("removed = %d, want 1", removed)
	}

	cleaned, err := os.ReadFile(cssPath)
	if err != nil {
		t.Fatalf("read cleaned css: %v", err)
	}
	if strings.Contains(string(cleaned), "ui-unused") {
		t.Fatalf("unused rule still present:\n%s", cleaned)
	}
}

func writeFile(t *testing.T, path string, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
}
