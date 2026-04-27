# UI8Kit packages

- `ui` (`github.com/fastygo/ui8kit/ui`): primitives and form controls.
- `components` (`github.com/fastygo/ui8kit/components`): composite components built from `ui` primitives.
- `layout` (`github.com/fastygo/ui8kit/layout`): shell, header, sidebar.
- `utils` (`github.com/fastygo/ui8kit/utils`): class helpers and variants.
- `styles` (`github.com/fastygo/ui8kit/styles`): embedded CSS layers.

Primitives are imported from `ui`; composites are imported from `components`; page shell helpers are imported from `layout`. `styles` is optional but commonly served for theme and icons.

```text
github.com/fastygo/ui8kit/ui -> depends on utils
github.com/fastygo/ui8kit/components -> depends on ui, utils
github.com/fastygo/ui8kit/layout -> depends on utils
github.com/fastygo/ui8kit/styles -> standalone assets
github.com/fastygo/ui8kit (root) -> version constant only
```

Docs: [ui](ui.md), [layout](layout.md), [layout shell](layout-shell.md), [utils](utils.md), [styles](styles.md).

Use [`../examples.md`](../examples.md) and `tests/examples/` for the standard composition style expected from future `Elements` and `Blocks` packages.
