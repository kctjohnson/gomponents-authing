package components

import (
	g "github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func Input(inputType string, name string, placeholder string, required bool) g.Node {
	return html.Input(
		html.Class("border rounded basis-full mb-4"),
		html.Type(inputType),
		html.Name(name),
		g.If(placeholder != "", html.Placeholder(placeholder)),
		g.If(required, html.Required()),
	)
}
