# Shell component

`layout.Shell` builds the full dashboard document shell.

## Behavior

- Desktop: fixed sidebar + header + main content.
- Mobile: dialog-backed sheet panel controlled by `@ui8kit/aria`.
- Optional header slot before theme toggle.

## Accessibility

- `aria-label` and `aria-controls` on menu controls.
- `role="dialog"` on mobile panel.

## Extensibility

- Set custom stylesheet path through `ShellProps.CSSPath`.
- Set theme and app script paths through `ShellProps.ThemeJSPath` and `ShellProps.AppJSPath`.
- Pass extra header actions via `HeaderExtra`.

Theme behavior is handled by `theme.js`; interactive layout behavior comes from the vendored `ui8kit.js` bundle.
