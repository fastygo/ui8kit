# Testing

UI8Kit maintains a comprehensive test suite that validates the utility system, variant helpers, component rendering, layout output, embedded assets, and release integrity.

## Running tests

```bash
go test ./...
```

With race detector (recommended, used in CI):

```bash
go test ./... -race
```

Verbose output to see individual test names:

```bash
go test ./... -v
```

## Test summary

**5 packages, 39 tests, all passing.**

| Package | Test file | Tests | What is covered |
|---------|-----------|-------|-----------------|
| `ui8kit` (root) | `ui8kit_test.go` | 2 | Version constant format and non-empty check |
| `utils` | `cn_test.go` | 5 | `Cn` class joiner: empty, single, multiple, blank filtering, whitespace trimming |
| `utils` | `props_test.go` | 18 | `UtilityProps.Resolve`: every prop category (flex, gap, padding, margin, rounded, shadow, border, grow, shrink, hidden, truncate, grid) |
| `utils` | `variants_test.go` | 7 | All variant helpers: `ButtonStyleVariant` (7 variants + fallback), `ButtonSizeVariant` (7 sizes), `BadgeStyleVariant` (9 variants), `BadgeSizeVariant` (4 sizes), `TypographyClasses`, `FieldVariant`/`FieldSizeVariant`, `FieldControlVariant`/`FieldControlSizeVariant` |
| `ui` | `components_test.go` | 15 | Render tests for every component: Button (primary, link, disabled), Badge (5 variants), Text, Title (orders 1–6, default), Icon, Field (input, textarea, select), Box (with UtilityProps), Stack, Group (grow), Container |
| `layout` | `layout_test.go` | 7 | Render tests for Header, Sidebar (desktop + mobile), Shell (full HTML document, default brand, default CSS path, custom CSS path) |
| `styles` | `embed_test.go` | 1 | Verifies all 3 CSS files (`base.css`, `components.css`, `latty.css`) are embedded and non-empty |

## Test architecture

### Render tests (`ui/`, `layout/`)

Each component is rendered to a `bytes.Buffer` via `templ.Component.Render`, and the resulting HTML is checked for expected content with string assertions:

```go
func render(t *testing.T, c interface{ Render(context.Context, io.Writer) error }) string {
    t.Helper()
    var buf bytes.Buffer
    if err := c.Render(context.Background(), &buf); err != nil {
        t.Fatalf("Render failed: %v", err)
    }
    return buf.String()
}
```

Tests assert:

- Correct HTML tags (`<button`, `<a`, `<h1`–`<h6>`, `<input`, `<textarea`, `<select`)
- Expected CSS classes (`bg-primary`, `flex flex-col`, `gap-4`, `opacity-50`)
- Attribute presence (`disabled`, `type="email"`, `selected`, `href`)
- Content (`label` text, `title`, `brand name`)

### Unit tests (`utils/`)

Table-driven tests verify pure functions with exact string comparison:

```go
{name: "multiple fields", props: UtilityProps{P: "4", Mx: "auto", Bg: "card"}, want: "p-4 mx-auto bg-card"},
```

Variant tests iterate over all documented variant names and assert non-empty output. Fallback behavior (unknown variant passed through as raw class) is also tested.

### Embed tests (`styles/`)

Reads each file from `embed.FS` and verifies it exists and is not empty. Catches accidental removal of CSS files or broken `//go:embed` directives.

### Version tests (root)

Regex-validates that `const Version` matches semver `X.Y.Z` format. Guards against accidental edits that would break release tooling.

## What each layer verifies

```
┌─────────────────────────────────────────────────┐
│  Release integrity                              │
│  ui8kit_test.go: Version format                 │
├─────────────────────────────────────────────────┤
│  Component rendering                            │
│  ui/components_test.go: HTML output of each     │
│  component with various props and variants      │
├─────────────────────────────────────────────────┤
│  Layout rendering                               │
│  layout/layout_test.go: Shell document          │
│  structure, Header, Sidebar (desktop + mobile)  │
├─────────────────────────────────────────────────┤
│  Embedded assets                                │
│  styles/embed_test.go: CSS files present        │
├─────────────────────────────────────────────────┤
│  Utility system                                 │
│  utils/cn_test.go: class joining                │
│  utils/props_test.go: UtilityProps → classes    │
│  utils/variants_test.go: variant → classes      │
└─────────────────────────────────────────────────┘
```

## CI integration

Tests run automatically in two contexts:

1. **On every push/PR to `main`** (`ci.yml`) — matrix of Go 1.23 and 1.24, with `-race` flag.
2. **On tag push** (`release.yml`) — full suite must pass before a GitHub Release is created.

See [Versioning](versioning.md) for how tests gate the release pipeline.

## Writing new tests

When adding a component:

1. Add render tests in `ui/components_test.go` (or `layout/layout_test.go` for layout components).
2. Use the existing `render` and `assertContains` helpers.
3. Test at minimum: correct HTML tag, expected classes for default variant, and at least one non-default variant or prop combination.
4. If adding a new variant helper in `utils/variants.go`, add cases to `variants_test.go`.

Run before committing:

```bash
go test ./... -race
go vet ./...
```
