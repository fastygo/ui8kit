package layout

import (
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/fastygo/ui8kit/utils"
)

func rawScript(js string) templ.Component {
	return templ.ComponentFunc(func(_ context.Context, w io.Writer) error {
		_, err := io.WriteString(w, "<script>"+js+"</script>")
		return err
	})
}

func sidebarLinkClass(active, path string) string {
	if active == path {
		return "bg-accent text-accent-foreground"
	}
	return "text-muted-foreground hover:bg-accent"
}

func sidebarItemClasses(active, path string) string {
	return utils.Cn("flex items-center gap-2 rounded px-4 py-2 text-sm", sidebarLinkClass(active, path))
}

func shellBrand(name string) string {
	if name == "" {
		return "App"
	}
	return name
}

func shellCSS(path string) string {
	if path == "" {
		return "/static/css/app.css"
	}
	return path
}

const themeScript = `(function(){var r=document.documentElement,s=localStorage.getItem("ui8kit-theme"),d=window.matchMedia&&window.matchMedia("(prefers-color-scheme:dark)").matches;r.classList.toggle("dark",s==="dark"||(!s&&d));var u=function(){var i=document.getElementById("theme-toggle-icon");if(i)i.className=r.classList.contains("dark")?"latty latty-sun h-4 w-4":"latty latty-moon h-4 w-4"};u();window.ui8kitToggleTheme=function(){var n=!r.classList.contains("dark");r.classList.toggle("dark",n);localStorage.setItem("ui8kit-theme",n?"dark":"light");u()};document.addEventListener("DOMContentLoaded",u)})();`

const sidebarScript = `window.ui8kitOpenSidebar=function(){var s=document.getElementById("ui8kit-mobile-sidebar"),b=document.getElementById("ui8kit-sidebar-backdrop");if(s&&b){s.classList.remove("hidden");b.classList.remove("hidden");document.body.style.overflow="hidden"}};window.ui8kitCloseSidebar=function(){var s=document.getElementById("ui8kit-mobile-sidebar"),b=document.getElementById("ui8kit-sidebar-backdrop");if(s&&b){s.classList.add("hidden");b.classList.add("hidden");document.body.style.overflow=""}};window.addEventListener("keydown",function(e){if(e&&e.key==="Escape")ui8kitCloseSidebar()});`
