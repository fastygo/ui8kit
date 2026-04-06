package layout

import "github.com/fastygo/ui8kit/utils"

// MobileSheetCheckboxID toggles the mobile nav sheet (checkbox + label pattern, same idea as ui8kit-core Sheet.tsx).
const MobileSheetCheckboxID = "ui8kit-mobile-sheet"

// MobileSheetPanelID is the dialog surface referenced by aria-controls on the menu trigger.
const MobileSheetPanelID = "ui8kit-mobile-sheet-panel"

func sidebarLinkClass(active, path string) string {
	if active == path {
		return "bg-accent text-accent-foreground"
	}
	return "text-muted-foreground hover:bg-accent"
}

func sidebarItemClasses(active, path string) string {
	return utils.Cn("flex items-center gap-2 rounded px-4 py-2 text-sm", sidebarLinkClass(active, path))
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

func shellLang(value string) string {
	if value == "" {
		return "en"
	}
	return value
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
