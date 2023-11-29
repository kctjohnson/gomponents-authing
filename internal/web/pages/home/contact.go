package home

import (
	"authing/internal/web"

	g "github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func ContactPage() (string, web.BodyFunc) {
	return "Contact", func() g.Node {
		return html.Div(
			html.H1(g.Text("Contact us")),
			html.P(g.Text("Just do it.")),
		)
	}
}
