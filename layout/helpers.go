package layout

import (
	"strings"

	"github.com/fastygo/ui8kit/utils"
)

// MobileSheetTriggerID opens the mobile navigation dialog in the shell header.
const MobileSheetTriggerID = "ui8kit-mobile-sheet-trigger"

// MobileSheetPanelID is the dialog surface referenced by aria-controls on the menu trigger.
const MobileSheetPanelID = "ui8kit-mobile-sheet-panel"

func sidebarItemStateClass(active, path string) string {
	if active == path {
		return "ui-sidebar-item-active"
	}
	return "ui-sidebar-item-inactive"
}

func sidebarItemClasses(active, path string) string {
	return utils.Cn("ui-sidebar-item", sidebarItemStateClass(active, path))
}

func shellBrand(name string) string {
	if name == "" {
		return "App"
	}
	return name
}

func shellCSS(path string) string {
	if path == "" {
		return "/static/css/app.css"
	}
	return path
}

func shellJS(path string) string {
	if path == "" {
		return "/static/js/ui8kit.js"
	}
	return path
}

func shellLang(value string) string {
	if value == "" {
		return "en"
	}
	return value
}

func shellBodyClass(props ShellProps) string {
	if props.MarketingShell {
		return "ui-shell-body ui-shell-body--marketing"
	}
	return "ui-shell-body"
}

func isExternalNavLink(path string) bool {
	return strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://")
}

func themeToggleLabel(value string) string {
	if value == "" {
		return "Toggle theme"
	}
	return value
}

func themeToggleSwitchToDarkLabel(value string) string {
	if value == "" {
		return "Switch to dark theme"
	}
	return value
}

func themeToggleSwitchToLightLabel(value string) string {
	if value == "" {
		return "Switch to light theme"
	}
	return value
}
