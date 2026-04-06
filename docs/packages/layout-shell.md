# Shell component

`layout.Shell` builds the full dashboard document shell.

## Behavior

- Desktop: fixed sidebar + header + main content.
- Mobile: CSS sheet panel controlled by checkbox + label.
- Optional header slot before theme toggle.

## Accessibility

- `aria-label` and `aria-controls` on menu controls.
- `role="dialog"` on mobile panel.

## Extensibility

- Set custom stylesheet path through `ShellProps.CSSPath`.
- Pass extra header actions via `HeaderExtra`.

Theme behavior should be handled by app script; `Shell` exposes attributes and data to your existing toggle logic.
