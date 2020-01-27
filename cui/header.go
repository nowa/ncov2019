package cui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func HeaderView(g *gocui.Gui, x, y, maxX, maxY int) error {

	if v, err := g.SetView("header", x, y, maxX, maxY); err != nil {
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

		welcome := "⣿ A realtime tracker for nCov2019"

		pad := maxX - len(welcome) + 1

		var i int
		for i <= pad {
			i++
			welcome += " "
		}
		fmt.Fprintf(v,
			StringFormatBoth(
				Color.White,
				Color.Header,
				welcome,
				[]string{"1"},
			),
		)
	}

	return nil
}
