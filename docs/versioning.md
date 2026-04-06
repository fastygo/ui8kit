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
- `v0.2.0`
- `v1.0.0`

## Release flow

Run the release script from the repository root:

```bash
bash ./scripts/preflight.sh
bash ./scripts/release.sh 0.2.0
```

Or, if you prefer to think of it as explicit steps:

```bash
bash ./scripts/preflight.sh
bash ./scripts/release.sh 0.2.0
git push origin main --tags
```

The intended sequence is:

1. `bash ./scripts/preflight.sh`
2. Commit any generated or formatted changes if preflight produced them.
3. `bash ./scripts/release.sh 0.2.0`
4. `git push origin main --tags`

If you already know the tree is clean and preflight already passed, the release script may be run directly:

```bash
./scripts/release.sh 0.2.0
```

The script:

1. Validates semver format.
2. Ensures the working tree is clean.
3. Runs `templ generate ./...`.
4. Runs `go test ./... -count=1`.
5. Updates `Version` in `ui8kit.go`.
6. Creates a release commit.
7. Creates an annotated tag.

Then publish:

```bash
git push origin main --tags
```

## Local tests vs CI tests

The local release script is intentionally practical for cross-platform maintainers:

- It always runs `go test ./... -count=1`.
- It does **not** require `-race` by default.
- `-race` remains enforced in CI on the release tag.

This avoids blocking releases on Windows or other environments where `go test -race` requires CGO and a local C toolchain.

If your machine supports CGO and you want the script to run race tests locally too:

```bash
RELEASE_RUN_RACE=1 ./scripts/release.sh 0.2.0
```

If CGO is unavailable, the script will skip local race tests and print a notice. CI still enforces them.

## CI responsibilities

After `git push origin main --tags`, GitHub Actions will:

1. Verify that the git tag matches `const Version` in `ui8kit.go`.
2. Run `templ generate ./...`.
3. Run `go vet ./...`.
4. Run `go test ./... -race -count=1`.
5. Create a GitHub Release.
6. Trigger Go proxy indexing.

## Important release rules

- Never delete a published tag.
- Never force-push over a tagged commit.
- If a release is bad, publish a new patch version instead of rewriting history.
