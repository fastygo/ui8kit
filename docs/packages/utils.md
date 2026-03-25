# Package `utils`

Import:

```go
import "github.com/fastygo/ui8kit/utils"
```

## Purpose

Shared building blocks for CSS class strings: merging fragments, mapping semantic props to Tailwind utilities, and applying consistent variants for buttons, badges, fields, and typography.

## `Cn`

`Cn` concatenates non-empty class strings with a single space:

```go
utils.Cn("px-4", "", "py-2", "bg-card") // "px-4 py-2 bg-card"
```

## `UtilityProps`

Struct fields mirror common Tailwind prefixes. Calling `Resolve()` returns a single class string.

Examples:

| Fields | Output (illustrative) |
|--------|------------------------|
| `P: "4"` | `p-4` |
| `Flex: "col"` | `flex flex-col` |
| `Gap: "lg"` | `gap-6` (semantic gap scale) |
| `Rounded: "lg"` | `rounded-lg` |
| `Rounded: "default"` | `rounded` |
| `Border: "true"` | `border` |
| `Hidden: true` | `hidden` |

Embed `UtilityProps` in your own structs if you build custom components that should stay consistent with UI8Kit.

## Variant helpers

| Function | Used for |
|----------|----------|
| `ButtonStyleVariant`, `ButtonSizeVariant` | Button appearance |
| `BadgeStyleVariant`, `BadgeSizeVariant` | Badge appearance |
| `FieldVariant`, `FieldSizeVariant` | Text inputs |
| `FieldControlVariant`, `FieldControlSizeVariant` | Checkbox / radio |
| `TypographyClasses` | Text and title |

These return Tailwind class strings; they are plain Go and safe to call from your own code.

## Extensibility

Passing an unknown `Variant` string often appends it as raw classes (see implementation in `variants.go`). Prefer documented variant names for stable upgrades.
