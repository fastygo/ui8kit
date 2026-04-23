# Versioning and releases

UI8Kit follows semantic versioning and uses annotated git tags as the source of truth for Go module releases.

## Rules

- `PATCH`: docs, bug fixes, internal cleanups, and non-breaking implementation changes.
- `MINOR`: new components, props, variants, or helpers that remain backward compatible.
- `MAJOR`: breaking API changes.

Use tags in this format:

```text
vMAJOR.MINOR.PATCH
```

Examples:

- `v0.1.1`
- `v0.3.0`
- `v1.0.0`

## Release flow

Run the release script from the repository root (example version `0.3.0`):

```bash
bash ./scripts/preflight.sh
bash ./scripts/release.sh 0.3.0
```

Or, if you prefer to think of it as explicit steps:

```bash
bash ./scripts/preflight.sh
bash ./scripts/release.sh 0.3.0
git push origin main --tags
```

The intended sequence is:

1. `bash ./scripts/preflight.sh` (with Bun available; see [scripts.md](scripts.md)).
2. Commit any generated or formatted changes if preflight produced them.
3. `bash ./scripts/release.sh 0.3.0`
4. `git push origin <branch> --tags` (use your current branch if not `main`)

If you already know the tree is clean and preflight already passed, the release script may be run directly:

```bash
./scripts/release.sh 0.3.0
```

`release.sh`:

1. Validates semver format (no `v` prefix).
2. Ensures the working tree is clean and the tag does not already exist.
3. Runs `PREFLIGHT_REQUIRE_BUN=1 bash ./scripts/preflight.sh`, forwarding `RELEASE_RUN_RACE` to `PREFLIGHT_RUN_RACE` when set.
4. Updates `Version` in `ui8kit.go` via `go run ./scripts/cmd/set-version`.
5. Creates a release commit and an annotated tag.

Then publish:

```bash
git push origin main --tags
```

## Local tests vs CI tests

The local release path is intentionally practical for cross-platform maintainers:

- Preflight runs `go test ./... -count=1` (and `go test ./scripts/cmd/sync-assets` when Bun is required).
- It does **not** require `-race` by default.
- `-race` remains enforced in CI (push, PR, and tag workflows).

This avoids blocking releases on Windows or other environments where `go test -race` requires CGO and a local C toolchain.

If your machine supports CGO and you want race tests during the preflight that `release.sh` invokes:

```bash
RELEASE_RUN_RACE=1 ./scripts/release.sh 0.3.0
```

If CGO is unavailable, preflight skips local race tests and prints a notice. CI still enforces them.

## CI responsibilities

### Push and pull requests (`.github/workflows/ci.yml`)

On each run, GitHub Actions installs Go (matrix), Bun 1.3.x, and templ, then runs:

`PREFLIGHT_RUN_RACE=1 PREFLIGHT_REQUIRE_BUN=1 bash ./scripts/preflight.sh`

That includes `go generate ./...`, formatting checks, `go vet ./...`, `go build ./...`, tests, and sync-assets package tests when Bun is present.

### Tag releases (`.github/workflows/release.yml`)

After `git push origin <branch> --tags` with a new `v*` tag:

1. Verify that the git tag matches `const Version` in `ui8kit.go`.
2. Run the same preflight as CI (`PREFLIGHT_RUN_RACE=1 PREFLIGHT_REQUIRE_BUN=1`).
3. Create a GitHub Release.
4. Trigger Go proxy indexing (`go list -m` against `proxy.golang.org`).

## Important release rules

- Never delete a published tag.
- Never force-push over a tagged commit.
- If a release is bad, publish a new patch version instead of rewriting history.
