# Package `styles`

Import:

```go
import "github.com/fastygo/ui8kit/styles"
```

## Purpose

Ships the design system CSS as embedded files so you can **serve** or **read** them without maintaining a separate asset repository.

## Embedded files

The following files are included in the binary via `embed.FS`:

| File | Role |
|------|------|
| `base.css` | Theme variables (light/dark), `@theme inline` bridge, base layer |
| `components.css` | Component-layer classes (card, table, form labels, etc.) |
| `latty.css` | Latty icon masks and `.latty-*` classes |

Access:

```go
data, err := styles.FS.ReadFile("base.css")
```

Or serve the whole tree:

```go
http.Handle("/static/css/ui8kit/", http.StripPrefix("/static/css/ui8kit/",
    http.FileServer(http.FS(styles.FS))))
```

## Relationship to Tailwind

These files are written for **Tailwind v4**-style pipelines: they use `@import "tailwindcss"` layering assumptions where applicable (`@layer base`, `@layer components`), and theme tokens under `:root` / `.dark`.

For a **single bundled `app.css`** built with the Tailwind CLI, you normally **import** the same three files from your `input.css` (or copies on disk) so one pass produces one CSS file. See [Tailwind setup](../integration/tailwind-setup.md).

## Dark mode

`base.css` defines `.dark` and uses `@custom-variant dark (&:is(.dark *));` for Tailwind v4. The layout script adds `class="dark"` to the document element when the user selects dark mode.

## Icons

`latty.css` defines CSS variables for SVG masks. Icon components render `<span class="latty latty-{name}">`. Ensure `latty.css` is included in your final CSS if you use `layout` or `ui.Icon`.
