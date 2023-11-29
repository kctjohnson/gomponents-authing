package components

import (
	g "github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"

	c "github.com/maragudk/gomponents/components"
)

type PageLink struct {
	Path string
	Name string
}

func Navbar(currentPath string, startLinks []PageLink, endLinks []PageLink) g.Node {
	return html.Nav(html.Class("bg-gray-700 mb-4"),
		Container(
			html.Div(
				html.Class("flex items-center justify-between space-x-4 h-16"),
				html.Div(
					html.Class("space-x-4"),
					g.Group(g.Map(startLinks, func(l PageLink) g.Node {
						return NavbarLink(l.Path, l.Name, currentPath == l.Path)
					})),
				),
				html.Div(
					html.Class("space-x-4"),
					g.Group(g.Map(endLinks, func(l PageLink) g.Node {
						return NavbarLink(l.Path, l.Name, currentPath == l.Path)
					})),
				),
			),
		),
	)
}

// NavbarLink is a link in the Navbar.
func NavbarLink(path, text string, active bool) g.Node {
	return html.A(html.Href(path), g.Text(text),
		// Apply CSS classes conditionally
		c.Classes{
			"px-3 py-2 rounded-md text-sm font-medium focus:outline-none focus:text-white focus:bg-gray-700": true,
			"text-white bg-gray-900":                           active,
			"text-gray-300 hover:text-white hover:bg-gray-700": !active,
		},
	)
}
