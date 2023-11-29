package user

import (
	"authing/internal/web"

	g "github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func AdminPanelPage() (string, web.BodyFunc) {
	return "Authorized Page", func() g.Node {
		return html.Div(
			html.H1(g.Text("Admin Panel")),
			html.P(g.Text("Welcome to the Admin Panel.")),
		)
	}
}
