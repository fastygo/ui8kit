# UI8Kit documentation

Welcome to the UI8Kit documentation. UI8Kit is a **Go + templ + Tailwind CSS** component kit in the style of shadcn/ui.

## Contents

### Getting started

- [Overview](getting-started/README.md) — what UI8Kit is and when to use it
- [Installation](getting-started/installation.md) — Go module, templ CLI, optional Node.js for Tailwind
- [Project structure](getting-started/project-structure.md) — how to lay out an app that uses UI8Kit

### Packages

- [Packages overview](packages/README.md)
- [`ui`](packages/ui.md) — Box, Stack, Group, Container, Button, Badge, Text, Title, Field, Icon
- [`layout`](packages/layout.md) — Shell, Header, Sidebar
- [`utils`](packages/utils.md) — UtilityProps, `Cn`, variant helpers
- [`styles`](packages/styles.md) — embedded CSS, theme tokens, Latty icons

### Integration

- [Integration overview](integration/README.md)
- [Tailwind CSS v4 setup](integration/tailwind-setup.md) — compiler, `input.css`, watch scripts, scanning `.go` / `.templ`
- [HTTP server and static assets](integration/http-server.md) — serving CSS, matching `layout.Shell` defaults

### Testing

- [Testing](testing.md) — test suite overview, architecture, CI integration, writing new tests

### Versioning

- [Versioning and releases](versioning.md) — semver, release script, git tags, Go proxy

## Quick links

| Topic | Document |
|--------|----------|
| `go get` and imports | [Installation](getting-started/installation.md) |
| npm + `@tailwindcss/cli` | [Tailwind setup](integration/tailwind-setup.md) |
| `UtilityProps` | [utils](packages/utils.md) |
| Full-page layout | [layout](packages/layout.md) |

## Module path

```
github.com/fastygo/ui8kit
```

## License

MIT — see the repository [LICENSE](../LICENSE) file.
