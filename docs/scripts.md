# UI8Kit scripts

This document describes the helper scripts that live in the repository `scripts/` directory.

## `cmd/sync-assets`

### Purpose

Vendors UI8Kit CSS, Framework font assets, `theme.js`, and a generated
`ui8kit.js` bundle into an application's `web/static/` directory.

### Run

```bash
go run github.com/fastygo/ui8kit/scripts/cmd/sync-assets web/static
```

### Notes

- By default the CLI builds a **subset** IIFE from `@ui8kit/aria` using Bun.
- `theme.js` is emitted separately for first-paint theme bootstrap.
- `ui8kit.js` contains the ARIA bundle plus locale behavior.

## `preflight.sh`

### Purpose

Runs the same quality checks locally that should pass in CI before a release or a merge:

- `go generate ./...`
- `gofmt -w .`
- `go vet ./...`
- `go build ./...`
- `go test ./... -count=1`
- when `PREFLIGHT_REQUIRE_BUN=1`: require Bun on PATH, then `go test ./scripts/cmd/sync-assets -count=1`
- optional `go test ./... -race -count=1`
- `git diff` / `git diff --cached` against a clean tree

### Run

```bash
bash ./scripts/preflight.sh
```

With optional local race tests:

```bash
PREFLIGHT_RUN_RACE=1 bash ./scripts/preflight.sh
```

Match CI and release scripts (Bun required for sync-assets tests):

```bash
PREFLIGHT_REQUIRE_BUN=1 bash ./scripts/preflight.sh
```

### Notes

- On systems without CGO support, local race tests are skipped with a warning.
- CI still enforces `go test ./... -race`.
- If `preflight.sh` changes generated or formatted files, it fails on the final `git diff` check so you can review and commit those changes before releasing.

## `release.sh`

### Purpose

Maintains tag + release flow for the UI8Kit module.

It is outside the runtime path and is usually used by maintainers.

The script runs `preflight.sh` with `PREFLIGHT_REQUIRE_BUN=1` (subset sync-assets tests), then bumps `ui8kit.go`, commits, and tags.

### Run

```bash
bash ./scripts/release.sh 0.3.0
```

### Release order

1. Run `bash ./scripts/preflight.sh` (Bun recommended so local checks match CI).
2. Commit any generated or formatted changes.
3. Run `bash ./scripts/release.sh 0.3.0`.
4. Push your branch and tags: `git push origin main --tags` (or replace `main` with your release branch).

## Notes

- Utility classes stay explicit in `.templ`, `.go`, and CSS `@apply` source files.
- Validate them with `npx ui8px@latest lint ui components utils styles tests/examples`.
- Use `npx ui8px@latest validate patterns ...` manually when reviewing repeated compositions for possible `ui-*` promotion.
