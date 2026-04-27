# UI8Kit - Technical Overview

## What is UI8Kit

UI8Kit is a component kit for **Go + templ + Tailwind CSS** in the style of shadcn/ui design tokens. It provides a set of composable, type-safe UI primitives that render server-side HTML with Tailwind utility classes.

Module: `github.com/fastygo/ui8kit`

## Core Principles

### 1. Server-Side Rendering
All components render to HTML on the server via Go's templ engine. No JavaScript framework, no hydration overhead, no virtual DOM. The browser receives clean, semantic HTML.

### 2. Single Dependency
The only external dependency is `github.com/a-h/templ`. No routers, no ORMs, no CSS-in-JS runtimes. Consumers bring their own HTTP stack.

### 3. Type-Safe Props
Every component accepts a typed Props struct for behavior, variants, semantic tags, and explicit class extension. Compile-time validation, zero reflection, zero interface{}.

### 4. shadcn/ui Design Tokens
CSS variables follow the shadcn/ui design tokens convention (`--primary`, `--destructive`, `--muted`, etc.) with light and dark themes. Consumers can override tokens to match their brand.

## Architecture

```
Consumer App
    │
    ├── import "github.com/fastygo/ui8kit/ui"         ← primitives
    ├── import "github.com/fastygo/ui8kit/components" ← composites
    ├── import "github.com/fastygo/ui8kit/layout"     ← page shell
    ├── import "github.com/fastygo/ui8kit/utils"      ← class utilities
    └── import "github.com/fastygo/ui8kit/styles"     ← embedded CSS
```

### Package Responsibilities

| Package | Purpose |
|---------|---------|
| `ui` | Visual primitives: Box, Stack, Group, Container, Button, Badge, Text, Title, Field, Icon |
| `components` | Composite components: Card, Accordion, Sheet, Dialog, Alert, Breadcrumb, Tabs, Combobox, Tooltip |
| `layout` | Page structure: Shell (sidebar + header + main), Header, Sidebar |
| `utils` | Cn (class joiner), variant helpers, tag helpers, aria helpers |
| `styles` | Embedded CSS via `embed.FS`: base theme, component classes, Latty icon font |

### Component Pattern

Each component follows the same structure:

```
props.go     →  type ButtonProps struct { ... }     (single source of truth)
helpers.go   →  func buttonClasses(p) string        (internal logic)
button.templ →  templ Button(props, label) { ... }  (template)
```

Props are never redeclared in `.templ` files. Generated `_templ.go` files are excluded from version control.

### Explicit Class Policy

Component APIs keep Tailwind utility classes visible to the source scanner:

```go
ui.Box(ui.BoxProps{
    Class: "flex flex-col gap-4 rounded-lg bg-card p-4 shadow",
})
```

The `ui8px` CLI validates classes against the repository policy so layout files keep an 8px spacing rhythm while compact controls can use finer 4px steps.

### Variant System

Components with visual variants (Button, Badge, Field) use helper functions:

```go
utils.ButtonStyleVariant("primary")    → "... bg-primary text-primary-foreground ..."
utils.ButtonStyleVariant("outline")    → "... border border-border bg-background ..."
utils.ButtonSizeVariant("sm")          → "h-8 px-3 text-sm"
```

Variants accept arbitrary strings as fallback, enabling custom class injection.

## CSS Strategy

CSS is embedded in the Go binary via `embed.FS` and served directly:

```go
http.Handle("/static/css/", http.StripPrefix("/static/css/",
    http.FileServer(http.FS(styles.FS))))
```

Three CSS files:

- `base.css` — theme tokens (light + dark), base layer resets
- `components.css` — reusable component classes (card, table, form-label)
- `latty.css` — CSS-only icon system using SVG masks

## Design Decisions

**Why templ over html/template?** Type safety, composability, IDE support, and compile-time error checking. templ components are Go functions that return `templ.Component`.

**Why Tailwind over custom CSS?** Utility-first approach aligns with component-level styling. Explicit class strings keep Tailwind builds minimal, while `ui8px` enforces the 8px design-grid policy.

**Why embed.FS?** Single `go get` installs everything for CSS primitives. Consumers serve CSS from Go binary or copy it into their own static directory.

**Why a separate JS CLI?** Theme bootstrap stays tiny and embedded in `js/theme.js`, while interactive ARIA behavior comes from the published `@ui8kit/aria` package and is vendored into apps through `scripts/cmd/sync-assets`.
