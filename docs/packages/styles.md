# Package `styles`

Embedded theme and shared UI CSS.

## Import

```go
import "github.com/fastygo/ui8kit/styles"
```

## Files in `styles.FS`

- `base.css` — tokens and base styles.
- `shell.css` — shell/layout classes.
- `components.css` — reusable components.
- `latty.css` — icon masks.

## Usage

```go
data, err := styles.FS.ReadFile("base.css")
```

Or serve package files directly with `http.FileServer(http.FS(styles.FS))` for quick checks. In production, prefer one compiled `app.css` via Tailwind.
