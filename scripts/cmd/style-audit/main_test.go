package main

import (
	"slices"
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
	if !slices.Equal(cfg.Extensions, []string{".templ"}) {
		t.Fatalf("Extensions = %#v, want [.templ]", cfg.Extensions)
	}
}
