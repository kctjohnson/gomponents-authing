package pages

import (
	"authing/internal/web"

	g "github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func InvalidPage() (string, web.BodyFunc) {
	return "Invalid URL", func() g.Node {
		return html.Div(
			html.H1(g.Text("Invalid URL")),
			html.P(
				g.Text(
					"The page you requested either doesn't exist, or you don't have the proper credentials to view it.",
				),
			),
		)
	}
}
