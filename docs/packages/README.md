# UI8Kit packages

- `ui` (`github.com/fastygo/ui8kit/ui`): primitives and form controls.
- `layout` (`github.com/fastygo/ui8kit/layout`): shell, header, sidebar.
- `utils` (`github.com/fastygo/ui8kit/utils`): class helpers and variants.
- `styles` (`github.com/fastygo/ui8kit/styles`): embedded CSS layers.

All components are imported from `ui`/`layout`; `styles` is optional but commonly served for theme and icons.

```text
github.com/fastygo/ui8kit/ui -> depends on utils
github.com/fastygo/ui8kit/layout -> depends on utils
github.com/fastygo/ui8kit/styles -> standalone assets
github.com/fastygo/ui8kit (root) -> version constant only
```

Docs: [ui](ui.md), [layout](layout.md), [layout shell](layout-shell.md), [utils](utils.md), [styles](styles.md).
