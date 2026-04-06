# Tailwind setup

## Goal

Compile `static/css/app.css` from `static/css/input.css`.

## Required `package.json`

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

## Setup steps

```bash
npm install
npm run build:css
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

Use sync helper before build:

```bash
bash ./scripts/sync-ui8kit-css.sh
npm run build:css
```
