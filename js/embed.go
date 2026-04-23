// Package js contains embedded browser assets that still ship with ui8kit.
//
// Only theme.js remains embedded here. Interactive ARIA behavior is sourced
// from @ui8kit/aria and vendored into applications via scripts/cmd/sync-assets.
package js

import "embed"

//go:embed theme.js
var FS embed.FS
