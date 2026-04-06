# Sync UI8Kit CSS

Use this command before CSS build:

```bash
bash ./scripts/sync-ui8kit-css.sh
```

The script copies:

- `base.css`
- `components.css`
- `shell.css`
- `latty.css`

into `static/css/ui8kit/` and prints confirmation.

Then:

```bash
templ generate ./...
npm run sync:ui8kit
npm run build:css
```

Import in `static/css/input.css`:

```css
@import "./ui8kit/base.css";
@import "./ui8kit/components.css";
@import "./ui8kit/shell.css";
@import "./ui8kit/latty.css";
```
