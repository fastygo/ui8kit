# Sync UI8Kit assets

Use this command before CSS build:

```bash
go run github.com/fastygo/ui8kit/scripts/cmd/sync-assets web/static
```

The CLI copies:

- `base.css`
- `components.css`
- `shell.css`
- `prose.css`
- `latty.css`
- Framework font assets into `web/static/fonts/`
- `theme.js`
- `ui8kit.js` built from `@ui8kit/aria`

into `web/static/` and prints confirmation plus a JS manifest with hashed filenames.

Then:

```bash
templ generate ./...
bun run vendor:assets
bun run build:css
```

Import in `static/css/input.css`:

```css
@import "./ui8kit/base.css";
@import "./ui8kit/components.css";
@import "./ui8kit/shell.css";
@import "./ui8kit/latty.css";
```
