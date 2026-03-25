package layout_test

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/fastygo/ui8kit/layout"
)

func render(t *testing.T, c interface{ Render(context.Context, io.Writer) error }) string {
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

func TestHeaderRender(t *testing.T) {
	h := layout.Header(layout.HeaderProps{Title: "My Page"})
	html := render(t, h)
	assertContains(t, html, "My Page")
	assertContains(t, html, "<header")
	assertContains(t, html, "ui8kitToggleTheme")
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
	assertContains(t, html, "bg-accent")
}

func TestSidebarMobile(t *testing.T) {
	items := []layout.NavItem{{Path: "/", Label: "Home", Icon: "box"}}
	s := layout.Sidebar(layout.SidebarProps{Items: items, Active: "/", Mobile: true})
	html := render(t, s)
	assertContains(t, html, "ui8kitCloseSidebar")
}

func TestShellRender(t *testing.T) {
	nav := []layout.NavItem{{Path: "/", Label: "Home", Icon: "box"}}
	sh := layout.Shell(layout.ShellProps{
		Title:     "Test App",
		BrandName: "Brand",
		Active:    "/",
		NavItems:  nav,
	})
	html := render(t, sh)
	assertContains(t, html, "<!doctype html>")
	assertContains(t, html, "<title>Test App</title>")
	assertContains(t, html, "Brand")
	assertContains(t, html, "Home")
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
