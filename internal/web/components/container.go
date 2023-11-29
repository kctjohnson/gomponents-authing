package components

import (
	g "github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func Container(children ...g.Node) g.Node {
	return html.Div(html.Class("px-2 sm:px-6 lg:px-8"), g.Group(children))
}
