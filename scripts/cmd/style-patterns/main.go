package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/fastygo/ui8kit/scripts/internal/stylepatterns"
)

type config struct {
	CSSFiles    []string
	Output      string
	SourceRoots []string
	Extensions  []string
	Check       bool
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

func parseFlags(args []string) (config, error) {
	cfg := config{
		CSSFiles: []string{
			filepath.FromSlash("styles/components.css"),
			filepath.FromSlash("styles/shell.css"),
		},
		Output:      filepath.FromSlash(".ui8px/policy/patterns.json"),
		SourceRoots: []string{"components", "ui", "layout"},
		Extensions:  []string{".templ", ".go"},
	}

	fs := flag.NewFlagSet("style-patterns", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	cssFiles := fs.String("css", strings.Join(cfg.CSSFiles, ","), "comma-separated CSS files to read")
	sourceRoots := fs.String("source", strings.Join(cfg.SourceRoots, ","), "comma-separated source directories to scan for source-only ui-* classes")
	extensions := fs.String("ext", strings.Join(cfg.Extensions, ","), "comma-separated source extensions")
	fs.StringVar(&cfg.Output, "out", cfg.Output, "patterns policy output path")
	fs.BoolVar(&cfg.Check, "check", false, "verify output is up to date without writing")
	if err := fs.Parse(args); err != nil {
		return cfg, err
	}
	if fs.NArg() != 0 {
		return cfg, errors.New("usage: style-patterns [--css styles/components.css,styles/shell.css] [--source components,ui,layout] [--out .ui8px/policy/patterns.json] [--check]")
	}
	cfg.CSSFiles = splitCSV(*cssFiles)
	if len(cfg.CSSFiles) == 0 {
		return cfg, errors.New("at least one CSS file is required")
	}
	cfg.SourceRoots = splitCSV(*sourceRoots)
	cfg.Extensions = splitCSV(*extensions)
	return cfg, nil
}

func run(cfg config) error {
	policy, err := stylepatterns.Generate(stylepatterns.Config{
		CSSFiles:    cfg.CSSFiles,
		SourceRoots: cfg.SourceRoots,
		Extensions:  cfg.Extensions,
	})
	if err != nil {
		return err
	}
	data := stylepatterns.Format(policy)
	if cfg.Check {
		current, err := os.ReadFile(cfg.Output)
		if err != nil {
			return fmt.Errorf("read output policy: %w", err)
		}
		if string(current) != string(data) {
			return fmt.Errorf("%s is out of date; run style-patterns", cfg.Output)
		}
		fmt.Printf("style-patterns: %s is up to date (%d patterns)\n", cfg.Output, len(policy.Patterns))
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(cfg.Output), 0o755); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}
	if err := os.WriteFile(cfg.Output, data, 0o644); err != nil {
		return fmt.Errorf("write output policy: %w", err)
	}
	fmt.Printf("style-patterns: wrote %d patterns to %s\n", len(policy.Patterns), cfg.Output)
	return nil
}

func splitCSV(value string) []string {
	parts := strings.Split(value, ",")
	values := make([]string, 0, len(parts))
	for _, part := range parts {
		item := strings.TrimSpace(part)
		if item != "" {
			values = append(values, item)
		}
	}
	return values
}
