package home

import (
	"authing/internal/db/models"
	"authing/internal/repositories"
	"authing/internal/web"
	"authing/internal/web/components"
	"database/sql"

	g "github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func UsersPage(accRepo *repositories.Accounts) (string, web.BodyFunc) {
	return "Users", func() g.Node {
		accs, err := accRepo.GetAll()
		if err != nil && err != sql.ErrNoRows {
			return html.Div(
				html.H1(g.Text("Users")),
				html.P(g.Text("Failed to load users")),
			)
		}
		return html.Div(
			html.H1(g.Text("Users")),
			g.Group(
				g.Map(accs, func(a models.Account) g.Node {
					return components.UserView(a)
				}),
			),
		)
	}
}
