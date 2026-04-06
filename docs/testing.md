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
- `ui` and `layout` render output.
- embedded `styles` assets.
- root package version sanity.

## Adding tests

1. Add render/behavior tests for new component paths.
2. Add variant entries when API adds styling behavior.
3. Keep CI command set consistent:

```bash
bash ./scripts/preflight.sh
```
