# UI8Kit Go

`github.com/fastygo/ui8kit` is a Go component kit for server-rendered dashboards powered by `templ` + Tailwind.

## Install

```bash
go get github.com/fastygo/ui8kit@latest
go install github.com/a-h/templ/cmd/templ@latest
go build ./...
go test ./...
```

## Packages

- `ui` — primitives (`github.com/fastygo/ui8kit/ui`)
- `layout` — shell and navigation (`github.com/fastygo/ui8kit/layout`)
- `utils` — props, variants, utility composition (`github.com/fastygo/ui8kit/utils`)
- `styles` — embedded CSS assets (`github.com/fastygo/ui8kit/styles`)

## Quick start

```go
import (
    "github.com/fastygo/ui8kit/layout"
    "github.com/fastygo/ui8kit/ui"
)

templ Dashboard(nav []layout.NavItem) {
    @layout.Shell(layout.ShellProps{Title: "HubRelay", Active: "/"}) {
        @ui.Button(ui.ButtonProps{Variant: "primary"}, "Run")
    }
}
```

## CSS flow

```bash
npm install
npm run build:css
./scripts/gen-css.sh
```

Serve either the compiled `static/css/app.css` path from the app, or `styles.FS` in local checks.

See full guides in [docs](docs/README.md).

