package styles_test

import (
	"testing"

	"github.com/fastygo/ui8kit/styles"
)

func TestEmbeddedFiles(t *testing.T) {
	for _, name := range []string{"base.css", "shell.css", "components.css", "latty.css"} {
		data, err := styles.FS.ReadFile(name)
		if err != nil {
			t.Fatalf("ReadFile(%q): %v", name, err)
		}
		if len(data) == 0 {
			t.Fatalf("%q is empty", name)
		}
	}
}
