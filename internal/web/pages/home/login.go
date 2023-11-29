package home

import (
	"authing/internal/web"
	"authing/internal/web/components"

	g "github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func LoginPage() (string, web.BodyFunc) {
	return "Login", func() g.Node {
		return html.Div(
			html.H1(g.Text("User Login")),
			html.Div(
				html.Class("flex flex-row"),
				html.FormEl(
					html.Class("flex flex-col basis-1/2"),
					html.Action("/auth/login"),
					html.Method("post"),
					components.Input("user", "user", "Username...", true),
					components.Input("password", "password", "Password...", true),
					components.Button("submit", "Login", "bg-green"),
				),
			),
		)
	}
}
