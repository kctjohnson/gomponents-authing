package home

import (
	"authing/internal/web"
	"authing/internal/web/components"

	g "github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func RegisterPage() (string, web.BodyFunc) {
	return "Register", func() g.Node {
		return html.Div(
			html.H1(g.Text("Register")),
			html.Div(
				html.Class("flex flex-row"),
				html.FormEl(
					html.Class("flex flex-col basis-1/2"),
					html.Action("/auth/register"),
					html.Method("post"),
					components.Input("user", "user", "Username...", true),
					components.Input("password", "password", "Password...", true),
					components.Button("submit", "Register", "bg-green"),
				),
			),
		)
	}
}
