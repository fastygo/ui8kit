# HTTP server and static assets

`layout.Shell` expects `/static/css/app.css` by default.

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
- set `layout.ShellProps.CSSPath` only when route differs.
- set caching headers according to your asset strategy.
