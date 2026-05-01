package ui_test

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/a-h/templ"
	"github.com/fastygo/ui8kit/ui"
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
	assertContains(t, html, "ui-button-disabled")
}

func TestButtonUnstyledDoesNotApplyDefaultSize(t *testing.T) {
	btn := ui.Button(ui.ButtonProps{Variant: "unstyled", Class: "ui-header-theme-btn"}, "Theme")
	html := render(t, btn)
	assertContains(t, html, "ui-header-theme-btn")
	assertNotContains(t, html, "h-9")
	assertNotContains(t, html, "px-4")
	assertNotContains(t, html, "py-2")
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
	assertContains(t, html, "ui-icon-lg")
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

func TestBoxWithExplicitClass(t *testing.T) {
	box := ui.Box(ui.BoxProps{
		Class: "p-4 bg-card rounded-lg extra",
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
	assertContains(t, html, "ui-stack")
}

func TestGroupGrow(t *testing.T) {
	group := ui.Group(ui.GroupProps{Grow: true})
	html := render(t, group)
	assertContains(t, html, "ui-group-grow")
}

func TestContainerRender(t *testing.T) {
	c := ui.Container(ui.ContainerProps{})
	html := render(t, c)
	assertContains(t, html, "ui-container")
}

func TestBoxWithTag(t *testing.T) {
	box := ui.Box(ui.BoxProps{Tag: "section"})
	html := render(t, box)
	assertContains(t, html, "<section")
}

func TestBoxInvalidTag(t *testing.T) {
	box := ui.Box(ui.BoxProps{Tag: "table"})
	html := render(t, box)
	assertContains(t, html, "<div")
}

func TestStackWithTag(t *testing.T) {
	stack := ui.Stack(ui.StackProps{Tag: "ul"})
	html := render(t, stack)
	assertContains(t, html, "<ul")
}

func TestGroupWithTag(t *testing.T) {
	group := ui.Group(ui.GroupProps{Tag: "fieldset"})
	html := render(t, group)
	assertContains(t, html, "<fieldset")
}

func TestTextWithTag(t *testing.T) {
	txt := ui.Text(ui.TextProps{Tag: "span"}, "Hello")
	html := render(t, txt)
	assertContains(t, html, "<span")
}

func TestContainerWithTag(t *testing.T) {
	c := ui.Container(ui.ContainerProps{Tag: "section"})
	html := render(t, c)
	assertContains(t, html, "<section")
}

func TestButtonDisabledLink(t *testing.T) {
	btn := ui.Button(ui.ButtonProps{Href: "/x", Disabled: true}, "X")
	html := render(t, btn)
	assertContains(t, html, `aria-disabled="true"`)
	assertContains(t, html, `tabindex="-1"`)
	assertContains(t, html, `role="link"`)
}

func TestButtonBlockWithDOMAttributes(t *testing.T) {
	btn := ui.ButtonBlock(ui.ButtonProps{
		ID:   "trigger",
		Role: "tab",
		Attrs: templ.Attributes{
			"data-tabs-trigger": true,
			"aria-selected":     "true",
		},
		Variant: "unstyled",
	})
	html := render(t, btn)
	assertContains(t, html, `id="trigger"`)
	assertContains(t, html, `role="tab"`)
	assertContains(t, html, `data-tabs-trigger`)
	assertContains(t, html, `aria-selected="true"`)
}

func TestFieldWithLabelAndError(t *testing.T) {
	field := ui.Field(ui.FieldProps{
		ID:    "email",
		Label: "Email",
		Hint:  "Enter email",
		Error: "Required",
	})
	html := render(t, field)
	assertContains(t, html, "<label")
	assertContains(t, html, `for="email"`)
	assertContains(t, html, "Enter email")
	assertContains(t, html, `role="alert"`)
}

func TestFieldSwitch(t *testing.T) {
	field := ui.Field(ui.FieldProps{
		Type:    "checkbox",
		Switch:  true,
		Checked: true,
	})
	html := render(t, field)
	assertContains(t, html, `role="switch"`)
	assertContains(t, html, `aria-checked="true"`)
}

func TestFieldControlWithDOMAttributes(t *testing.T) {
	field := ui.FieldControl(ui.FieldProps{
		ID:   "search",
		Role: "combobox",
		Attrs: templ.Attributes{
			"aria-expanded": "false",
			"aria-controls": "search-options",
		},
		Variant: "unstyled",
		Type:    "text",
	})
	html := render(t, field)
	assertContains(t, html, `id="search"`)
	assertContains(t, html, `role="combobox"`)
	assertContains(t, html, `aria-expanded="false"`)
	assertContains(t, html, `aria-controls="search-options"`)
}

func TestImageRender(t *testing.T) {
	img := ui.Image(ui.ImageProps{
		Src:      "/x.png",
		Alt:      "X",
		Fit:      "cover",
		Position: "center",
		Aspect:   "video",
	})
	html := render(t, img)
	assertContains(t, html, "<img")
	assertContains(t, html, `src="/x.png"`)
	assertContains(t, html, "object-cover")
	assertContains(t, html, "aspect-video")
}

func TestGridRender(t *testing.T) {
	grid := ui.Grid(ui.GridProps{Cols: "3"})
	html := render(t, grid)
	assertContains(t, html, "grid-cols-3")
}

func TestGridColRender(t *testing.T) {
	col := ui.GridCol(ui.GridColProps{Span: 4, Start: 2, Order: 1})
	html := render(t, col)
	assertContains(t, html, "col-span-4")
	assertContains(t, html, "col-start-2")
	assertContains(t, html, "order-1")
}

func TestTablePrimitivesRender(t *testing.T) {
	assertContains(t, render(t, ui.Table(ui.TableProps{})), "<table")
	assertContains(t, render(t, ui.TableCaption(ui.TableCaptionProps{})), "<caption")
	assertContains(t, render(t, ui.TableHead(ui.TableSectionProps{})), "<thead")
	assertContains(t, render(t, ui.TableBody(ui.TableSectionProps{})), "<tbody")
	assertContains(t, render(t, ui.TableFoot(ui.TableSectionProps{})), "<tfoot")
	assertContains(t, render(t, ui.TableRow(ui.TableRowProps{})), "<tr")
	assertContains(t, render(t, ui.TableHeadCell(ui.TableCellProps{Scope: "col"})), `scope="col"`)
	assertContains(t, render(t, ui.TableCell(ui.TableCellProps{ColSpan: 2})), `colspan="2"`)
}

func TestTableColumnsRender(t *testing.T) {
	assertContains(t, render(t, ui.TableColGroup(ui.TableColGroupProps{Span: 2})), "<colgroup")
	assertContains(t, render(t, ui.TableCol(ui.TableColProps{Span: 2})), "<col")
	assertContains(t, render(t, ui.TableCol(ui.TableColProps{Span: 2})), `span="2"`)
}

func TestListPrimitivesRender(t *testing.T) {
	assertContains(t, render(t, ui.List(ui.ListProps{Tag: "ol"})), "<ol")
	assertContains(t, render(t, ui.ListItem(ui.ListItemProps{Value: 2})), "<li")
	assertContains(t, render(t, ui.ListItem(ui.ListItemProps{Value: 2})), `value="2"`)
}

func TestDescriptionListPrimitivesRender(t *testing.T) {
	assertContains(t, render(t, ui.DescriptionList(ui.DescriptionListProps{})), "<dl")
	assertContains(t, render(t, ui.DescriptionTerm(ui.DescriptionTermProps{})), "<dt")
	assertContains(t, render(t, ui.DescriptionDetails(ui.DescriptionDetailsProps{})), "<dd")
}

func TestMediaPrimitivesRender(t *testing.T) {
	assertContains(t, render(t, ui.Picture(ui.PictureProps{})), "<picture")
	sourceHTML := render(t, ui.Source(ui.SourceProps{SrcSet: "/hero.avif", Type: "image/avif"}))
	assertContains(t, sourceHTML, "<source")
	assertContains(t, sourceHTML, `srcset="/hero.avif"`)
}

func TestFigurePrimitivesRender(t *testing.T) {
	assertContains(t, render(t, ui.Figure(ui.FigureProps{})), "<figure")
	assertContains(t, render(t, ui.FigureCaption(ui.FigureCaptionProps{})), "<figcaption")
}

func TestDisclosurePrimitivesRender(t *testing.T) {
	html := render(t, ui.Disclosure(ui.DisclosureProps{Open: true}))
	assertContains(t, html, "<details")
	assertContains(t, html, "open")
	assertContains(t, render(t, ui.DisclosureSummary(ui.DisclosureSummaryProps{})), "<summary")
}

func TestFormPrimitivesRender(t *testing.T) {
	html := render(t, ui.Form(ui.FormProps{ID: "settings", Method: "post", NoValidate: true}))
	assertContains(t, html, "<form")
	assertContains(t, html, `method="post"`)
	assertContains(t, html, "novalidate")
	assertContains(t, render(t, ui.Fieldset(ui.FieldsetProps{Name: "profile"})), "<fieldset")
	assertContains(t, render(t, ui.Legend(ui.LegendProps{})), "<legend")
}

func TestFormValuePrimitivesRender(t *testing.T) {
	output := ui.Output(ui.OutputProps{ID: "total", Name: "total", For: "amount", Value: "42"})
	meter := ui.Meter(ui.MeterProps{Value: "3", Min: "0", Max: "5"})
	progress := ui.Progress(ui.ProgressProps{Value: "2", Max: "10"})

	outputHTML := render(t, output)
	meterHTML := render(t, meter)
	progressHTML := render(t, progress)

	assertContains(t, outputHTML, "<output")
	assertContains(t, outputHTML, `for="amount"`)
	assertContains(t, meterHTML, "<meter")
	assertContains(t, meterHTML, `max="5"`)
	assertContains(t, progressHTML, "<progress")
	assertContains(t, progressHTML, `max="10"`)
}
