package layout_test

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/a-h/templ"
	"github.com/fastygo/ui8kit/layout"
)

func render(t *testing.T, c interface {
	Render(context.Context, io.Writer) error
}) string {
	t.Helper()
	var buf bytes.Buffer
	if err := c.Render(context.Background(), &buf); err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	return buf.String()
}

func assertContains(t *testing.T, html, want string) {
	t.Helper()
	if !strings.Contains(html, want) {
		t.Errorf("expected HTML to contain %q, got:\n%s", want, html)
	}
}

func assertNotContains(t *testing.T, html, unwanted string) {
	t.Helper()
	if strings.Contains(html, unwanted) {
		t.Errorf("expected HTML not to contain %q", unwanted)
	}
}

func TestHeaderRender(t *testing.T) {
	h := layout.Header(layout.HeaderProps{
		Title: "My Page",
		Extra: templ.ComponentFunc(func(_ context.Context, w io.Writer) error {
			_, err := io.WriteString(w, `<button id="header-extra">extra</button>`)
			return err
		}),
		Trailing: templ.ComponentFunc(func(_ context.Context, w io.Writer) error {
			_, err := io.WriteString(w, `<button id="header-trailing">trailing</button>`)
			return err
		}),
		ThemeToggle: layout.ThemeToggleProps{
			Label:              "Change theme",
			SwitchToDarkLabel:  "Dark mode",
			SwitchToLightLabel: "Light mode",
		},
	})
	html := render(t, h)
	assertContains(t, html, "My Page")
	assertContains(t, html, "<header")
	assertContains(t, html, "ui-header")
	assertContains(t, html, "ui8kit-theme-toggle")
	assertContains(t, html, "aria-controls=\"ui8kit-mobile-sheet-panel\"")
	assertContains(t, html, "data-ui8kit-dialog-open")
	assertContains(t, html, "header-extra")
	assertContains(t, html, "header-trailing")
	assertContains(t, html, "data-switch-to-dark-label=\"Dark mode\"")
	assertContains(t, html, "data-switch-to-light-label=\"Light mode\"")
	assertContains(t, html, "aria-label=\"Change theme\"")
	assertNotContains(t, html, "onclick=\"ui8kitToggleTheme()\"")
}

func TestSidebarRender(t *testing.T) {
	items := []layout.NavItem{
		{Path: "/", Label: "Home", Icon: "box"},
		{Path: "/settings", Label: "Settings", Icon: "boxes"},
	}
	s := layout.Sidebar(layout.SidebarProps{Items: items, Active: "/"})
	html := render(t, s)
	assertContains(t, html, "Home")
	assertContains(t, html, "Settings")
	assertContains(t, html, "latty-box")
	assertContains(t, html, "ui-sidebar-item-active")
}

func TestSidebarMobile(t *testing.T) {
	items := []layout.NavItem{{Path: "/", Label: "Home", Icon: "box"}}
	s := layout.Sidebar(layout.SidebarProps{Items: items, Active: "/", Mobile: true})
	html := render(t, s)
	assertContains(t, html, "Home")
	assertNotContains(t, html, "ui8kitCloseSidebar")
}

func TestShellRender(t *testing.T) {
	nav := []layout.NavItem{{Path: "/", Label: "Home", Icon: "box"}}
	sh := layout.Shell(layout.ShellProps{
		Title:     "Test App",
		Lang:      "ru",
		BrandName: "Brand",
		Active:    "/",
		NavItems:  nav,
		HeaderExtra: templ.ComponentFunc(func(_ context.Context, w io.Writer) error {
			_, err := io.WriteString(w, `<button id="shell-extra">shell</button>`)
			return err
		}),
		HeaderTrailing: templ.ComponentFunc(func(_ context.Context, w io.Writer) error {
			_, err := io.WriteString(w, `<button id="shell-trailing">trail</button>`)
			return err
		}),
		ThemeToggle: layout.ThemeToggleProps{
			Label:              "Theme",
			SwitchToDarkLabel:  "To dark",
			SwitchToLightLabel: "To light",
		},
	})
	html := render(t, sh)
	assertContains(t, html, "<!doctype html>")
	assertContains(t, html, "<html lang=\"ru\">")
	assertContains(t, html, "<title>Test App</title>")
	assertContains(t, html, "Brand")
	assertContains(t, html, "Home")
	assertContains(t, html, "aria-modal=\"true\"")
	assertContains(t, html, "id=\"ui8kit-mobile-sheet-panel\"")
	assertContains(t, html, "data-ui8kit-dialog")
	assertNotContains(t, html, "type=\"checkbox\"")
	assertContains(t, html, "ui-shell-mobile-sheet-overlay")
	assertContains(t, html, "ui-shell-mobile-sheet-panel")
	assertContains(t, html, "/static/js/theme.js")
	assertContains(t, html, "/static/js/ui8kit.js")
	assertNotContains(t, html, "popover=")
	assertNotContains(t, html, "ui8kitOpenSidebar")
	assertNotContains(t, html, "ui8kitToggleTheme")
	assertContains(t, html, "ui-shell-body")
	assertContains(t, html, "ui-shell-main")
	assertContains(t, html, "shell-extra")
	assertContains(t, html, "shell-trailing")
	assertContains(t, html, "data-switch-to-dark-label=\"To dark\"")
}

func TestShellMarketingBodyClass(t *testing.T) {
	sh := layout.Shell(layout.ShellProps{Title: "Marketing", MarketingShell: true})
	html := render(t, sh)
	assertContains(t, html, "ui-shell-body--marketing")
}

func TestShellDefaultBrand(t *testing.T) {
	sh := layout.Shell(layout.ShellProps{Title: "X"})
	html := render(t, sh)
	assertContains(t, html, "App")
}

func TestShellDefaultCSS(t *testing.T) {
	sh := layout.Shell(layout.ShellProps{Title: "X"})
	html := render(t, sh)
	assertContains(t, html, "/static/css/app.css")
}

func TestShellCustomCSS(t *testing.T) {
	sh := layout.Shell(layout.ShellProps{Title: "X", CSSPath: "/assets/style.css"})
	html := render(t, sh)
	assertContains(t, html, "/assets/style.css")
}

func TestShellCustomJS(t *testing.T) {
	sh := layout.Shell(layout.ShellProps{Title: "X", ThemeJSPath: "/assets/theme.min.js", AppJSPath: "/assets/ui8kit.min.js"})
	html := render(t, sh)
	assertContains(t, html, "/assets/theme.min.js")
	assertContains(t, html, "/assets/ui8kit.min.js")
}

func TestShellLegacyJSAlias(t *testing.T) {
	sh := layout.Shell(layout.ShellProps{Title: "X", JSPath: "/assets/legacy-ui8kit.min.js"})
	html := render(t, sh)
	assertContains(t, html, "/assets/legacy-ui8kit.min.js")
}

func TestShellScriptIntegrity(t *testing.T) {
	sh := layout.Shell(layout.ShellProps{
		Title:            "X",
		ThemeJSIntegrity: "sha384-theme",
		AppJSIntegrity:   "sha384-app",
	})
	html := render(t, sh)
	assertContains(t, html, `integrity="sha384-theme" crossorigin="anonymous"`)
	assertContains(t, html, `integrity="sha384-app" crossorigin="anonymous"`)
}
