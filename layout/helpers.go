package layout

import (
	"context"
	"html"
	"io"
	"strings"

	"github.com/a-h/templ"
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

func shellThemeJS(path string) string {
	if path == "" {
		return "/static/js/theme.js"
	}
	return path
}

func shellAppJS(appPath, legacyPath string) string {
	if appPath != "" {
		return appPath
	}
	if legacyPath != "" {
		return legacyPath
	}
	return "/static/js/ui8kit.js"
}

func shellScriptTag(src, integrity string, deferred bool) templ.Component {
	return templ.ComponentFunc(func(_ context.Context, w io.Writer) error {
		var b strings.Builder
		b.WriteString(`<script src="`)
		b.WriteString(html.EscapeString(src))
		b.WriteString(`"`)
		if deferred {
			b.WriteString(` defer`)
		}
		if integrity != "" {
			b.WriteString(` integrity="`)
			b.WriteString(html.EscapeString(integrity))
			b.WriteString(`" crossorigin="anonymous"`)
		}
		b.WriteString(`></script>`)
		_, err := io.WriteString(w, b.String())
		return err
	})
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
