# Testing

## Run tests

```bash
go test ./...

go test ./... -race
```

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
go test ./... -race

go vet ./...
```
