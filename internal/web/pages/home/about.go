package home

import (
	"authing/internal/web"

	g "github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func AboutPage() (string, web.BodyFunc) {
	return "About", func() g.Node {
		return html.Div(
			html.H1(g.Text("About this site")),
			html.P(g.Text("This is a site showing off gomponents.")),
		)
	}
}
