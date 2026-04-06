# Project structure

Typical consuming app layout:

- `go.mod` + app modules
- `cmd/` with HTTP server
- `internal/` domain logic
- `views/` or `templates/` (`.templ`)
- `static/css/input.css` and generated `static/css/app.css`
- optional `static/css/ui8kit/*` synced theme files
- `package.json` if CSS build runs in repo

`templ generate` compiles `.templ` into `_templ.go`.
