package components

import (
	"authing/internal/db/models"
	"fmt"

	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	"github.com/maragudk/gomponents/html"
)

func UserView(a models.Account) g.Node {
	return html.P(g.Text(a.Username))
}

func AuthedUserView(a models.Account) g.Node {
	return html.Div(
		html.ID(fmt.Sprintf("userview%d", a.ID)),
		html.Class("flex items-center"),
		html.P(
			g.Text(fmt.Sprintf("%d | %s | %s", a.ID, a.Username, a.Password)),
		),
		Button(
			"button",
			"Delete",
			"bg-red",
			hx.Target(fmt.Sprintf("#userview%d", a.ID)),
			hx.Post(fmt.Sprintf("/api/delete?id=%d", a.ID)),
			hx.Swap("delete"),
		),
	)
}
