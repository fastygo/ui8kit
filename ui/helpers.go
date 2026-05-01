package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/a-h/templ"
	"github.com/fastygo/ui8kit/utils"
)

func resolveTag(tag, fallback string) string {
	return utils.ResolveTag(tag, fallback, utils.TagGroupLayout)
}

func buttonClasses(p ButtonProps) string {
	state := ""
	if p.Disabled {
		state = "ui-button-disabled"
	}
	if strings.TrimSpace(p.Variant) == "unstyled" && strings.TrimSpace(p.Size) == "" {
		return utils.Cn(utils.ButtonStyleVariant(p.Variant), state, p.Class)
	}
	return utils.Cn(utils.ButtonStyleVariant(p.Variant), utils.ButtonSizeVariant(p.Size), state, p.Class)
}

func domAttrs(p DOMProps) templ.Attributes {
	attrs := templ.Attributes{}
	for key, value := range p.Attrs {
		attrs[key] = value
	}
	if strings.TrimSpace(p.ID) != "" {
		attrs["id"] = p.ID
	}
	if strings.TrimSpace(p.Role) != "" {
		attrs["role"] = p.Role
	}
	if strings.TrimSpace(p.TabIndex) != "" {
		attrs["tabindex"] = p.TabIndex
	}
	return attrs
}

func buttonAttrs(p ButtonProps) templ.Attributes {
	attrs := domAttrs(DOMProps{ID: p.ID, Role: p.Role, TabIndex: p.TabIndex, Attrs: p.Attrs})
	if strings.TrimSpace(p.AriaLabel) != "" {
		attrs["aria-label"] = p.AriaLabel
	}
	if strings.TrimSpace(p.Href) != "" && p.Disabled {
		attrs["aria-disabled"] = "true"
		attrs["tabindex"] = "-1"
		if _, ok := attrs["role"]; !ok {
			attrs["role"] = "link"
		}
	}
	return attrs
}

func blockAttrs(p BlockProps) templ.Attributes {
	return domAttrs(DOMProps{ID: p.ID})
}

func buttonType(t string) string {
	if strings.TrimSpace(t) == "" {
		return "button"
	}
	return t
}

func buttonRel(disabled bool) string {
	if disabled {
		return "nofollow noopener noreferrer"
	}
	return ""
}

func fieldClasses(p FieldProps) string {
	if p.Type == "checkbox" || p.Type == "radio" {
		return utils.Cn(utils.FieldControlVariant(p.Variant), utils.FieldControlSizeVariant(p.Size), p.Class)
	}
	return utils.Cn(utils.FieldVariant(p.Variant), utils.FieldSizeVariant(p.Size), p.Class)
}

func fieldAttrs(p FieldProps) templ.Attributes {
	attrs := domAttrs(DOMProps{ID: p.ID, Role: p.Role, TabIndex: p.TabIndex, Attrs: p.Attrs})
	delete(attrs, "id")
	if strings.TrimSpace(p.AriaLabel) != "" {
		attrs["aria-label"] = p.AriaLabel
	}
	if p.Type == "checkbox" && p.Switch {
		attrs["role"] = "switch"
		attrs["aria-checked"] = strconv.FormatBool(p.Checked)
	}
	return attrs
}

func stackClasses(p StackProps) string {
	return utils.Cn("ui-stack", p.Class)
}

func titleTag(order int) string {
	if order < 1 || order > 6 {
		return "h2"
	}
	return fmt.Sprintf("h%d", order)
}

func titleClasses(p TitleProps) string {
	fontSize := p.FontSize
	if fontSize == "" {
		fontSize = "2xl"
	}
	fontWeight := p.FontWeight
	if fontWeight == "" {
		fontWeight = "semibold"
	}
	return utils.Cn(utils.TypographyClasses(fontSize, fontWeight, p.LineHeight, p.LetterSpacing, p.TextColor, p.TextAlign, p.Truncate), p.Class)
}

func textClasses(p TextProps) string {
	return utils.Cn(utils.TypographyClasses(p.FontSize, p.FontWeight, p.LineHeight, p.LetterSpacing, p.TextColor, p.TextAlign, p.Truncate), p.Class)
}

func iconClasses(p IconProps) string {
	size := p.Size
	switch size {
	case "xs":
		size = "ui-icon-xs"
	case "", "sm":
		size = "ui-icon-sm"
	case "md":
		size = "ui-icon-md"
	case "lg":
		size = "ui-icon-lg"
	}
	return utils.Cn("ui-icon", "latty", "latty-"+p.Name, size, p.Class)
}

func fieldRows(rows int) int {
	if rows <= 0 {
		return 4
	}
	return rows
}

func fieldInputType(t string) string {
	if strings.TrimSpace(t) == "" {
		return "text"
	}
	return t
}

func fieldID(p FieldProps) string {
	if strings.TrimSpace(p.ID) != "" {
		return p.ID
	}
	if strings.TrimSpace(p.Name) != "" {
		return p.Name
	}
	return ""
}

func imageLoading(value string) string {
	switch strings.TrimSpace(value) {
	case "eager":
		return "eager"
	default:
		return "lazy"
	}
}

func imageClasses(p ImageProps) string {
	return utils.Cn(
		"ui-image",
		utils.ImageFitVariant(p.Fit),
		utils.ImagePositionVariant(p.Position),
		utils.ImageAspectVariant(p.Aspect),
		p.Class,
	)
}

func sourceAttrs(p SourceProps) templ.Attributes {
	attrs := templ.Attributes{}
	if strings.TrimSpace(p.SrcSet) != "" {
		attrs["srcset"] = p.SrcSet
	}
	if strings.TrimSpace(p.Src) != "" {
		attrs["src"] = p.Src
	}
	if strings.TrimSpace(p.Media) != "" {
		attrs["media"] = p.Media
	}
	if strings.TrimSpace(p.Type) != "" {
		attrs["type"] = p.Type
	}
	if strings.TrimSpace(p.Sizes) != "" {
		attrs["sizes"] = p.Sizes
	}
	return attrs
}

func gridClasses(p GridProps) string {
	return utils.Cn("ui-grid", utils.GridColsVariant(p.Cols), p.Class)
}

func gridColClasses(p GridColProps) string {
	return utils.Cn(utils.GridColVariant(p.Span, p.Start, p.End, p.Order), p.Class)
}

func tableAttrs(p TableProps) templ.Attributes {
	return domAttrs(DOMProps{Attrs: p.Attrs})
}

func tableCellAttrs(p TableCellProps, heading bool) templ.Attributes {
	attrs := templ.Attributes{}
	if heading {
		scope := tableScope(p.Scope)
		if scope != "" {
			attrs["scope"] = scope
		}
		if strings.TrimSpace(p.Abbr) != "" {
			attrs["abbr"] = p.Abbr
		}
	}
	if p.ColSpan > 0 {
		attrs["colspan"] = strconv.Itoa(p.ColSpan)
	}
	if p.RowSpan > 0 {
		attrs["rowspan"] = strconv.Itoa(p.RowSpan)
	}
	if strings.TrimSpace(p.Headers) != "" {
		attrs["headers"] = p.Headers
	}
	return attrs
}

func tableScope(value string) string {
	switch strings.TrimSpace(value) {
	case "row", "col", "rowgroup", "colgroup":
		return strings.TrimSpace(value)
	default:
		return ""
	}
}

func spanAttrs(span int) templ.Attributes {
	if span <= 0 {
		return templ.Attributes{}
	}
	return templ.Attributes{"span": strconv.Itoa(span)}
}

func listTag(tag string) string {
	return utils.ResolveTag(tag, "ul", utils.TagGroupList)
}

func listItemAttrs(p ListItemProps) templ.Attributes {
	if p.Value <= 0 {
		return templ.Attributes{}
	}
	return templ.Attributes{"value": strconv.Itoa(p.Value)}
}

func formAttrs(p FormProps) templ.Attributes {
	attrs := domAttrs(DOMProps{ID: p.ID, Attrs: p.Attrs})
	if strings.TrimSpace(p.Action) != "" {
		attrs["action"] = p.Action
	}
	if strings.TrimSpace(p.Method) != "" {
		attrs["method"] = p.Method
	}
	if strings.TrimSpace(p.Enctype) != "" {
		attrs["enctype"] = p.Enctype
	}
	if strings.TrimSpace(p.Autocomplete) != "" {
		attrs["autocomplete"] = p.Autocomplete
	}
	if strings.TrimSpace(p.Name) != "" {
		attrs["name"] = p.Name
	}
	if strings.TrimSpace(p.Target) != "" {
		attrs["target"] = p.Target
	}
	if p.NoValidate {
		attrs["novalidate"] = true
	}
	return attrs
}

func fieldsetAttrs(p FieldsetProps) templ.Attributes {
	attrs := domAttrs(DOMProps{Attrs: p.Attrs})
	if strings.TrimSpace(p.Name) != "" {
		attrs["name"] = p.Name
	}
	if strings.TrimSpace(p.Form) != "" {
		attrs["form"] = p.Form
	}
	if p.Disabled {
		attrs["disabled"] = true
	}
	return attrs
}

func dataListAttrs(p DataListProps) templ.Attributes {
	return domAttrs(DOMProps{ID: p.ID})
}

func optGroupAttrs(p OptGroupProps) templ.Attributes {
	attrs := templ.Attributes{}
	if strings.TrimSpace(p.Label) != "" {
		attrs["label"] = p.Label
	}
	if p.Disabled {
		attrs["disabled"] = true
	}
	return attrs
}

func outputAttrs(p OutputProps) templ.Attributes {
	attrs := domAttrs(DOMProps{ID: p.ID, Attrs: p.Attrs})
	if strings.TrimSpace(p.Name) != "" {
		attrs["name"] = p.Name
	}
	if strings.TrimSpace(p.For) != "" {
		attrs["for"] = p.For
	}
	return attrs
}

func meterAttrs(p MeterProps) templ.Attributes {
	attrs := domAttrs(DOMProps{ID: p.ID, Attrs: p.Attrs})
	if strings.TrimSpace(p.Value) != "" {
		attrs["value"] = p.Value
	}
	if strings.TrimSpace(p.Min) != "" {
		attrs["min"] = p.Min
	}
	if strings.TrimSpace(p.Max) != "" {
		attrs["max"] = p.Max
	}
	if strings.TrimSpace(p.Low) != "" {
		attrs["low"] = p.Low
	}
	if strings.TrimSpace(p.High) != "" {
		attrs["high"] = p.High
	}
	if strings.TrimSpace(p.Optimum) != "" {
		attrs["optimum"] = p.Optimum
	}
	return attrs
}

func progressAttrs(p ProgressProps) templ.Attributes {
	attrs := domAttrs(DOMProps{ID: p.ID, Attrs: p.Attrs})
	if strings.TrimSpace(p.Value) != "" {
		attrs["value"] = p.Value
	}
	if strings.TrimSpace(p.Max) != "" {
		attrs["max"] = p.Max
	}
	return attrs
}

func intAttr(n int) string {
	return strconv.Itoa(n)
}
