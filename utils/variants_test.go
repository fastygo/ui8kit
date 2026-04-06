package utils

import (
	"strings"
	"testing"
)

func TestButtonStyleVariant(t *testing.T) {
	for _, v := range []string{"", "default", "primary", "destructive", "outline", "secondary", "ghost", "link"} {
		if got := ButtonStyleVariant(v); got == "" {
			t.Fatalf("ButtonStyleVariant(%q) returned empty", v)
		}
	}
	if got := ButtonStyleVariant("custom-class"); !strings.Contains(got, "custom-class") {
		t.Fatal("fallback should include the raw variant string")
	}
}

func TestButtonSizeVariant(t *testing.T) {
	for _, s := range []string{"xs", "sm", "", "default", "md", "lg", "xl", "icon"} {
		if got := ButtonSizeVariant(s); got == "" {
			t.Fatalf("ButtonSizeVariant(%q) returned empty", s)
		}
	}
}

func TestBadgeStyleVariant(t *testing.T) {
	for _, v := range []string{"", "default", "primary", "success", "secondary", "destructive", "outline", "warning", "info"} {
		if got := BadgeStyleVariant(v); got == "" {
			t.Fatalf("BadgeStyleVariant(%q) returned empty", v)
		}
	}
}

func TestBadgeSizeVariant(t *testing.T) {
	for _, s := range []string{"xs", "sm", "", "default", "lg"} {
		if got := BadgeSizeVariant(s); got == "" {
			t.Fatalf("BadgeSizeVariant(%q) returned empty", s)
		}
	}
}

func TestCardVariant(t *testing.T) {
	for _, v := range []string{"", "default", "raised", "kpi", "muted", "ghost", "compact", "flat", "accent"} {
		if got := CardVariant(v); got == "" {
			t.Fatalf("CardVariant(%q) returned empty", v)
		}
	}
	if got := CardVariant("kpi"); !strings.Contains(got, "kit-card--kpi") {
		t.Fatalf("expected kpi modifier in %q", got)
	}
	if got := CardVariant("custom-extra"); !strings.Contains(got, "custom-extra") {
		t.Fatal("fallback should append raw classes")
	}
}

func TestTypographyClasses(t *testing.T) {
	got := TypographyClasses("sm", "medium", "6", "tight", "muted-foreground", "left", true)
	if got == "" {
		t.Fatal("TypographyClasses returned empty")
	}
	for _, want := range []string{"text-sm", "font-medium", "leading-6", "tracking-tight", "text-muted-foreground", "text-left", "truncate"} {
		if !strings.Contains(got, want) {
			t.Fatalf("expected %q in %q", want, got)
		}
	}
}

func TestFieldVariants(t *testing.T) {
	for _, v := range []string{"", "default", "outline", "ghost"} {
		if got := FieldVariant(v); got == "" {
			t.Fatalf("FieldVariant(%q) returned empty", v)
		}
	}
	for _, s := range []string{"xs", "sm", "", "default", "md", "lg"} {
		if got := FieldSizeVariant(s); got == "" {
			t.Fatalf("FieldSizeVariant(%q) returned empty", s)
		}
	}
}

func TestFieldControlVariants(t *testing.T) {
	for _, v := range []string{"", "default", "outline", "ghost"} {
		if got := FieldControlVariant(v); got == "" {
			t.Fatalf("FieldControlVariant(%q) returned empty", v)
		}
	}
	for _, s := range []string{"xs", "sm", "", "default", "md", "lg"} {
		if got := FieldControlSizeVariant(s); got == "" {
			t.Fatalf("FieldControlSizeVariant(%q) returned empty", s)
		}
	}
}
