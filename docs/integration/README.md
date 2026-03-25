# Integration

These guides describe how to connect UI8Kit to a real Go application: compiling CSS with Tailwind v4 and serving assets so `layout.Shell` and browsers stay in sync.

## Guides

- [Tailwind CSS v4 setup](tailwind-setup.md) — `package.json`, `input.css`, watch scripts, syncing UI8Kit styles
- [HTTP server and static assets](http-server.md) — routes, `embed.FS`, matching default CSS paths

## Summary

1. Add UI8Kit with `go get github.com/fastygo/ui8kit`.
2. Add Tailwind CLI to your app (dev dependency) and create `static/css/input.css`.
3. Copy or reference UI8Kit’s `styles/*.css` into the Tailwind pipeline (see Tailwind doc).
4. Run `tailwindcss` to produce `static/css/app.css`.
5. Serve `/static/css/app.css` (or set `layout.ShellProps.CSSPath` to your URL).
