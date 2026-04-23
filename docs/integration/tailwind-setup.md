# Tailwind setup

## Goal

Compile `static/css/app.css` from `static/css/input.css`.

## Required `package.json`

```json
{
  "private": true,
  "type": "module",
  "devDependencies": {
    "@tailwindcss/cli": "^4.2.1",
    "@ui8kit/aria": "0.1.0",
    "tailwindcss": "^4.2.1"
  },
  "scripts": {
    "vendor:assets": "go run github.com/fastygo/ui8kit/scripts/cmd/sync-assets web/static",
    "dev:css": "tailwindcss -i ./static/css/input.css -o ./static/css/app.css --watch",
    "build:css": "tailwindcss -i ./static/css/input.css -o ./static/css/app.css --minify"
  }
}
```

## Setup steps

```bash
bun install
bun run vendor:assets
bun run build:css
```

## `input.css` minimum

```css
@import "tailwindcss";
@import "./ui8kit/base.css";
@import "./ui8kit/components.css";
@import "./ui8kit/shell.css";
@import "./ui8kit/latty.css";

@source "../../**/*.templ";
@source "../../**/*.go";
```

## Helper

Use the sync-assets CLI before build:

```bash
go run github.com/fastygo/ui8kit/scripts/cmd/sync-assets web/static
bun run build:css
```
