# Getting started

UI8Kit is a Go + `templ` component kit with Tailwind in the generated CSS pipeline.

## Requirements

- Go and `templ`
- Bun (for Tailwind build and subset bundling)

## Install

```bash
go get github.com/fastygo/ui8kit@latest
go install github.com/a-h/templ/cmd/templ@latest
```

## Run first steps

```bash
templ generate
```

Then set up CSS:

```bash
bun install
bun run vendor:assets
bun run build:css
```

## Module docs to read

- [Project structure](project-structure.md)
- [Tailwind setup](../integration/tailwind-setup.md)
- [HTTP server](../integration/http-server.md)
