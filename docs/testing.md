# Testing

## Run tests

```bash
bash ./scripts/preflight.sh

go test ./...

go test ./... -race
```

`preflight.sh` is the recommended local entry point because it also runs generation, formatting, vet, build, and clean-tree checks before tests.

## Scope

- `utils` helpers and variants.
- `utils/tags.go` semantic ownership groups and invalid cross-group cases.
- `ui` and `layout` render output.
- Dedicated HTML5 primitives such as tables, lists, media, disclosure, and form wrappers.
- embedded `styles` assets.
- root package version sanity.

## Adding tests

1. Add render/behavior tests for new component paths.
2. Add tag group tests when a component exposes new semantic ownership.
3. Add variant entries when API adds styling behavior.
4. Keep generated `*_templ.go` current with `templ generate ./...`.
5. Keep CI command set consistent:

```bash
bash ./scripts/preflight.sh
```
