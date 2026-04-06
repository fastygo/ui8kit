# Versioning and releases

## Rules

- `PATCH`: docs, bugfixes, internals.
- `MINOR`: new components/props/variants (backward compatible).
- `MAJOR`: breaking API changes.

Use tags like `vMAJOR.MINOR.PATCH`.

## Release

```bash
./scripts/release.sh 0.2.0
```

The script validates semver, runs tests, updates `Version`, creates commit and tag.

For the new release:

```bash
git push origin main --tags
```

CI checks tag/version consistency, then publishes GitHub release and updates module metadata.
