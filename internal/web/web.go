package web

import (
	"authing/internal/web/components"
	"time"

	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

/*
This exists so that pages that need to reload DB data can do that on reload

- TODO: Find a better way to handle that kinda stuff, this isn't the best
*/
type BodyFunc func() g.Node

func Page(authed bool, title, path string, body g.Node) g.Node {
	var startLinks []components.PageLink
	var endLinks []components.PageLink
	if authed {
		startLinks = []components.PageLink{
			{Path: "/user/adminpanel", Name: "Admin Panel"},
			{Path: "/user/users", Name: "Users"},
		}
		endLinks = []components.PageLink{
			{Path: "/auth/logout", Name: "Logout"},
		}
	} else {
		startLinks = []components.PageLink{
			{Path: "/", Name: "Home"},
			{Path: "/contact", Name: "Contact"},
			{Path: "/about", Name: "About"},
			{Path: "/users", Name: "Users"},
		}
		endLinks = []components.PageLink{
			{Path: "/login", Name: "Login"},
			{Path: "/register", Name: "Register"},
		}
	}

	// HTML5 boilerplate document
	return c.HTML5(c.HTML5Props{
		Title:    title,
		Language: "en",
		Head: []g.Node{
			html.Script(html.Src("https://cdn.tailwindcss.com?plugins=typography")),
		},
		Body: []g.Node{
			components.Navbar(path, startLinks, endLinks),
			components.Container(
				Prose(body),
				PageFooter(),
			),
		},
	})
}

func Prose(children ...g.Node) g.Node {
	return html.Div(html.Class("prose"), g.Group(children))
}

func PageFooter() g.Node {
	return html.Footer(html.Class("prose prose-sm prose-indigo"),
		html.P(
			// We can use string interpolation directly, like fmt.Sprintf.
			g.Textf("Rendered %v. ", time.Now().Format(time.RFC3339)),

			// Conditional inclusion
			g.If(time.Now().Second()%2 == 0, g.Text("It's an even second.")),
			g.If(time.Now().Second()%2 == 1, g.Text("It's an odd second.")),
		),

		html.P(html.A(html.Href("https://www.gomponents.com"), g.Text("gomponents"))),
	)
}
