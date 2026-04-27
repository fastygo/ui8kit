package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/fastygo/ui8kit/scripts/internal/styleaudit"
)

type config struct {
	CSSPath      string
	Root         string
	Extensions   []string
	TemplateDirs []string
	DryRun       bool
}

func main() {
	cfg, err := parseFlags(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	result, removed, err := run(cfg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if cfg.DryRun {
		fmt.Printf("style-clean: would remove %d CSS rules/selectors for %d unused classes\n", removed, len(result.Unused))
		return
	}
	fmt.Printf("style-clean: removed %d CSS rules/selectors for %d unused classes\n", removed, len(result.Unused))
}

func parseFlags(args []string) (config, error) {
	cfg := config{
		CSSPath:    filepath.FromSlash("styles/components.css"),
		Root:       ".",
		Extensions: []string{".templ"},
	}

	fs := flag.NewFlagSet("style-clean", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	fs.StringVar(&cfg.CSSPath, "css", cfg.CSSPath, "CSS file to clean")
	fs.StringVar(&cfg.Root, "root", cfg.Root, "root directory to scan")
	fs.BoolVar(&cfg.DryRun, "dry-run", false, "print removable rules without writing CSS")
	extensions := fs.String("ext", ".templ", "comma-separated source extensions to scan")
	templatesMode := fs.Bool("templates", false, "treat remaining arguments as template/source directories")
	if err := fs.Parse(args); err != nil {
		return cfg, err
	}
	if fs.NArg() != 0 && !*templatesMode {
		return cfg, errors.New("usage: style-clean [--css styles/components.css] [--root .] [--ext .templ] [--dry-run] [--templates components ui layout]")
	}
	if *templatesMode {
		cfg.TemplateDirs = fs.Args()
		if len(cfg.TemplateDirs) == 0 {
			return cfg, errors.New("--templates requires at least one directory")
		}
	}
	cfg.Extensions = styleaudit.SplitExtensions(*extensions)
	if len(cfg.Extensions) == 0 {
		return cfg, errors.New("at least one source extension is required")
	}
	return cfg, nil
}

func run(cfg config) (styleaudit.Result, int, error) {
	auditCfg := styleaudit.Config{
		CSSPath:      cfg.CSSPath,
		Root:         cfg.Root,
		Extensions:   cfg.Extensions,
		TemplateDirs: cfg.TemplateDirs,
	}

	if cfg.DryRun {
		result, err := styleaudit.Run(auditCfg)
		var auditErr styleaudit.UnusedClassesError
		if err != nil && !errors.As(err, &auditErr) {
			return result, 0, err
		}
		css, err := os.ReadFile(cfg.CSSPath)
		if err != nil {
			return result, 0, fmt.Errorf("read css: %w", err)
		}
		_, removed := styleaudit.RemoveUnusedRules(string(css), result.Unused)
		return result, removed, nil
	}

	return styleaudit.CleanFile(auditCfg)
}
