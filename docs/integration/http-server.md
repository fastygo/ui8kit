# HTTP server and static assets

`layout.Shell` expects `/static/css/app.css`, `/static/js/theme.js`, and `/static/js/ui8kit.js` by default.

## Quick route

```go
fs := http.FileServer(http.Dir("static"))
http.Handle("/static/", http.StripPrefix("/static/", fs))
```

## Embedded fallback (optional)

```go
import "github.com/fastygo/ui8kit/styles"

http.Handle("/static/css/ui8kit/", http.StripPrefix("/static/css/ui8kit/",
	http.FileServer(http.FS(styles.FS))))
```

## Production recommendation

- serve one compiled `/static/css/app.css`.
- vendor `theme.js` and `ui8kit.js` with `sync-assets`.
- set `layout.ShellProps.CSSPath` only when route differs.
- set `ShellProps.ThemeJSPath` / `ShellProps.AppJSPath` only when script routes differ.
- set caching headers according to your asset strategy.
