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

- `ShellProps` — title, brand, active route, nav, CSS path, theme/app JS paths, optional SRI.
- `SidebarProps` — nav items and active item.
- `HeaderProps` — header title and extra actions.

## Default assumptions

- CSS is at `/static/css/app.css` by default (`CSSPath`).
- Theme bootstrap is at `/static/js/theme.js`.
- App behavior bundle is at `/static/js/ui8kit.js`.

## Best practice

Keep navigation data-driven (`[]layout.NavItem`) and avoid local class sprawl.
