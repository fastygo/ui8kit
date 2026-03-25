# Installation

## Go module

Add UI8Kit to your module:

```bash
go get github.com/fastygo/ui8kit@latest
```

Your `go.mod` should contain:

```go
require github.com/fastygo/ui8kit v0.1.0 // version may differ
```

## templ CLI

Install the templ generator (version should be compatible with the `github.com/a-h/templ` version required by UI8Kit):

```bash
go install github.com/a-h/templ/cmd/templ@latest
```

After editing `.templ` files:

```bash
templ generate ./...
```

## Import paths

Use subpackages directly (idiomatic for Go):

```go
import (
    "github.com/fastygo/ui8kit/ui"
    "github.com/fastygo/ui8kit/layout"
    "github.com/fastygo/ui8kit/utils"
    "github.com/fastygo/ui8kit/styles" // optional: embed.FS or ReadFile
)
```

The root package `github.com/fastygo/ui8kit` only exposes a version constant; components live under `ui`, `layout`, etc.

## Optional: Node.js for Tailwind

UI8Kit does not require Node.js at runtime. You only need Node/npm (or another package manager) if you **compile CSS** with the official Tailwind CLI, as described in [Tailwind setup](../integration/tailwind-setup.md).

Minimum versions used in this documentation:

- `tailwindcss` ^4.2.x
- `@tailwindcss/cli` ^4.2.x

## Verify

```bash
go build ./...
go test ./...
```

In your application module, import a component and compile:

```go
// example: unused import check
import _ "github.com/fastygo/ui8kit/ui"
```

Then follow [Project structure](project-structure.md) and [Tailwind setup](../integration/tailwind-setup.md) to add CSS compilation to your app.
