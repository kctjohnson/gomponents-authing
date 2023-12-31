package components

import (
	"fmt"

	g "github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func Button(btnType string, text string, color string, children ...g.Node) g.Node {
	btnClass := fmt.Sprintf(
		"%s-600 text-black px-2 rounded hover:%s-700 hover:text-white",
		color,
		color,
	)
	return html.Button(
		html.Class(btnClass),
		html.Type(btnType),
		g.Text(text),
		g.Group(children),
	)
}
