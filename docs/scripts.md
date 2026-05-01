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
- The CLI writes missing app-local `.ui8px/policy` files by default. This lets
  `npx ui8px lint ./...` treat vendored `web/static/css/ui8kit/**` and theme
  token CSS such as `web/static/css/shadcn.css` as control/theme CSS instead of
  app layout CSS. Pass `--ui8px-policy=false` when an application owns its policy
  completely.
- Existing app policy is preserved: `scopes.json`, `denied.json`, and
  `groups.json` are only written when missing. `allowed.json` utilities are
  merged so shared UI8Kit utilities such as `prose` stay available.
  `patterns.json` is merged so UI8Kit-owned `ui-*` patterns stay current while
  app-specific patterns remain.

## `preflight.sh`

### Purpose

Runs the same quality checks locally that should pass in CI before a release or a merge:

- `go generate ./...`
- `gofmt -w .`
- `go vet ./...`
- `go build ./...`
- `go test ./... -count=1`
- `npx ui8px@latest lint ui components utils styles tests/examples` when Bun or npm tooling is available
- `go run ./scripts/cmd/style-patterns --check`
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
- CI and release run preflight with Bun required, so ui8px and semantic pattern checks are expected to be available there.

## `cmd/style-patterns`

### Purpose

Generates `.ui8px/policy/patterns.json` from UI8Kit CSS `@apply` rules and source-only `ui-*` semantic classes.

### Run

```bash
go run ./scripts/cmd/style-patterns
```

Check that the generated policy is current without writing:

```bash
go run ./scripts/cmd/style-patterns --check
```

### Notes

- By default the CLI reads `styles/components.css` and `styles/shell.css`.
- It scans `components`, `ui`, and `layout` for `ui-*` classes that are used as semantic modifiers but do not have their own CSS rule.
- State, pseudo, media, and contextual selectors are intentionally not folded into base patterns.

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
- Regenerate semantic pattern policy with `go run ./scripts/cmd/style-patterns`.
- Use `npx ui8px@latest validate patterns ...` manually when reviewing repeated compositions for possible `ui-*` promotion.
