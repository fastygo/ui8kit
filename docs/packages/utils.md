# Package `utils`

Core style helpers for UI8Kit.

## Import

```go
import "github.com/fastygo/ui8kit/utils"
```

## Core API

- `Cn(...)` — join utility classes.
- `UtilityProps` + `Resolve()` — declarative class composition.
- Variant helpers: `ButtonStyleVariant`, `ButtonSizeVariant`, `BadgeStyleVariant`, `FieldVariant`, etc.

## Usage notes

- Use variants / props first.
- If a new visual system appears in multiple places, add it in `variants.go` or `props.go`, then update generator coverage.
