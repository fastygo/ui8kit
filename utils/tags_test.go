package utils

import "testing"

func TestResolveTag(t *testing.T) {
	if got := ResolveTag("", "div", TagGroupLayout); got != "div" {
		t.Fatalf("expected fallback div, got %q", got)
	}
	if got := ResolveTag("section", "div", TagGroupLayout); got != "section" {
		t.Fatalf("expected section, got %q", got)
	}
	if got := ResolveTag("table", "div", TagGroupLayout); got != "div" {
		t.Fatalf("expected fallback div for invalid tag, got %q", got)
	}
}

func TestIsAllowedTagGroups(t *testing.T) {
	cases := []struct {
		tag   string
		group TagGroup
		want  bool
	}{
		{"main", TagGroupLayout, true},
		{"p", TagGroupBlockText, true},
		{"span", TagGroupInline, true},
		{"h3", TagGroupHeading, true},
		{"ul", TagGroupList, true},
		{"fieldset", TagGroupForm, true},
		{"ul", TagGroupStack, true},
		{"fieldset", TagGroupGroup, true},
		{"span", TagGroupText, true},
		{"section", TagGroupContainer, true},
		{"table", TagGroupText, false},
	}

	for _, tc := range cases {
		if got := IsAllowedTag(tc.tag, tc.group); got != tc.want {
			t.Fatalf("IsAllowedTag(%q, %v) = %v, want %v", tc.tag, tc.group, got, tc.want)
		}
	}
}
