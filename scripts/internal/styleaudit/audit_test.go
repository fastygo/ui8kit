package styleaudit

import (
	"errors"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

func TestExtractSelectorClasses(t *testing.T) {
	css := `
@layer components {
  .ui-card {
    @apply rounded;
  }

  .ui-dialog[data-state="open"],
  .ui-dialog-overlay:hover {
    @apply block;
  }

  @media (max-width: 767px) {
    .ui-shell-body--marketing .ui-header {
      @apply flex;
    }
  }
}
`

	got := ExtractSelectorClasses(css)
	want := []string{
		"ui-card",
		"ui-dialog",
		"ui-dialog-overlay",
		"ui-header",
		"ui-shell-body--marketing",
	}
	for _, className := range want {
		if !slices.Contains(got, className) {
			t.Fatalf("missing class %q in %#v", className, got)
		}
	}
}

func TestRunReportsUnusedTemplClasses(t *testing.T) {
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "styles", "components.css"), `
@layer components {
  .ui-used { @apply block; }
  .ui-unused { @apply hidden; }
  .ui-card[data-state="open"] { @apply block; }
}
`)
	writeFile(t, filepath.Join(root, "components", "card.templ"), `
package components

templ Card() {
  <div class="ui-used ui-card"></div>
}
`)

	result, err := Run(Config{
		CSSPath:      filepath.Join(root, "styles", "components.css"),
		Root:         root,
		Extensions:   []string{".templ"},
		TemplateDirs: []string{"components"},
	})
	var auditErr UnusedClassesError
	if !errors.As(err, &auditErr) {
		t.Fatalf("expected UnusedClassesError, got %v", err)
	}
	if auditErr.Count != 1 {
		t.Fatalf("expected 1 unused class, got %d", auditErr.Count)
	}
	if len(result.Unused) != 1 || result.Unused[0] != "ui-unused" {
		t.Fatalf("unexpected unused classes: %#v", result.Unused)
	}
}

func TestSplitExtensions(t *testing.T) {
	got := SplitExtensions(".templ,go,  .txt ")
	want := []string{".templ", ".go", ".txt"}
	if !slices.Equal(got, want) {
		t.Fatalf("SplitExtensions() = %#v, want %#v", got, want)
	}
}

func TestRemoveUnusedRules(t *testing.T) {
	css := `
@layer components {
  .ui-used {
    @apply block;
  }

  .ui-unused {
    @apply hidden;
  }

  .ui-used .ui-unused-child,
  .ui-used:hover {
    @apply block;
  }

  .ui-preserved {
    @apply flex items-center;
  }
}
`

	cleaned, removed := RemoveUnusedRules(css, []string{"ui-unused", "ui-unused-child"})
	if removed != 1 {
		t.Fatalf("removed = %d, want 1", removed)
	}
	if strings.Contains(cleaned, "ui-unused {") {
		t.Fatalf("unused rule was not removed:\n%s", cleaned)
	}
	if !strings.Contains(cleaned, ".ui-used .ui-unused-child") {
		t.Fatalf("mixed selector rule should be preserved:\n%s", cleaned)
	}
	if strings.Contains(cleaned, "ui-unused-child") {
		t.Log("mixed selector with unused class was preserved to avoid rewriting live CSS formatting")
	}
	preserved := `  .ui-preserved {
    @apply flex items-center;
  }`
	if !strings.Contains(cleaned, preserved) {
		t.Fatalf("live rule formatting was not preserved:\n%s", cleaned)
	}
	if !strings.Contains(cleaned, ".ui-used:hover") {
		t.Fatalf("live selector was removed:\n%s", cleaned)
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
