package main

import (
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func TestCollectTemplUtilityClasses(t *testing.T) {
	tmpDir := t.TempDir()
	sourcePath := filepath.Join(tmpDir, "page_templ.go")
	source := `package views

import "github.com/fastygo/ui8kit/utils"

func view() {
	_ = utils.UtilityProps{P: "2", Rounded: "md", Hidden: true}
	_ = utils.UtilityProps{Flex: "col", Gap: "sm"}
}
`
	if err := os.WriteFile(sourcePath, []byte(source), 0o644); err != nil {
		t.Fatalf("write source: %v", err)
	}

	classes, err := collectTemplUtilityClasses(tmpDir)
	if err != nil {
		t.Fatalf("collectTemplUtilityClasses() error = %v", err)
	}

	want := []string{"flex", "flex-col", "gap-2", "hidden", "p-2", "rounded-md"}
	for _, className := range want {
		if !slices.Contains(classes, className) {
			t.Fatalf("missing class %q in %v", className, classes)
		}
	}
}

