package ui

import (
	"fmt"
	"strconv"
	"strings"

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
	return utils.Cn(utils.ButtonStyleVariant(p.Variant), utils.ButtonSizeVariant(p.Size), p.UtilityProps.Resolve(), state, p.Class)
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
		return utils.Cn(utils.FieldControlVariant(p.Variant), utils.FieldControlSizeVariant(p.Size), p.UtilityProps.Resolve(), p.Class)
	}
	return utils.Cn(utils.FieldVariant(p.Variant), utils.FieldSizeVariant(p.Size), p.UtilityProps.Resolve(), p.Class)
}

func stackClasses(p StackProps) string {
	return utils.Cn("ui-stack", p.UtilityProps.Resolve(), p.Class)
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
	return utils.Cn(utils.TypographyClasses(fontSize, fontWeight, p.LineHeight, p.LetterSpacing, p.TextColor, p.TextAlign, p.Truncate), p.UtilityProps.Resolve(), p.Class)
}

func textClasses(p TextProps) string {
	return utils.Cn(utils.TypographyClasses(p.FontSize, p.FontWeight, p.LineHeight, p.LetterSpacing, p.TextColor, p.TextAlign, p.Truncate), p.UtilityProps.Resolve(), p.Class)
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
		p.UtilityProps.Resolve(),
		p.Class,
	)
}

func gridClasses(p GridProps) string {
	return utils.Cn("ui-grid", utils.GridColsVariant(p.Cols), p.UtilityProps.Resolve(), p.Class)
}

func gridColClasses(p GridColProps) string {
	return utils.Cn(utils.GridColVariant(p.Span, p.Start, p.End, p.Order), p.UtilityProps.Resolve(), p.Class)
}

func cardClasses(p CardProps) string {
	return utils.Cn(utils.CardVariant(p.Variant), p.UtilityProps.Resolve(), p.Class)
}

func cardTag(tag string) string {
	return utils.ResolveTag(tag, "div", utils.TagGroupLayout)
}

func cardTitleTag(order int) string {
	if order < 1 || order > 6 {
		return "h3"
	}
	return fmt.Sprintf("h%d", order)
}

func accordionType(value string) string {
	switch strings.TrimSpace(value) {
	case "multiple":
		return "multiple"
	default:
		return "single"
	}
}

func accordionState(open bool) string {
	if open {
		return "open"
	}
	return "closed"
}

func accordionTriggerID(value string) string {
	return "trigger-" + strings.TrimSpace(value)
}

func accordionPanelID(value string) string {
	return "panel-" + strings.TrimSpace(value)
}

func sheetID(value string) string {
	if strings.TrimSpace(value) == "" {
		return "sheet"
	}
	return value
}

func sheetPanelID(value string) string {
	return sheetID(value) + "-panel"
}

func sheetTitleID(value string) string {
	return sheetID(value) + "-title"
}

func sheetPanelClasses(p SheetProps) string {
	return utils.Cn(
		"ui-sheet-panel",
		utils.SheetSideVariant(p.Side),
		utils.SheetSizeVariant(p.Size),
		p.UtilityProps.Resolve(),
		p.Class,
	)
}

func intAttr(n int) string {
	return strconv.Itoa(n)
}
