# Versioning and releases

UI8Kit follows [Semantic Versioning](https://semver.org/) and uses annotated git tags as the single source of truth for Go module versions.

## Version scheme

```
v{MAJOR}.{MINOR}.{PATCH}
```

| Increment | When |
|-----------|------|
| PATCH | Bug fixes, documentation, internal refactors (no API change) |
| MINOR | New components, new Props fields, new variant values (backward compatible) |
| MAJOR | Breaking changes: removed/renamed types, changed function signatures |

While the project is at `v0.x.y`, the API is considered unstable. Breaking changes may occur in MINOR bumps.

## How to create a release (maintainers)

Use the release script from the repository root:

```bash
./scripts/release.sh 0.2.0
```

The script:

1. Validates semver format.
2. Checks that the working tree is clean.
3. Warns if not on `main`.
4. Runs `go test ./... -race`.
5. Updates `Version` in `ui8kit.go` to match the new version.
6. Creates a commit: `chore: release v0.2.0`.
7. Creates an annotated tag: `v0.2.0`.

Then push:

```bash
git push origin main --tags
```

CI will:

- Verify that the git tag matches `const Version` in `ui8kit.go`.
- Run `templ generate`, `go vet`, and `go test -race`.
- Create a GitHub Release with auto-generated changelog.
- Trigger Go proxy indexing so `go get` resolves the new version immediately.

## Version constant

The `Version` constant in `ui8kit.go` always reflects the latest release:

```go
import "github.com/fastygo/ui8kit"

fmt.Println(ui8kit.Version) // "0.2.0"
```

A unit test enforces that `Version` is valid semver. The release workflow verifies that the git tag and `Version` constant match — a mismatch fails the pipeline.

## Rules for published tags

- **Never delete** a published tag. Go proxy caches are immutable.
- **Never force-push** over a tagged commit. Checksums in `go.sum` will break.
- If a tag has a bug, release a new PATCH (e.g. `v0.1.0` is broken → release `v0.1.1`).
- To deprecate a version, use Go's `retract` directive in `go.mod`:

```go
retract v0.1.0 // contains broken import paths
```

## For consumers

```bash
# Install latest
go get github.com/fastygo/ui8kit@latest

# Pin to specific version
go get github.com/fastygo/ui8kit@v0.2.0

# Upgrade
go get -u github.com/fastygo/ui8kit
```
