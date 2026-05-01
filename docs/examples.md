# Reference examples

`tests/examples/` contains small reference compositions for future `Elements` and `Blocks` packages.

They are not product screens. They are standards for how higher layers should start from UI8Kit:

- Use `ui` primitives for layout and controls.
- Use `components` composites for neutral reusable UI.
- Prefer props such as `Tag`, `Variant`, `Size`, `Cols`, `Label`, `Hint`, and `Options`.
- Do not introduce raw HTML tags in example compositions.
- Do not introduce local utility classes, inline styles, brand classes, gradients, or app CSS.

## Element Examples

Element examples represent reusable widgets that can later move into an `Elements` package:

- `ActionCardElement`
- `FormFieldElement`
- `DisclosureElement`
- `DataTableElement`
- `DescriptionListElement`
- `NativeDisclosureElement`
- `MediaElement`
- `FormElement`

These examples compose `Card`, `Field`, `Accordion`, `Table`, `DescriptionList`, `Disclosure`, `Picture`, `Figure`, `Form`, `Text`, `Group`, and `Button` without app-specific styling.

## Block Examples

Block examples represent page sections that can later move into a `Blocks` package:

- `SummaryBlockExample`
- `SettingsBlockExample`
- `DataTableBlockExample`
- `SemanticFormBlockExample`

Blocks should compose existing primitives and elements first. If a block needs a new visual treatment, prefer adding a neutral UI8Kit variant or component policy before adding product-local CSS.

## Validation

Run the same checks used for the kit:

```bash
templ generate ./...
go test ./...
npx ui8px@latest lint ui components utils styles tests/examples
```
