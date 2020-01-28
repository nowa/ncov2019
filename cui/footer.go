package cui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func FooterView(g *gocui.Gui, x, y, maxX, maxY int) error {

	if v, err := g.SetView("footer", x, y, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.FgColor = gocui.Attribute(15 + 1)
		v.BgColor = gocui.Attribute(0)

		v.Autoscroll = false
		v.Editable = false
		v.Wrap = false
		v.Frame = false
		v.Overwrite = true

		fmt.Fprintf(v, "...")
	}

	return nil
}
