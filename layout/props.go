package layout

import "github.com/a-h/templ"

// NavItem describes a single sidebar navigation link.
type NavItem struct {
	Path  string
	Label string
	Icon  string
}

// SidebarProps configures the sidebar navigation.
type SidebarProps struct {
	Items  []NavItem
	Active string
	Mobile bool
}

// HeaderProps configures the top header bar.
type HeaderProps struct {
	Title       string
	Extra       templ.Component
	Trailing    templ.Component
	ThemeToggle ThemeToggleProps
}

// ThemeToggleProps configures copy for the theme toggle button.
type ThemeToggleProps struct {
	Label              string
	SwitchToDarkLabel  string
	SwitchToLightLabel string
}

// ShellProps configures the full page shell (sidebar + header + main).
type ShellProps struct {
	Title          string           // HTML <title>
	Lang           string           // HTML lang attribute (defaults to "en")
	BrandName      string           // sidebar brand label (defaults to "App")
	Active         string           // current path for nav highlight
	NavItems       []NavItem        // sidebar links
	CSSPath        string           // path to CSS file (defaults to "/static/css/app.css")
	HeadExtra      templ.Component  // optional extra <head> content
	HeaderExtra    templ.Component  // optional extra header actions before theme toggle
	HeaderTrailing templ.Component  // optional trailing header actions after theme toggle
	ThemeToggle    ThemeToggleProps // optional copy for the theme toggle button
}
