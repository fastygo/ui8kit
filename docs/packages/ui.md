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
    UtilityProps: utils.UtilityProps{P: "2", Rounded: "md"},
    Variant: "primary",
    Size: "sm",
}, "Save")
```

## Styling notes

- Prefer `utils.UtilityProps` and variants instead of raw class strings.
- Dynamic utility values can be added to generated safelist via `scripts/gen-ui8kit-css.go`.

## Quick commands

```bash
templ generate
./scripts/gen-css.sh
```
