package stylepatterns

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

func TestExtractPatternsUsesOnlyBaseUIRules(t *testing.T) {
	css := `
@layer components {
  .ui-card {
    @apply rounded border border-border;
  }

  .ui-card {
    @apply bg-card;
  }

  .ui-card:hover {
    @apply shadow;
  }

  .ui-shell-body--marketing .ui-header {
    @apply grid;
  }

  .ui-dialog[data-state="open"] {
    @apply block;
  }

  .not-ui {
    @apply hidden;
  }

  @media (max-width: 48rem) {
    .ui-card {
      @apply p-4;
    }
  }
}
`

	patterns := ExtractPatterns(css)
	want := []string{"rounded", "border", "border-border", "bg-card"}
	if !slices.Equal(patterns["ui-card"], want) {
		t.Fatalf("ui-card pattern = %#v, want %#v", patterns["ui-card"], want)
	}
	for _, absent := range []string{"ui-shell-body--marketing", "ui-header", "ui-dialog", "not-ui"} {
		if _, ok := patterns[absent]; ok {
			t.Fatalf("unexpected pattern for %q: %#v", absent, patterns[absent])
		}
	}
}

func TestFormatIsStable(t *testing.T) {
	out := string(Format(Policy{Patterns: map[string][]string{
		"ui-b": {"flex"},
		"ui-a": {"grid", "gap-2"},
	}}))
	want := `{
  "patterns": {
    "ui-a": ["grid","gap-2"],
    "ui-b": ["flex"]
  }
}
`
	if out != want {
		t.Fatalf("Format() =\n%s\nwant:\n%s", out, want)
	}
}

func TestGenerateIncludesSourceOnlyClasses(t *testing.T) {
	root := t.TempDir()
	cssPath := filepath.Join(root, "components.css")
	sourceRoot := filepath.Join(root, "components")
	if err := os.MkdirAll(sourceRoot, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := os.WriteFile(cssPath, []byte(`.ui-card { @apply rounded; }`), 0o644); err != nil {
		t.Fatalf("write css: %v", err)
	}
	if err := os.WriteFile(filepath.Join(sourceRoot, "card.templ"), []byte(`class="ui-card ui-source-only"`), 0o644); err != nil {
		t.Fatalf("write source: %v", err)
	}

	policy, err := Generate(Config{
		CSSFiles:    []string{cssPath},
		SourceRoots: []string{sourceRoot},
		Extensions:  []string{".templ"},
	})
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}
	if !slices.Equal(policy.Patterns["ui-card"], []string{"rounded"}) {
		t.Fatalf("ui-card = %#v", policy.Patterns["ui-card"])
	}
	if pattern, ok := policy.Patterns["ui-source-only"]; !ok || len(pattern) != 0 {
		t.Fatalf("source-only pattern = %#v, exists=%v", pattern, ok)
	}
}

func TestRemoveNestedAtRules(t *testing.T) {
	css := `.ui-a { @apply flex; }
@media (max-width: 48rem) {
  .ui-a { @apply grid; }
}
.ui-b { @apply block; }`
	got := removeNestedAtRules(css, []string{"@media"})
	if strings.Contains(got, "grid") {
		t.Fatalf("media block was not removed:\n%s", got)
	}
	if !strings.Contains(got, ".ui-a") || !strings.Contains(got, ".ui-b") {
		t.Fatalf("base rules were removed:\n%s", got)
	}
}
