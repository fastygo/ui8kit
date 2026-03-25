# Package `layout`

Import:

```go
import "github.com/fastygo/ui8kit/layout"
```

## Purpose

Page-level structure: a full HTML document with sidebar navigation, header, and main content region. Suitable for signed-in dashboards and admin-style UIs.

## Types

- `NavItem` — `Path`, `Label`, `Icon` (Latty icon name without the `latty-` prefix).
- `SidebarProps` — `Items`, `Active` (current path for highlight), `Mobile` (adds close handler on links).
- `HeaderProps` — `Title` (shown in the header bar).
- `ShellProps` — document title, brand, navigation, optional head slot, CSS path.

## Shell

`Shell` renders `<!DOCTYPE html>`, loads CSS (default `/static/css/app.css`), injects small scripts for:

- Theme toggle (`ui8kitToggleTheme`) — toggles `.dark` on `<html>`, persists `ui8kit-theme` in `localStorage`
- Mobile sidebar (`ui8kitOpenSidebar`, `ui8kitCloseSidebar`)

### ShellProps fields

| Field | Description |
|-------|-------------|
| `Title` | `<title>` and main header title |
| `BrandName` | Sidebar brand; default `"App"` if empty |
| `Active` | Path string matched against `NavItem.Path` for active link styling |
| `NavItems` | Slice of `NavItem` |
| `CSSPath` | Stylesheet `href`; default `/static/css/app.css` |
| `HeadExtra` | Optional `templ.Component` appended inside `<head>` (analytics, fonts, etc.) |

Ensure your server actually serves the file at `CSSPath` (see [HTTP server](../integration/http-server.md)).

## Sidebar and Header

You can use `Sidebar` and `Header` outside `Shell` if you compose your own page layout; they are independent templates.

## Accessibility

The header includes buttons with `aria-label` for menu and theme. Extend with landmarks (`nav`, `main`) as needed for your compliance targets.
