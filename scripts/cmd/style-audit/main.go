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
}

func main() {
	cfg, err := parseFlags(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	result, err := styleaudit.Run(styleaudit.Config(cfg))
	if err != nil {
		var auditErr styleaudit.UnusedClassesError
		if errors.As(err, &auditErr) {
			printResult(os.Stderr, result)
			os.Exit(1)
		}
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	printResult(os.Stdout, result)
}

func parseFlags(args []string) (config, error) {
	cfg := config{
		CSSPath:    filepath.FromSlash("styles/components.css"),
		Root:       ".",
		Extensions: []string{".templ"},
	}

	fs := flag.NewFlagSet("style-audit", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	fs.StringVar(&cfg.CSSPath, "css", cfg.CSSPath, "CSS file to audit")
	fs.StringVar(&cfg.Root, "root", cfg.Root, "root directory to scan")
	extensions := fs.String("ext", ".templ", "comma-separated source extensions to scan")
	templatesMode := fs.Bool("templates", false, "treat remaining arguments as template/source directories")
	if err := fs.Parse(args); err != nil {
		return cfg, err
	}
	if fs.NArg() != 0 && !*templatesMode {
		return cfg, errors.New("usage: style-audit [--css styles/components.css] [--root .] [--ext .templ] [--templates components ui layout]")
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

func printResult(w io.Writer, result styleaudit.Result) {
	if len(result.Unused) == 0 {
		fmt.Fprintf(w, "style-audit: %d selector classes are used across %d source files\n", len(result.Classes), len(result.SourceFiles))
		return
	}

	fmt.Fprintf(w, "style-audit: %d of %d selector classes are unused across %d source files\n", len(result.Unused), len(result.Classes), len(result.SourceFiles))
	for _, className := range result.Unused {
		fmt.Fprintf(w, "- %s\n", className)
	}
}
