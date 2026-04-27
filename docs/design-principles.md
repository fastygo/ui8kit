# UI8Kit design principles

## Defaults

- 8px/4-8 grid spacing.
- Keep templates structural; avoid inline styles and ad-hoc utility sprawl.
- Use semantic design tokens from `styles/base.css`.
- Reuse existing props and variants before adding local classes.
- Higher-layer examples should start from `ui` primitives and `components` composites with no raw HTML tags or product-local classes.

## Styling workflow

1. Use `utils/props.go` and `utils/variants.go` for reusable style API.
2. If pattern is reusable across packages, add in variants/shared props.
3. Keep app-only overrides in app CSS as `app-*` classes.
4. For new kit-wide patterns use `styles/components.css` or `styles/shell.css` via `@apply`.

## Elements and Blocks baseline

- `tests/examples/elements.templ` shows reusable widget-level composition.
- `tests/examples/blocks.templ` shows section-level composition.
- These examples intentionally avoid app-specific classes and CSS so they can become a clean starting point for external `Elements` and `Blocks` repositories.

## Interaction model

- Prefer CSS-only interactions where possible.
- Keep theme switching and behavior in app scripts.
- Small scripts only for app-specific state.

## Accessibility

- Preserve ARIA on interactive elements.
- Keep keyboard focus/outline visible.
- Keep icon classes consistent through `latty.css`.
