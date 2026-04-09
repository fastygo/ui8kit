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
	if got := CardVariant("kpi"); !strings.Contains(got, "ui-card--kpi") {
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

func TestImageVariants(t *testing.T) {
	if got := ImageFitVariant("contain"); got != "object-contain" {
		t.Fatalf("unexpected fit: %q", got)
	}
	if got := ImagePositionVariant("left-top"); got != "object-left-top" {
		t.Fatalf("unexpected position: %q", got)
	}
	if got := ImageAspectVariant("video"); got != "aspect-video" {
		t.Fatalf("unexpected aspect: %q", got)
	}
}

func TestGridVariants(t *testing.T) {
	if got := GridColsVariant("3"); got != "grid-cols-3" {
		t.Fatalf("unexpected cols: %q", got)
	}
	if got := GridColsVariant("1-3"); got == "" || !strings.Contains(got, "xl:grid-cols-3") {
		t.Fatalf("unexpected preset cols: %q", got)
	}
	if got := GridColVariant(3, 2, 0, 1); !strings.Contains(got, "col-span-3") || !strings.Contains(got, "col-start-2") || !strings.Contains(got, "order-1") {
		t.Fatalf("unexpected col variant: %q", got)
	}
}

func TestSheetVariants(t *testing.T) {
	if got := SheetSizeVariant("lg"); got != "w-96" {
		t.Fatalf("unexpected sheet size: %q", got)
	}
	if got := SheetSideVariant("left"); got != "left-0 border-r" {
		t.Fatalf("unexpected sheet side: %q", got)
	}
}
