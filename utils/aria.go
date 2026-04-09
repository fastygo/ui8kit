package utils

import (
	"strconv"
	"strings"

	"github.com/a-h/templ"
)

func AriaExpanded(expanded bool) templ.Attributes {
	return templ.Attributes{"aria-expanded": strconv.FormatBool(expanded)}
}

func AriaControls(id string) templ.Attributes {
	if strings.TrimSpace(id) == "" {
		return templ.Attributes{}
	}
	return templ.Attributes{"aria-controls": id}
}

func AriaLabelledBy(id string) templ.Attributes {
	if strings.TrimSpace(id) == "" {
		return templ.Attributes{}
	}
	return templ.Attributes{"aria-labelledby": id}
}

func AriaDescribedBy(id string) templ.Attributes {
	if strings.TrimSpace(id) == "" {
		return templ.Attributes{}
	}
	return templ.Attributes{"aria-describedby": id}
}

func AriaModal(modal bool) templ.Attributes {
	return templ.Attributes{"aria-modal": strconv.FormatBool(modal)}
}

func AriaHidden(hidden bool) templ.Attributes {
	return templ.Attributes{"aria-hidden": strconv.FormatBool(hidden)}
}

func AriaSelected(selected bool) templ.Attributes {
	return templ.Attributes{"aria-selected": strconv.FormatBool(selected)}
}

func AriaDisabled(disabled bool) templ.Attributes {
	return templ.Attributes{"aria-disabled": strconv.FormatBool(disabled)}
}

func AriaRequired(required bool) templ.Attributes {
	return templ.Attributes{"aria-required": strconv.FormatBool(required)}
}

func AriaLive(mode string) templ.Attributes {
	switch strings.TrimSpace(strings.ToLower(mode)) {
	case "assertive", "off":
		return templ.Attributes{"aria-live": mode}
	default:
		return templ.Attributes{"aria-live": "polite"}
	}
}
