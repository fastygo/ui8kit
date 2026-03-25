package utils

import "strings"

// UtilityProps maps semantic component props onto Tailwind utility classes.
// Embed this struct in any component Props to gain shorthand styling.
type UtilityProps struct {
	Flex     string
	Gap      string
	Items    string
	Justify  string
	P        string
	Px       string
	Py       string
	Pt       string
	Pb       string
	Pl       string
	Pr       string
	M        string
	Mt       string
	Mb       string
	Ml       string
	Mr       string
	Mx       string
	My       string
	W        string
	H        string
	MinW     string
	MaxW     string
	Rounded  string
	Shadow   string
	Border   string
	Bg       string
	Text     string
	Font     string
	Leading  string
	Tracking string
	Overflow string
	Z        string
	Hidden   bool
	Truncate bool
	Grid     string
	Col      string
	Row      string
	Order    string
	Grow     string
	Shrink   string
	Basis    string
}

// Resolve converts populated utility props into a Tailwind class string.
func (u UtilityProps) Resolve() string {
	classes := make([]string, 0, 32)

	if flex := strings.TrimSpace(u.Flex); flex != "" {
		classes = append(classes, resolveFlex(flex)...)
	}
	if gap := strings.TrimSpace(u.Gap); gap != "" {
		classes = append(classes, "gap-"+resolveGap(gap))
	}

	appendIf := func(prefix, value string) {
		if v := strings.TrimSpace(value); v != "" {
			classes = append(classes, prefix+v)
		}
	}

	appendIf("items-", u.Items)
	appendIf("justify-", u.Justify)
	appendIf("p-", u.P)
	appendIf("px-", u.Px)
	appendIf("py-", u.Py)
	appendIf("pt-", u.Pt)
	appendIf("pb-", u.Pb)
	appendIf("pl-", u.Pl)
	appendIf("pr-", u.Pr)
	appendIf("m-", u.M)
	appendIf("mt-", u.Mt)
	appendIf("mb-", u.Mb)
	appendIf("ml-", u.Ml)
	appendIf("mr-", u.Mr)
	appendIf("mx-", u.Mx)
	appendIf("my-", u.My)
	appendIf("w-", u.W)
	appendIf("h-", u.H)
	appendIf("min-w-", u.MinW)
	appendIf("max-w-", u.MaxW)
	appendIf("bg-", u.Bg)
	appendIf("text-", u.Text)
	appendIf("font-", u.Font)
	appendIf("leading-", u.Leading)
	appendIf("tracking-", u.Tracking)
	appendIf("overflow-", u.Overflow)
	appendIf("z-", u.Z)
	appendIf("grid-cols-", u.Grid)
	appendIf("col-span-", u.Col)
	appendIf("row-span-", u.Row)
	appendIf("order-", u.Order)
	appendIf("basis-", u.Basis)

	if rounded := strings.TrimSpace(u.Rounded); rounded != "" {
		if rounded == "default" || rounded == "true" {
			classes = append(classes, "rounded")
		} else {
			classes = append(classes, "rounded-"+rounded)
		}
	}
	if shadow := strings.TrimSpace(u.Shadow); shadow != "" {
		if shadow == "default" || shadow == "true" {
			classes = append(classes, "shadow")
		} else {
			classes = append(classes, "shadow-"+shadow)
		}
	}
	if border := strings.TrimSpace(u.Border); border != "" {
		if border == "default" || border == "true" {
			classes = append(classes, "border")
		} else if strings.HasPrefix(border, "border-") {
			classes = append(classes, border)
		} else {
			classes = append(classes, "border-"+border)
		}
	}
	if grow := strings.TrimSpace(u.Grow); grow != "" {
		if grow == "default" || grow == "true" {
			classes = append(classes, "grow")
		} else {
			classes = append(classes, "grow-"+grow)
		}
	}
	if shrink := strings.TrimSpace(u.Shrink); shrink != "" {
		if shrink == "default" || shrink == "true" {
			classes = append(classes, "shrink")
		} else {
			classes = append(classes, "shrink-"+shrink)
		}
	}

	if u.Hidden {
		classes = append(classes, "hidden")
	}
	if u.Truncate {
		classes = append(classes, "truncate")
	}

	return Cn(classes...)
}

func resolveFlex(value string) []string {
	switch value {
	case "row":
		return []string{"flex", "flex-row"}
	case "col":
		return []string{"flex", "flex-col"}
	case "row-reverse":
		return []string{"flex", "flex-row-reverse"}
	case "col-reverse":
		return []string{"flex", "flex-col-reverse"}
	case "wrap":
		return []string{"flex", "flex-wrap"}
	case "nowrap":
		return []string{"flex", "flex-nowrap"}
	case "inline":
		return []string{"inline-flex"}
	default:
		if strings.HasPrefix(value, "flex") || strings.HasPrefix(value, "inline-flex") {
			return []string{value}
		}
		return []string{"flex", "flex-" + value}
	}
}

func resolveGap(value string) string {
	switch value {
	case "xs":
		return "1"
	case "sm":
		return "2"
	case "md":
		return "4"
	case "lg":
		return "6"
	case "xl":
		return "8"
	default:
		return value
	}
}
