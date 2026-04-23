# Contributing

Thanks for your interest in contributing to ui8kit! This is a Go component library built on [templ](https://templ.guide) and Tailwind CSS. Beginners are welcome.

## Prerequisites

- [Go](https://go.dev/dl/) 1.23+
- [templ CLI](https://templ.guide/quick-start/installation): `go install github.com/a-h/templ/cmd/templ@latest`
- [Bun](https://bun.sh/) 1.3+ if you change `scripts/cmd/sync-assets`, layout asset wiring, or run `PREFLIGHT_REQUIRE_BUN=1` / `./scripts/release.sh` (CI uses Bun 1.3.3)

## Quick start

1. Fork this repository and clone your fork:
   ```bash
   git clone <your-fork-url>
   cd ui8kit
   ```
2. Create a branch for your change:
   ```bash
   git checkout -b feat/my-change
   ```
3. Generate templ code, run tests:
   ```bash
   templ generate ./...
   go test ./...
   ```
4. Commit with sign-off (required):
   ```bash
   git add -A
   git commit -s -m "feat: describe your change"
   git push -u origin feat/my-change
   ```
5. Open a Pull Request from your branch to `main`.

## Project structure

```
ui8kit/
├── ui/         UI primitives (Box, Stack, Button, Badge, Field, Icon, ...)
├── layout/     Page shell (Shell, Header, Sidebar)
├── utils/      UtilityProps, Cn(), variant helpers
├── styles/     Embedded CSS (base theme, components, Latty icons)
└── ui8kit.go   Root package, version constant
```

Each package follows the same pattern:

- `props.go` — all Props structs (single source of truth, never redeclared in `.templ`)
- `helpers.go` — internal helper functions
- `*.templ` — templ templates that use the Props and helpers
- `*_test.go` — render tests

## Development workflow

```bash
# Generate Go code from templ files
templ generate ./...

# Build all packages
go build ./...

# Run all tests with race detector
go test ./... -race

# Format code
gofmt -w .

# Vet
go vet ./...
```

## Adding a new component

1. Define the Props struct in the appropriate `props.go` (usually `ui/props.go`).
2. Create a `.templ` file for the component (e.g. `ui/card.templ`).
3. Add helper functions in `helpers.go` if needed.
4. Add render tests in `components_test.go`.
5. Run `templ generate ./...` and `go test ./...`.

## Commit style

Use [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/):

```
feat(button): add loading state
fix(utils): correct border class resolution
docs(readme): add Card component example

Signed-off-by: Your Name <your.email@example.com>
```

## Sign-off policy (DCO)

We require a Developer Certificate of Origin sign-off on all commits.

- **Command line**: use `git commit -s -m "..."`.
- **GitHub UI**: check "Sign off on this commit" when committing.

### Fixing a missing sign-off

```bash
git commit -s --amend --no-edit
git push --force-with-lease
```

## Contributing via GitHub web UI

1. Navigate to the file you want to change and click "Edit".
2. Make your changes.
3. Check "Sign off on this commit" in the commit form.
4. Propose changes and open a Pull Request.

This works well for small fixes (docs, typos). For larger changes, prefer local development.

## Requests for new components

Open a GitHub Issue with:

- Use case description
- Proposed API (Props struct sketch)
- Examples of expected HTML output

## License

By contributing, you agree that your contributions are licensed under the [MIT License](LICENSE) and that you certify the DCO by signing off your commits.
