# UI8Kit Go

A component kit for **Go + [templ](https://templ.guide) + Tailwind CSS** in the style of shadcn/ui design tokens.

## Install

```bash
go get github.com/fastygo/ui8kit@latest
```

## Documentation

Full guides (installation, packages, Tailwind v4 CLI setup for consuming apps) live in the [docs](docs/README.md) folder.

## Packages


| Package  | Import                             | Description                                                                       |
| -------- | ---------------------------------- | --------------------------------------------------------------------------------- |
| `ui`     | `github.com/fastygo/ui8kit/ui`     | Primitives: Box, Stack, Group, Container, Button, Badge, Text, Title, Field, Icon |
| `layout` | `github.com/fastygo/ui8kit/layout` | Shell (sidebar + header + main), Sidebar, Header                                  |
| `utils`  | `github.com/fastygo/ui8kit/utils`  | UtilityProps, Cn(), variant helpers                                               |
| `styles` | `github.com/fastygo/ui8kit/styles` | Embedded CSS (base theme, components, Latty icons)                                |


## Quick start

```go
package pages

import (
    "github.com/fastygo/ui8kit/ui"
    "github.com/fastygo/ui8kit/layout"
)

templ Dashboard(nav []layout.NavItem) {
    @layout.Shell(layout.ShellProps{
        Title:     "Dashboard",
        BrandName: "My App",
        Active:    "/",
        NavItems:  nav,
    }) {
        @ui.Stack(ui.StackProps{}) {
            @ui.Title(ui.TitleProps{Order: 1}, "Welcome")
            @ui.Group(ui.GroupProps{}) {
                @ui.Button(ui.ButtonProps{Variant: "primary"}, "New")
                @ui.Button(ui.ButtonProps{Variant: "outline"}, "Settings")
            }
        }
    }
}
```

## Serving embedded CSS

```go
import (
    "net/http"
    "github.com/fastygo/ui8kit/styles"
)

http.Handle("/static/css/", http.StripPrefix("/static/css/",
    http.FileServer(http.FS(styles.FS))))
```

## License

MIT