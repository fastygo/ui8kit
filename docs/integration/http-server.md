# HTTP server and static assets

UI8Kit’s `layout.Shell` defaults to:

```html
<link rel="stylesheet" href="/static/css/app.css"/>
```

Your server must serve the compiled file at that URL, or you must change `layout.ShellProps.CSSPath`.

## Option A: `os.DirFS` or `http.FileServer`

If `app.css` lives on disk at `./static/css/app.css` relative to the working directory:

```go
fs := http.FileServer(http.Dir("static"))
http.Handle("/static/", http.StripPrefix("/static/", fs))
```

Then `GET /static/css/app.css` returns the file.

## Option B: embed UI8Kit `styles` only

You can serve UI8Kit’s embedded package for quick experiments:

```go
import "github.com/fastygo/ui8kit/styles"

http.Handle("/static/css/ui8kit/", http.StripPrefix("/static/css/ui8kit/",
    http.FileServer(http.FS(styles.FS))))
```

That serves `base.css`, `components.css`, and `latty.css` **separately**. It does **not** replace a full Tailwind build: your app still needs a compiled `app.css` if you rely on utilities generated from your own templates.

**Recommended production setup:** one `app.css` from [Tailwind setup](tailwind-setup.md) and a single `/static/css/app.css` route.

## Custom `CSSPath`

```go
layout.Shell(layout.ShellProps{
    Title:   "My App",
    CSSPath: "/assets/styles.css",
    // ...
})
```

Serve `/assets/styles.css` from your router accordingly.

## Security headers

This library outputs HTML; follow standard Go practices: set `Content-Type`, consider CSP for `script-src` and `style-src` if you extend Shell with third-party scripts.

## Caching

For production, set `Cache-Control` on hashed CSS filenames if you use a build step that fingerprints assets. UI8Kit does not prescribe a caching strategy.
