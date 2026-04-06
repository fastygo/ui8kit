package utils

import "strings"

// CardVariant returns semantic dashboard card classes. Base and modifiers are defined in
// `styles/components.css` (kit-card + kit-card--*) via @apply.
func CardVariant(variant string) string {
	switch strings.TrimSpace(variant) {
	case "", "default", "raised":
		return "kit-card"
	case "kpi":
		return "kit-card kit-card--kpi"
	case "muted":
		return "kit-card kit-card--muted"
	case "ghost":
		return "kit-card kit-card--ghost"
	case "compact":
		return "kit-card kit-card--compact"
	case "flat":
		return "kit-card kit-card--flat"
	case "accent":
		return "kit-card kit-card--accent"
	default:
		return Cn("kit-card", variant)
	}
}

// ButtonStyleVariant returns base + color classes for a button variant.
func ButtonStyleVariant(variant string) string {
	base := "inline-flex items-center justify-center gap-2 whitespace-nowrap text-sm font-medium rounded transition-colors shrink-0 outline-none"
	switch variant {
	case "", "default", "primary":
		return Cn(base, "bg-primary text-primary-foreground hover:opacity-90")
	case "destructive":
		return Cn(base, "bg-destructive text-destructive-foreground hover:opacity-90")
	case "outline":
		return Cn(base, "border border-border bg-background hover:bg-accent hover:text-accent-foreground")
	case "secondary":
		return Cn(base, "bg-secondary text-secondary-foreground hover:opacity-90")
	case "ghost":
		return Cn(base, "hover:bg-accent hover:text-accent-foreground")
	case "link":
		return Cn(base, "text-primary underline-offset-4 hover:underline")
	default:
		return Cn(base, variant)
	}
}

// ButtonSizeVariant returns size classes for a button.
func ButtonSizeVariant(size string) string {
	switch size {
	case "xs":
		return "h-7 px-2 text-xs"
	case "sm":
		return "h-8 px-3 text-sm"
	case "", "default", "md":
		return "h-9 px-4 py-2"
	case "lg":
		return "h-10 px-6 text-base"
	case "xl":
		return "h-11 px-8 text-base"
	case "icon":
		return "h-9 w-9"
	default:
		return size
	}
}

// BadgeStyleVariant returns base + color classes for a badge variant.
func BadgeStyleVariant(variant string) string {
	base := "inline-flex rounded px-2 py-1 text-xs font-medium"
	switch variant {
	case "", "default", "primary", "success":
		return Cn(base, "bg-primary text-primary-foreground")
	case "secondary":
		return Cn(base, "bg-secondary text-secondary-foreground")
	case "destructive":
		return Cn(base, "bg-destructive text-destructive-foreground")
	case "outline":
		return Cn(base, "border border-border bg-background")
	case "warning":
		return Cn(base, "bg-accent text-accent-foreground")
	case "info":
		return Cn(base, "bg-muted text-muted-foreground")
	default:
		return Cn(base, variant)
	}
}

// BadgeSizeVariant returns size classes for a badge.
func BadgeSizeVariant(size string) string {
	switch size {
	case "xs":
		return "px-1.5 py-0.5 text-[10px]"
	case "sm":
		return "px-2 py-0.5 text-xs"
	case "", "default":
		return "px-2 py-1 text-xs"
	case "lg":
		return "px-3 py-1 text-sm"
	default:
		return size
	}
}

// TypographyClasses builds Tailwind typography classes from individual settings.
func TypographyClasses(fontSize, fontWeight, lineHeight, letterSpacing, textColor, textAlign string, truncate bool) string {
	var classes []string
	if fontSize != "" {
		classes = append(classes, "text-"+fontSize)
	}
	if fontWeight != "" {
		classes = append(classes, "font-"+fontWeight)
	}
	if lineHeight != "" {
		classes = append(classes, "leading-"+lineHeight)
	}
	if letterSpacing != "" {
		classes = append(classes, "tracking-"+letterSpacing)
	}
	if textColor != "" {
		classes = append(classes, "text-"+textColor)
	}
	if textAlign != "" {
		classes = append(classes, "text-"+textAlign)
	}
	if truncate {
		classes = append(classes, "truncate")
	}
	return Cn(classes...)
}

// FieldVariant returns base + color classes for an input field.
func FieldVariant(variant string) string {
	base := "w-full rounded border px-3 py-2 text-sm outline-none transition-colors focus-visible:ring-1 focus-visible:ring-ring disabled:opacity-50"
	switch variant {
	case "", "default", "outline":
		return Cn(base, "border-border bg-background")
	case "ghost":
		return Cn(base, "border-transparent bg-muted")
	default:
		return Cn(base, variant)
	}
}

// FieldSizeVariant returns size classes for an input field.
func FieldSizeVariant(size string) string {
	switch size {
	case "xs":
		return "min-h-8 px-2 py-1 text-xs"
	case "sm":
		return "min-h-9 px-2.5 py-1.5 text-sm"
	case "", "default", "md":
		return "min-h-10 px-3 py-2 text-sm"
	case "lg":
		return "min-h-11 px-4 py-3 text-base"
	default:
		return size
	}
}

// FieldControlVariant returns classes for checkbox/radio inputs.
func FieldControlVariant(variant string) string {
	base := "rounded border border-primary text-primary focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:opacity-50"
	switch variant {
	case "", "default", "outline":
		return Cn(base, "border-primary")
	case "ghost":
		return Cn(base, "border-primary/50 bg-muted/30")
	default:
		return Cn(base, variant)
	}
}

// FieldControlSizeVariant returns size classes for checkbox/radio.
func FieldControlSizeVariant(size string) string {
	switch size {
	case "xs":
		return "h-3 w-3"
	case "sm":
		return "h-3.5 w-3.5"
	case "", "default", "md":
		return "h-4 w-4"
	case "lg":
		return "h-5 w-5"
	default:
		return size
	}
}
