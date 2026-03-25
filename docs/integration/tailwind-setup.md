# Tailwind CSS v4 setup for applications using UI8Kit

UI8Kit components emit **Tailwind utility class names** from Go and templ. Those strings are not visible to the Tailwind compiler unless you configure **content scanning** (Tailwind v4: `@source` in CSS) and feed **theme CSS** from UI8Kit into the same pipeline.

This page describes a typical setup using the **official Tailwind CLI** and **npm scripts**.

## Prerequisites

- Node.js 18+ and npm (or pnpm/yarn — adjust commands accordingly)
- Go toolchain and UI8Kit already added to your module ([Installation](../getting-started/installation.md))

## 1. Add Tailwind CLI to your application

In your **application** repository (not inside the UI8Kit module), create or update `package.json`:

```json
{
  "private": true,
  "type": "module",
  "dependencies": {
    "@tailwindcss/cli": "^4.2.1",
    "tailwindcss": "^4.2.1"
  },
  "scripts": {
    "dev:css": "tailwindcss -i ./static/css/input.css -o ./static/css/app.css --watch",
    "build:css": "tailwindcss -i ./static/css/input.css -o ./static/css/app.css --minify"
  }
}
```

Install:

```bash
npm install
```

### Scripts explained

| Script | Purpose |
|--------|---------|
| `dev:css` | Watch `input.css` and source files; rebuild `app.css` on change |
| `build:css` | One-off production build with minification |

Paths assume `static/css/input.css` and `static/css/app.css` at the app root. Change `-i` / `-o` if your layout differs.

## 2. Sync UI8Kit CSS into your static folder

UI8Kit ships three files: `base.css`, `components.css`, `latty.css` (see [styles package](../packages/styles.md)). The Tailwind compiler must process them together with your app’s entry file.

**Recommended:** copy them next to your entry CSS so imports are stable and reproducible:

```bash
mkdir -p static/css/ui8kit
KIT="$(go list -m -f '{{.Dir}}' github.com/fastygo/ui8kit)"
cp "$KIT/styles/base.css" "$KIT/styles/components.css" "$KIT/styles/latty.css" static/css/ui8kit/
```

Re-run this when you upgrade UI8Kit (or automate with `go generate` — see below).

**Local development with `replace`:** if `go.mod` uses `replace github.com/fastygo/ui8kit => ../ui8kit`, the same `go list` command resolves to your local checkout.

## 3. Create `static/css/input.css`

Tailwind v4 uses CSS-first configuration. A practical `input.css` for UI8Kit:

```css
@import "tailwindcss";

/* Scan your templates and Go code for class names (adjust globs to your repo layout) */
@source "../../**/*.templ";
@source "../../**/*.go";

/* UI8Kit theme + components + icons (copied in step 2) */
@import "./ui8kit/base.css";
@import "./ui8kit/components.css";
@import "./ui8kit/latty.css";
```

Notes:

- **`@source`** — Tailwind v4 discovers utilities from these globs. Point at every directory that contains `.templ` / `.go` files with Tailwind classes. The `../../` depth depends on where `static/css/` sits relative to your views; fix the path to match your repo.
- **Library-only classes** — UI8Kit’s own `.go` / `.templ` files live inside the module. If a utility is stripped because it appears only in the dependency, add another `@source` that points at your module cache copy of UI8Kit (path varies by OS and version suffix), vendor the module, or add a small safelist. Most apps scan their **app** templates and rely on imported UI8Kit CSS for layers; add `@source` for UI8Kit when you see missing styles in practice.
- **Import order** — `@import "tailwindcss"` first, then UI8Kit files, so layers and theme tokens apply consistently.

If you prefer not to copy UI8Kit CSS, you can `@import` absolute paths from `go list -m -f '{{.Dir}}'` in a small codegen step that writes a temporary `input.part.css`; the copy approach is simpler for most teams.

## 4. Generate CSS

Development (watch):

```bash
npm run dev:css
```

Production build:

```bash
npm run build:css
```

Commit `app.css` only if your deployment does not run Node; otherwise generate it in CI before building the Go binary.

## 5. `go generate` helper (optional)

Add a generator script to refresh `static/css/ui8kit/*.css` from the module:

```bash
# scripts/sync-ui8kit-css.sh
#!/usr/bin/env bash
set -euo pipefail
mkdir -p static/css/ui8kit
KIT="$(go list -m -f '{{.Dir}}' github.com/fastygo/ui8kit)"
cp "$KIT/styles/base.css" "$KIT/styles/components.css" "$KIT/styles/latty.css" static/css/ui8kit/
```

Wire it from a `//go:generate` line in your main package or a dedicated `tools.go`.

## 6. Fonts (optional)

`base.css` references **Geist** / **Geist Mono** as the default font stack. If you do not load those fonts, the browser falls back to `sans-serif` / `monospace`. Add font links in `layout.ShellProps.HeadExtra` or your base template.

## Troubleshooting

| Symptom | What to check |
|---------|----------------|
| Utilities missing in `app.css` | `@source` globs must cover files where classes appear; rebuild after adding new folders |
| Theme variables not applied | Ensure `base.css` is imported after `@import "tailwindcss"` and that `app.css` is the file the browser loads |
| Dark mode broken | Document must get `class="dark"` on `<html>`; `base.css` defines `.dark` variables |
| Icons invisible | Include `latty.css`; verify `latty` classes are not purged (they appear in `layout` / `ui`) |

## Reference: minimal file tree

```
myapp/
├── package.json
├── static/css/
│   ├── input.css
│   ├── app.css          # generated
│   └── ui8kit/
│       ├── base.css     # copied from UI8Kit
│       ├── components.css
│       └── latty.css
└── views/
    └── *.templ
```

For wiring `app.css` to HTTP, see [HTTP server](http-server.md).
