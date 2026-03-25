# Getting started

## What is UI8Kit?

UI8Kit provides typed, composable **templ** components that render HTML with **Tailwind CSS** utility classes. It targets server-rendered Go applications: you write `.templ` files, run `templ generate`, and serve HTML. There is no client-side React or Vue runtime.

## What you need

| Requirement | Purpose |
|-------------|---------|
| Go 1.23+ | Build and run your application |
| [templ](https://templ.guide) | Generate Go code from `.templ` templates |
| Tailwind CSS v4 (recommended) | Compile `app.css` from your `input.css` and scanned sources |

UI8Kit itself ships **embedded CSS** (`styles` package) for theme tokens and icon masks. Your app still runs the **Tailwind compiler** so that classes produced by Go code (e.g. `bg-primary`, `flex`, `gap-4`) exist in the final CSS bundle.

## Documentation map

1. [Installation](installation.md) — add the module and tooling.
2. [Project structure](project-structure.md) — folders for `static/`, templates, and CSS entry.
3. [Tailwind setup](../integration/tailwind-setup.md) — `package.json`, `input.css`, watch mode.
4. [HTTP server](../integration/http-server.md) — wire `/static/css/app.css` to your handlers.

## Design assumptions

- You control the HTTP server and static file routes.
- Shell layout defaults to loading CSS from `/static/css/app.css` (configurable via `layout.ShellProps.CSSPath`).
- Dark mode uses a `.dark` class on the root element; layout scripts toggle it and persist preference in `localStorage`.

For a deeper architectural overview, see [DESCRIPTION.md](../../DESCRIPTION.md) in the repository root.
