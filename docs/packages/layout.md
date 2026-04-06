# Package `layout`

Dashboard page chrome and shell layout.

## Import

```go
import "github.com/fastygo/ui8kit/layout"
```

## What it provides

- `Shell`: full page wrapper and document shell.
- `Sidebar`, `Header`, nav items.
- Desktop and mobile sheet behavior.

## Key props

- `ShellProps` — title, brand, active route, nav, CSS path.
- `SidebarProps` — nav items and active item.
- `HeaderProps` — header title and extra actions.

## Default assumptions

- CSS is at `/static/css/app.css` by default (`CSSPath`).
- Mobile navigation is CSS-only (checkbox + label pattern).

## Best practice

Keep navigation data-driven (`[]layout.NavItem`) and avoid local class sprawl.
