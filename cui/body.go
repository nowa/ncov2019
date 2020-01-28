package cui

import (
	"fmt"

	"github.com/bitly/go-notify"
	"github.com/jroimartin/gocui"
)

func BodyView(g *gocui.Gui, x, y, maxX, maxY int) error {
	if v, err := g.SetView("body", x, y, maxX, maxY); err != nil {
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

		fmt.Fprintf(v, " ")

		c := make(chan interface{})
		notify.Start("_COUNTRY_DATA_UPDATED_", c)
		go func() {
			for {
				data := <-c
				cd, ok := data.(map[string]map[string]int)

				if ok && len(cd) > 1 {
					g.Update(func(g *gocui.Gui) error {
						m := ""
						for k, ct := range cd {
							m += StringRandom(fmt.Sprintf("%s has comfirmed %d, suspected %d, cured %d, dead %d.\n", k, ct["C"], ct["S"], ct["H"], ct["D"]))
						}
						return UpdateView(g, "body", m)
					})
				}
			}
		}()
	}

	return nil
}
