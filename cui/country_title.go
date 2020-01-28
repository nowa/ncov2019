package cui

import (
	"fmt"

	"github.com/bitly/go-notify"
	"github.com/jroimartin/gocui"
)

func CountryTitleView(g *gocui.Gui, x, y, maxX, maxY int) error {
	if v, err := g.SetView("countrytitle", x, y, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			panic(err)
		}

		v.FgColor = gocui.Attribute(15 + 1)
		v.BgColor = gocui.Attribute(0)

		v.Autoscroll = false
		v.Editable = false
		v.Wrap = false
		v.Frame = false
		v.Overwrite = true

		fmt.Fprintf(v, "Loading...")

		c := make(chan interface{})
		notify.Start("_COUNTRY_DATA_UPDATED_", c)
		go func() {
			for {
				data := <-c
				cd, ok := data.([]map[string]int)

				if ok && len(cd) > 1 {
					g.Update(func(g *gocui.Gui) error {
						var title = fmt.Sprintf("Infections in %d countries.", len(cd[0]))
						title = StringFormat(Color.QueryHeader, title, []string{"7"})
						return UpdateView(g, "countrytitle", title)
					})
				}
			}
		}()
	}

	return nil
}
