# UI8Kit scripts

This document describes the helper scripts that live in the repository `scripts/` directory.

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

## `preflight.sh`

### Purpose

Runs the same quality checks locally that should pass in CI before a release or a merge:

- `go generate ./...`
- `gofmt -w .`
- `go vet ./...`
- `go build ./...`
- `go test ./... -count=1`
- optional `go test ./... -race -count=1`
- `git diff --exit-code`

### Run

```bash
bash ./scripts/preflight.sh
```

With optional local race tests:

```bash
PREFLIGHT_RUN_RACE=1 bash ./scripts/preflight.sh
```

### Notes

- On systems without CGO support, local race tests are skipped with a warning.
- CI still enforces `go test ./... -race`.
- If `preflight.sh` changes generated or formatted files, it fails on the final `git diff` check so you can review and commit those changes before releasing.

## `release.sh`

### Purpose

Maintains tag + release flow for the UI8Kit module.

It is outside the runtime path and is usually used by maintainers.

### Run

```bash
bash ./scripts/release.sh 0.2.0
```

### Release order

1. Run `bash ./scripts/preflight.sh`.
2. Commit any generated or formatted changes.
3. Run `bash ./scripts/release.sh 0.2.0`.
4. Push `main` and tags: `git push origin main --tags`.

## Notes

- The generated `styles/ui8kit.css` is intended for development and verification of utility presence.
- It is safe to keep this file in source control as long as your team policy accepts generated style artifacts.
- If you change generator behavior, update this document and the related documentation guides together.
