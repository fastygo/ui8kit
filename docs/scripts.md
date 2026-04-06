# UI8Kit scripts

This document describes the scripts that live in `apps/dashboard/ui8kit/scripts` in this monorepo.

## `gen-ui8kit-css.go`

### Purpose

Generates a concrete utility safelist file from:

- static utility classes from `utils/props.go`
- utility literals used in generated `*_templ.go` files

This helps keep Tailwind CSS generation stable when utility classes are produced from Go structs at compile time.

### How it works

- Parses `utils/props.go` and finds literal utility classes and value patterns.
- Scans all `*_templ.go` files and extracts `.Resolve()` outputs from `UtilityProps` composites.
- Writes CSS declarations into `styles/ui8kit.css` so classes remain detectable by Tailwind.

### Run

```bash
templ generate
go run ./scripts/gen-ui8kit-css.go
```

The default output is `styles/ui8kit.css`.

## `gen-css.sh`

### Purpose

Convenience wrapper that performs both steps:

- runs `templ generate`
- runs `go run ./scripts/gen-ui8kit-css.go`

### Run

```bash
./scripts/gen-css.sh
```

### When to use

Use when you changed:

- `.templ` files that affect utility props usage,
- `utils/props.go`,
- `utils/variants.go`.

## `release.sh`

### Purpose

Maintains tag + release flow for the UI8Kit module.

It is outside the runtime path and is usually used by maintainers.

## Notes

- The generated `styles/ui8kit.css` is intended for development and verification of utility presence.
- It is safe to keep this file in source control as long as your team policy accepts generated style artifacts.
- If you change generator behavior, update this document and the related documentation guides together.
