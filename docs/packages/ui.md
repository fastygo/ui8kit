# Package `ui`

Component layer for dashboard surfaces.

## Import

```go
import "github.com/fastygo/ui8kit/ui"
```

## Components (short)

Layout and text:
`Box`, `Stack`, `Group`, `Container`, `Block`, `Grid`, `GridCol`, `Text`, `Title`.

Controls and media:
`Button`, `Badge`, `Field`, `Icon`, `Image`, `Picture`, `Source`, `Figure`, `FigureCaption`.

Semantic structures:
`Table`, `TableCaption`, `TableHead`, `TableBody`, `TableFoot`, `TableRow`, `TableHeadCell`, `TableCell`, `TableColGroup`, `TableCol`, `List`, `ListItem`, `DescriptionList`, `DescriptionTerm`, `DescriptionDetails`, `Disclosure`, `DisclosureSummary`, `Form`, `Fieldset`, `Legend`, `DataList`, `OptGroup`, `Output`, `Meter`, `Progress`.

## Typical usage

```go
ui.Button(ui.ButtonProps{
    Class: "rounded-md p-2",
    Variant: "primary",
    Size: "sm",
}, "Save")
```

## Semantic tables

Tables use dedicated primitives instead of `Box{Tag: "table"}` because table content models are strict:

```go
ui.Table(ui.TableProps{}) {
    ui.TableHead(ui.TableSectionProps{}) {
        ui.TableRow(ui.TableRowProps{}) {
            ui.TableHeadCell(ui.TableCellProps{Scope: "col"}) { "Name" }
        }
    }
    ui.TableBody(ui.TableSectionProps{}) {
        ui.TableRow(ui.TableRowProps{}) {
            ui.TableCell(ui.TableCellProps{}) { "Acme" }
        }
    }
}
```

## Lists, media, disclosure, and forms

- Use `List` + `ListItem` for `ul`, `ol`, and `menu` content.
- Use `DescriptionList`, `DescriptionTerm`, and `DescriptionDetails` for `dl` / `dt` / `dd`.
- Use `Picture`, `Source`, `Image`, `Figure`, and `FigureCaption` for media composition.
- Use `Disclosure` and `DisclosureSummary` for native zero-JS `details` / `summary`.
- Use `Form`, `Fieldset`, and `Legend` around `Field` for form document semantics.

## Tag ownership

Generic layout components intentionally do not accept every HTML5 tag. `Box{Tag: "table"}` remains invalid. Tags with strict content models are owned by dedicated primitives so app templates can stay raw-tag-free and still validate as HTML5.

## Styling notes

- Prefer explicit utility classes in `Class` for local composition.
- Validate explicit classes with `npx ui8px@latest lint ...` from the repository root.

## Quick commands

```bash
templ generate
npx ui8px@latest lint ui components utils styles tests/examples
```
