// Package styles provides embedded CSS for ui8kit themes and icons.
//
// Serve via http.FileServer:
//
//	http.Handle("/static/css/", http.StripPrefix("/static/css/",
//	    http.FileServer(http.FS(styles.FS))))
//
// Or read individual files:
//
//	data, _ := styles.FS.ReadFile("base.css")
package styles

import "embed"

//go:embed base.css shell.css components.css latty.css
var FS embed.FS
