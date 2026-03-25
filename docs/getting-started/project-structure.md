# Recommended project structure

This is a sensible layout for a Go + templ + Tailwind app that consumes UI8Kit. Adjust names to match your repository.

## Layout

```
myapp/
├── go.mod
├── go.sum
├── cmd/
│   └── server/
│       └── main.go          # HTTP server, routes, static files
├── internal/
│   └── ...                  # app logic
├── views/                   # or templates/, pages/, etc.
│   ├── dashboard.templ
│   └── ...
├── static/
│   └── css/
│       ├── input.css        # Tailwind entry (sources + imports)
│       ├── app.css          # generated — do not edit by hand
│       └── ui8kit/          # optional: copied theme files (see Tailwind doc)
│           ├── base.css
│           ├── components.css
│           └── latty.css
├── package.json             # optional: Tailwind CLI scripts
└── node_modules/            # optional: local Tailwind install
```

## Generated files

- `templ generate` produces `*_templ.go` next to each `.templ` file. Many teams gitignore `*_templ.go` and regenerate in CI; others commit them. Pick one policy and document it for your team.

## Static assets

- **Compiled CSS** — `static/css/app.css` is the file your browser loads. Your Shell link should match the URL you serve (see [HTTP server](../integration/http-server.md)).
- **UI8Kit theme CSS** — The library ships `base.css`, `components.css`, and `latty.css` inside the `styles` package. For Tailwind processing, you typically **import copies** of these files from `static/css/ui8kit/` (or import from a known path under the module cache). Details are in [Tailwind setup](../integration/tailwind-setup.md).

## Content scanning

Tailwind must see class names that appear in:

- Your `.templ` and `.go` files
- Any string literals that emit Tailwind classes

Point `@source` (Tailwind v4) at those directories. If you vendor UI8Kit or use a local `replace`, you can also add `@source` paths into the module’s `ui` and `layout` packages so utility classes used inside the library are included in your bundle.
