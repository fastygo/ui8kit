# Package `ui`

Component layer for dashboard surfaces.

## Import

```go
import "github.com/fastygo/ui8kit/ui"
```

## Components (short)

`Box`, `Stack`, `Group`, `Container`, `Block`, `Button`, `Badge`, `Text`, `Title`, `Field`, `Icon`.

## Typical usage

```go
ui.Button(ui.ButtonProps{
    Class: "rounded-md p-2",
    Variant: "primary",
    Size: "sm",
}, "Save")
```

## Styling notes

- Prefer explicit utility classes in `Class` for local composition.
- Validate explicit classes with `npx ui8px@latest lint ...` from the repository root.

## Quick commands

```bash
templ generate
npx ui8px@latest lint ui components utils styles tests/examples
```
