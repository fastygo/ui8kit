package components_test

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/fastygo/ui8kit/components"
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

func TestCardRender(t *testing.T) {
	card := components.Card(components.CardProps{Variant: "accent", Tag: "article"})
	html := render(t, card)
	assertContains(t, html, "<article")
	assertContains(t, html, "ui-card--accent")
}

func TestCardTitle(t *testing.T) {
	title := components.CardTitle(components.CardTitleProps{Order: 2}, "Card")
	html := render(t, title)
	assertContains(t, html, "<h2")
	assertContains(t, html, "Card")
}

func TestAccordionRender(t *testing.T) {
	accordion := components.Accordion(components.AccordionProps{Type: "multiple"})
	html := render(t, accordion)
	assertContains(t, html, `data-ui8kit="accordion"`)
	assertContains(t, html, `data-accordion-type="multiple"`)
}

func TestAccordionParts(t *testing.T) {
	trigger := components.AccordionTrigger(components.AccordionTriggerProps{Value: "a", Open: true})
	content := components.AccordionContent(components.AccordionContentProps{Value: "a", Open: true})
	triggerHTML := render(t, trigger)
	contentHTML := render(t, content)
	assertContains(t, triggerHTML, `aria-controls="ui8kit-accordion-panel-a"`)
	assertContains(t, triggerHTML, `aria-expanded="true"`)
	assertContains(t, contentHTML, `role="region"`)
	assertContains(t, contentHTML, `aria-labelledby="ui8kit-accordion-trigger-a"`)
}

func TestSheetRender(t *testing.T) {
	sheet := components.Sheet(components.SheetProps{ID: "menu", Side: "left", Size: "lg", Title: "Menu"})
	html := render(t, sheet)
	assertContains(t, html, `data-ui8kit="sheet"`)
	assertContains(t, html, `role="dialog"`)
	assertContains(t, html, `aria-modal="true"`)
	assertContains(t, html, `id="menu-panel"`)
	assertContains(t, html, `aria-labelledby="menu-title"`)
}
