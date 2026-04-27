# Package `utils`

Core style helpers for UI8Kit.

## Import

```go
import "github.com/fastygo/ui8kit/utils"
```

## Core API

- `Cn(...)` — join utility classes.
- Variant helpers: `ButtonStyleVariant`, `ButtonSizeVariant`, `BadgeStyleVariant`, `FieldVariant`, etc.

## Usage notes

- Use variants for component states and sizes.
- Use explicit utility classes for local composition and validate them with `ui8px`.
- If a new visual system appears in multiple places, add a reviewed `ui-*` pattern or variant helper.
