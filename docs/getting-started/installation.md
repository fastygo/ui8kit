# Installation

```bash
go get github.com/fastygo/ui8kit@latest
```

Run from app root:

```bash
go install github.com/a-h/templ/cmd/templ@latest
templ generate ./...
go build ./...
go test ./...
```

Use subpackages directly:

```go
import (
    "github.com/fastygo/ui8kit/ui"
    "github.com/fastygo/ui8kit/layout"
    "github.com/fastygo/ui8kit/utils"
)
```

For CSS workflow, follow [Tailwind setup](../integration/tailwind-setup.md).
