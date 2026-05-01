package utils

import (
	"strconv"
	"strings"

	"github.com/a-h/templ"
)

func MergeAttrs(groups ...templ.Attributes) templ.Attributes {
	merged := templ.Attributes{}
	for _, group := range groups {
		for key, value := range group {
			merged[key] = value
		}
	}
	return merged
}

func AriaExpanded(expanded bool) templ.Attributes {
	return templ.Attributes{"aria-expanded": strconv.FormatBool(expanded)}
}

func AriaPressed(pressed bool) templ.Attributes {
	return templ.Attributes{"aria-pressed": strconv.FormatBool(pressed)}
}

func AriaChecked(checked bool) templ.Attributes {
	return templ.Attributes{"aria-checked": strconv.FormatBool(checked)}
}

func AriaControls(id string) templ.Attributes {
	if strings.TrimSpace(id) == "" {
		return templ.Attributes{}
	}
	return templ.Attributes{"aria-controls": id}
}

func AriaLabel(label string) templ.Attributes {
	if strings.TrimSpace(label) == "" {
		return templ.Attributes{}
	}
	return templ.Attributes{"aria-label": label}
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

func AriaHasPopup(kind string) templ.Attributes {
	switch strings.TrimSpace(strings.ToLower(kind)) {
	case "true", "menu", "listbox", "tree", "grid", "dialog":
		return templ.Attributes{"aria-haspopup": strings.TrimSpace(strings.ToLower(kind))}
	default:
		return templ.Attributes{}
	}
}

func AriaCurrent(value string) templ.Attributes {
	switch strings.TrimSpace(strings.ToLower(value)) {
	case "page", "step", "location", "date", "time", "true", "false":
		return templ.Attributes{"aria-current": strings.TrimSpace(strings.ToLower(value))}
	default:
		return templ.Attributes{}
	}
}

func AriaLive(mode string) templ.Attributes {
	switch strings.TrimSpace(strings.ToLower(mode)) {
	case "assertive", "off":
		return templ.Attributes{"aria-live": mode}
	default:
		return templ.Attributes{"aria-live": "polite"}
	}
}
