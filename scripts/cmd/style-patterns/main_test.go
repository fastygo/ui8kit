package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParseFlags(t *testing.T) {
	cfg, err := parseFlags([]string{"--css", "a.css,b.css", "--source", "components,ui", "--out", "patterns.json", "--check"})
	if err != nil {
		t.Fatalf("parseFlags() error = %v", err)
	}
	if len(cfg.CSSFiles) != 2 || cfg.CSSFiles[0] != "a.css" || cfg.CSSFiles[1] != "b.css" {
		t.Fatalf("unexpected css files: %#v", cfg.CSSFiles)
	}
	if cfg.Output != "patterns.json" {
		t.Fatalf("Output = %q", cfg.Output)
	}
	if len(cfg.SourceRoots) != 2 || cfg.SourceRoots[0] != "components" || cfg.SourceRoots[1] != "ui" {
		t.Fatalf("unexpected source roots: %#v", cfg.SourceRoots)
	}
	if !cfg.Check {
		t.Fatal("Check = false, want true")
	}
}

func TestRunWritesPolicy(t *testing.T) {
	root := t.TempDir()
	cssPath := filepath.Join(root, "components.css")
	outPath := filepath.Join(root, "patterns.json")
	if err := os.WriteFile(cssPath, []byte(`.ui-card { @apply rounded border; }`), 0o644); err != nil {
		t.Fatalf("write css: %v", err)
	}
	if err := run(config{CSSFiles: []string{cssPath}, Output: outPath}); err != nil {
		t.Fatalf("run() error = %v", err)
	}
	data, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatalf("read output: %v", err)
	}
	if !strings.Contains(string(data), `"ui-card": ["rounded","border"]`) {
		t.Fatalf("unexpected output:\n%s", data)
	}
}
