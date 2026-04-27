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
	return utils.Cn(utils.ButtonStyleVariant(p.Variant), utils.ButtonSizeVariant(p.Size), state, p.Class)
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

func gridClasses(p GridProps) string {
	return utils.Cn("ui-grid", utils.GridColsVariant(p.Cols), p.Class)
}

func gridColClasses(p GridColProps) string {
	return utils.Cn(utils.GridColVariant(p.Span, p.Start, p.End, p.Order), p.Class)
}

func intAttr(n int) string {
	return strconv.Itoa(n)
}
