package utils

import "testing"

func TestCn(t *testing.T) {
	tests := []struct {
		name    string
		classes []string
		want    string
	}{
		{"empty", nil, ""},
		{"single", []string{"foo"}, "foo"},
		{"multiple", []string{"a", "b", "c"}, "a b c"},
		{"skips blanks", []string{"a", "", "  ", "b"}, "a b"},
		{"trims spaces", []string{" a ", " b"}, "a b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cn(tt.classes...); got != tt.want {
				t.Fatalf("Cn(%v) = %q, want %q", tt.classes, got, tt.want)
			}
		})
	}
}
