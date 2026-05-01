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
		{"li", TagGroupListItem, true},
		{"dt", TagGroupListItem, true},
		{"dd", TagGroupListItem, true},
		{"fieldset", TagGroupForm, true},
		{"input", TagGroupFormControl, true},
		{"textarea", TagGroupFormControl, true},
		{"select", TagGroupFormControl, true},
		{"button", TagGroupFormControl, true},
		{"option", TagGroupFormControl, true},
		{"optgroup", TagGroupFormControl, true},
		{"datalist", TagGroupFormControl, true},
		{"output", TagGroupFormControl, true},
		{"meter", TagGroupFormControl, true},
		{"progress", TagGroupFormControl, true},
		{"label", TagGroupFormLabel, true},
		{"legend", TagGroupFormLabel, true},
		{"table", TagGroupTable, true},
		{"thead", TagGroupTableSection, true},
		{"tbody", TagGroupTableSection, true},
		{"tfoot", TagGroupTableSection, true},
		{"tr", TagGroupTableRow, true},
		{"th", TagGroupTableCell, true},
		{"td", TagGroupTableCell, true},
		{"colgroup", TagGroupTableColumn, true},
		{"col", TagGroupTableColumn, true},
		{"img", TagGroupMedia, true},
		{"picture", TagGroupMedia, true},
		{"source", TagGroupMedia, true},
		{"details", TagGroupDisclosure, true},
		{"summary", TagGroupDisclosure, true},
		{"ul", TagGroupStack, true},
		{"fieldset", TagGroupGroup, true},
		{"span", TagGroupText, true},
		{"section", TagGroupContainer, true},
		{"table", TagGroupLayout, false},
		{"table", TagGroupText, false},
		{"button", TagGroupLayout, false},
		{"input", TagGroupLayout, false},
		{"img", TagGroupLayout, false},
		{"li", TagGroupLayout, false},
		{"dt", TagGroupLayout, false},
		{"dd", TagGroupLayout, false},
	}

	for _, tc := range cases {
		if got := IsAllowedTag(tc.tag, tc.group); got != tc.want {
			t.Fatalf("IsAllowedTag(%q, %v) = %v, want %v", tc.tag, tc.group, got, tc.want)
		}
	}
}
