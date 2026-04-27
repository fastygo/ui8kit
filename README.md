# UI8Kit Go

`github.com/fastygo/ui8kit` is a Go component kit for server-rendered dashboards powered by `templ` + Tailwind. Styling is expressed as explicit Tailwind utility classes and validated with `ui8px`.

## Install

```bash
go get github.com/fastygo/ui8kit@latest
go install github.com/a-h/templ/cmd/templ@latest
bun install
go build ./...
go test ./...
```

## Packages

- `ui` — primitives (`github.com/fastygo/ui8kit/ui`)
- `components` — composites built from primitives (`github.com/fastygo/ui8kit/components`)
- `layout` — shell and navigation (`github.com/fastygo/ui8kit/layout`)
- `utils` — class composition, variants, tags, and aria helpers (`github.com/fastygo/ui8kit/utils`)
- `styles` — embedded CSS assets (`github.com/fastygo/ui8kit/styles`)

## Quick start

```go
import (
    cmp "github.com/fastygo/ui8kit/components"
    "github.com/fastygo/ui8kit/layout"
    "github.com/fastygo/ui8kit/ui"
)

templ Dashboard(nav []layout.NavItem) {
    @layout.Shell(layout.ShellProps{Title: "HubRelay", Active: "/"}) {
        @cmp.Card(cmp.CardProps{}) {
            @cmp.CardHeader(cmp.CardHeaderProps{}) {
                @cmp.CardTitle(cmp.CardTitleProps{Order: 2}, "Dashboard")
                @cmp.CardDescription(cmp.CardDescriptionProps{}, "Start with neutral UI8Kit primitives and composites.")
            }
            @cmp.CardContent(cmp.CardContentProps{}) {
                @ui.Button(ui.ButtonProps{Variant: "primary"}, "Run")
            }
        }
    }
}
```

Use `tests/examples/` as the reference for future `Elements` and `Blocks` work: compose `ui` primitives and `components` composites first, without app-specific tags, utility classes, inline styles, or brand CSS.

## CSS flow

```bash
bun install
bun run build:css
npx ui8px@latest lint ui components utils styles tests/examples
```

Keep `.ui8px/` in the repository root. The policy separates compact `control`
files from strict 8px `layout` examples so Tailwind classes remain explicit
while the design-grid rules stay enforceable.

## Asset CLI

UI8Kit ships a Go CLI for vendoring static assets into an application:

```bash
go run github.com/fastygo/ui8kit/scripts/cmd/sync-assets@latest web/static
```

The CLI copies UI8Kit CSS, Framework font assets, emits `theme.js`, and builds
`ui8kit.js` from `@ui8kit/aria` (full or subset mode).

Serve either the compiled `static/css/app.css` path from the app, or `styles.FS` in local checks.

See full guides in [docs](docs/README.md).

