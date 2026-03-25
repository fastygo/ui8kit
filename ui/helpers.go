package ui

import (
	"fmt"
	"strings"

	"github.com/fastygo/ui8kit/utils"
)

func resolveTag(tag, fallback string) string {
	if t := strings.TrimSpace(tag); t != "" {
		return t
	}
	return fallback
}

func buttonClasses(p ButtonProps) string {
	state := ""
	if p.Disabled {
		state = "pointer-events-none opacity-50"
	}
	return utils.Cn(utils.ButtonStyleVariant(p.Variant), utils.ButtonSizeVariant(p.Size), p.UtilityProps.Resolve(), state, p.Class)
}

func buttonType(t string) string {
	if strings.TrimSpace(t) == "" {
		return "button"
	}
	return t
}

func fieldClasses(p FieldProps) string {
	if p.Type == "checkbox" || p.Type == "radio" {
		return utils.Cn(utils.FieldControlVariant(p.Variant), utils.FieldControlSizeVariant(p.Size), p.UtilityProps.Resolve(), p.Class)
	}
	return utils.Cn(utils.FieldVariant(p.Variant), utils.FieldSizeVariant(p.Size), p.UtilityProps.Resolve(), p.Class)
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
		size = "h-3 w-3"
	case "", "sm":
		size = "h-4 w-4"
	case "md":
		size = "h-5 w-5"
	case "lg":
		size = "h-6 w-6"
	}
	return utils.Cn("latty", "latty-"+p.Name, size, p.Class)
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
