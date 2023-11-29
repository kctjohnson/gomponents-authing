package home

import (
	"authing/internal/web"

	g "github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func IndexPage() (string, web.BodyFunc) {
	return "Welcome!", func() g.Node {
		return html.Div(
			html.H1(g.Text("Welcome to this example page")),
			html.P(g.Text("I hope it will make you happy. 😄 It's using TailwindCSS for styling.")),
		)
	}
}
