package utils

import "strings"

type TagGroup int

const (
	TagGroupLayout TagGroup = iota
	TagGroupBlockText
	TagGroupInline
	TagGroupHeading
	TagGroupList
	TagGroupListItem
	TagGroupForm
	TagGroupFormControl
	TagGroupFormLabel
	TagGroupTable
	TagGroupTableSection
	TagGroupTableRow
	TagGroupTableCell
	TagGroupTableColumn
	TagGroupMedia
	TagGroupDisclosure
	TagGroupStack
	TagGroupGroup
	TagGroupText
	TagGroupContainer
)

func ResolveTag(tag, fallback string, group TagGroup) string {
	t := strings.TrimSpace(strings.ToLower(tag))
	if t == "" {
		return fallback
	}
	if IsAllowedTag(t, group) {
		return t
	}
	return fallback
}

func IsAllowedTag(tag string, group TagGroup) bool {
	t := strings.TrimSpace(strings.ToLower(tag))
	switch group {
	case TagGroupLayout:
		return oneOf(t, "div", "section", "article", "aside", "header", "footer", "main", "nav", "figure")
	case TagGroupBlockText:
		return oneOf(t, "p", "blockquote", "figcaption", "address", "pre")
	case TagGroupInline:
		return oneOf(t, "span", "em", "strong", "small", "abbr", "cite", "code", "kbd", "mark", "time", "data", "var", "samp", "sub", "sup", "b", "i", "u", "s", "q", "dfn", "bdo", "bdi")
	case TagGroupHeading:
		return oneOf(t, "h1", "h2", "h3", "h4", "h5", "h6")
	case TagGroupList:
		return oneOf(t, "ul", "ol", "dl", "menu")
	case TagGroupListItem:
		return oneOf(t, "li", "dt", "dd")
	case TagGroupForm:
		return oneOf(t, "form", "fieldset")
	case TagGroupFormControl:
		return oneOf(t, "input", "textarea", "select", "button", "option", "optgroup", "datalist", "output", "meter", "progress")
	case TagGroupFormLabel:
		return oneOf(t, "label", "legend")
	case TagGroupTable:
		return oneOf(t, "table")
	case TagGroupTableSection:
		return oneOf(t, "thead", "tbody", "tfoot")
	case TagGroupTableRow:
		return oneOf(t, "tr")
	case TagGroupTableCell:
		return oneOf(t, "th", "td")
	case TagGroupTableColumn:
		return oneOf(t, "colgroup", "col")
	case TagGroupMedia:
		return oneOf(t, "img", "picture", "source")
	case TagGroupDisclosure:
		return oneOf(t, "details", "summary")
	case TagGroupStack:
		return IsAllowedTag(t, TagGroupLayout) || oneOf(t, "ul", "ol")
	case TagGroupGroup:
		return IsAllowedTag(t, TagGroupLayout) || oneOf(t, "fieldset")
	case TagGroupText:
		return IsAllowedTag(t, TagGroupBlockText) || oneOf(t, "span", "small", "time", "em", "strong", "cite", "abbr", "code", "mark")
	case TagGroupContainer:
		return oneOf(t, "div", "main", "section")
	default:
		return false
	}
}

func oneOf(value string, options ...string) bool {
	for _, opt := range options {
		if value == opt {
			return true
		}
	}
	return false
}
