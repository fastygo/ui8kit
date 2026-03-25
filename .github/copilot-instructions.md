# Coding standards

## Project

ui8kit is a Go component library built on top of templ and Tailwind CSS in the style of shadcn/ui design tokens.

Module path: `github.com/fastygo/ui8kit`

## Behaviour

* Always run `go fmt` after making changes to Go code.
* Always run `go test ./...` after making changes.
* Run `templ generate ./...` after modifying `.templ` files.

## Architecture

The project is organized into four packages:

* `ui/` — UI primitives: Box, Stack, Group, Container, Button, Badge, Text, Title, Field, Icon.
* `layout/` — Page shell: Shell (sidebar + header + main), Header, Sidebar.
* `utils/` — UtilityProps, Cn(), variant helpers for Tailwind classes.
* `styles/` — Embedded CSS via `embed.FS` (base theme, components, Latty icons).

### Key rules

* Props are defined once in `props.go` — never redeclare in `.templ` files.
* Helper functions live in `helpers.go` within each package.
* The only external dependency is `github.com/a-h/templ`.
* Do not add application-specific logic to library components.

## Commit messages

The project uses https://www.conventionalcommits.org/en/v1.0.0/

Examples:

* `feat: add Card component with header and content slots`
* `fix: resolve class duplication in UtilityProps.Resolve`

## Writing style

* Use American English spelling to match the Go standard library, e.g. "color".
* Avoid use of emojis in code, comments, commit messages, and documentation.
* Be precise — avoid filler words like "just" or "really".

## Coding style

* Prefer early returns over `else` blocks.
* Use line breaks to separate "paragraphs" of code.
* Do not write comments that explain what the code does — comments explain why.
* All comments in code must be in English.

## Tests

* Utils tests are in `utils/` — `cn_test.go`, `props_test.go`, `variants_test.go`.
* Component render tests are in `ui/components_test.go`.
* Layout render tests are in `layout/layout_test.go`.
* Embedded CSS tests are in `styles/embed_test.go`.
* Use `go test ./...` to run all tests.

## Moving and renaming files

* templ files have the `.templ` extension.
* Generated files have the `_templ.go` suffix and are excluded from version control via `.gitignore`.
* If a `.templ` file is renamed, moved, or deleted — regenerate with `templ generate ./...`.
