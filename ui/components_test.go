package ui_test

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/fastygo/ui8kit/ui"
	"github.com/fastygo/ui8kit/utils"
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

func assertNotContains(t *testing.T, html, notWant string) {
	t.Helper()
	if strings.Contains(html, notWant) {
		t.Errorf("expected HTML NOT to contain %q, got:\n%s", notWant, html)
	}
}

func TestButtonPrimary(t *testing.T) {
	btn := ui.Button(ui.ButtonProps{Variant: "primary"}, "Submit")
	html := render(t, btn)
	assertContains(t, html, "Submit")
	assertContains(t, html, "bg-primary")
	assertContains(t, html, "<button")
	assertContains(t, html, `type="button"`)
}

func TestButtonLink(t *testing.T) {
	btn := ui.Button(ui.ButtonProps{Href: "/go"}, "Go")
	html := render(t, btn)
	assertContains(t, html, "<a")
	assertContains(t, html, `href="/go"`)
	assertContains(t, html, "Go")
}

func TestButtonDisabled(t *testing.T) {
	btn := ui.Button(ui.ButtonProps{Disabled: true}, "No")
	html := render(t, btn)
	assertContains(t, html, "disabled")
	assertContains(t, html, "opacity-50")
}

func TestBadgeVariants(t *testing.T) {
	for _, v := range []string{"", "success", "destructive", "warning", "info"} {
		b := ui.Badge(ui.BadgeProps{Variant: v}, "Status")
		html := render(t, b)
		assertContains(t, html, "Status")
		assertContains(t, html, "<span")
	}
}

func TestTextRender(t *testing.T) {
	txt := ui.Text(ui.TextProps{FontSize: "lg", FontWeight: "bold"}, "Hello")
	html := render(t, txt)
	assertContains(t, html, "<p")
	assertContains(t, html, "text-lg")
	assertContains(t, html, "font-bold")
	assertContains(t, html, "Hello")
}

func TestTitleOrders(t *testing.T) {
	for order := 1; order <= 6; order++ {
		title := ui.Title(ui.TitleProps{Order: order}, "Heading")
		html := render(t, title)
		assertContains(t, html, "Heading")
		tag := "h" + string(rune('0'+order))
		assertContains(t, html, "<"+tag)
	}
}

func TestTitleDefaultOrder(t *testing.T) {
	title := ui.Title(ui.TitleProps{}, "Default")
	html := render(t, title)
	assertContains(t, html, "<h2")
}

func TestIconRender(t *testing.T) {
	icon := ui.Icon(ui.IconProps{Name: "sun", Size: "lg"})
	html := render(t, icon)
	assertContains(t, html, "latty-sun")
	assertContains(t, html, "h-6 w-6")
}

func TestFieldInput(t *testing.T) {
	field := ui.Field(ui.FieldProps{
		Name:        "email",
		Type:        "email",
		Placeholder: "you@example.com",
	})
	html := render(t, field)
	assertContains(t, html, "<input")
	assertContains(t, html, `type="email"`)
	assertContains(t, html, `name="email"`)
}

func TestFieldTextarea(t *testing.T) {
	field := ui.Field(ui.FieldProps{Component: "textarea", Rows: 6})
	html := render(t, field)
	assertContains(t, html, "<textarea")
	assertContains(t, html, `rows="6"`)
}

func TestFieldSelect(t *testing.T) {
	field := ui.Field(ui.FieldProps{
		Component: "select",
		Value:     "b",
		Options:   []ui.FieldOption{{Value: "a", Label: "A"}, {Value: "b", Label: "B"}},
	})
	html := render(t, field)
	assertContains(t, html, "<select")
	assertContains(t, html, "<option")
	assertContains(t, html, "selected")
}

func TestBoxWithUtilityProps(t *testing.T) {
	box := ui.Box(ui.BoxProps{
		UtilityProps: utils.UtilityProps{P: "4", Bg: "card", Rounded: "lg"},
		Class:        "extra",
	})
	html := render(t, box)
	assertContains(t, html, "<div")
	assertContains(t, html, "p-4")
	assertContains(t, html, "bg-card")
	assertContains(t, html, "rounded-lg")
	assertContains(t, html, "extra")
}

func TestStackRender(t *testing.T) {
	stack := ui.Stack(ui.StackProps{})
	html := render(t, stack)
	assertContains(t, html, "flex flex-col")
	assertContains(t, html, "gap-4")
}

func TestGroupGrow(t *testing.T) {
	group := ui.Group(ui.GroupProps{Grow: true})
	html := render(t, group)
	assertContains(t, html, "w-full")
}

func TestContainerRender(t *testing.T) {
	c := ui.Container(ui.ContainerProps{})
	html := render(t, c)
	assertContains(t, html, "max-w-7xl")
	assertContains(t, html, "mx-auto")
}
