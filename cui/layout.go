package cui

import (
	"github.com/jroimartin/gocui"
)

func Layout(g *gocui.Gui) error {
	maxX, _ := g.Size()

	if err := HeaderView(g, -1, -1, maxX, 1); err != nil {
		panic(err)
	}

	return nil
}
