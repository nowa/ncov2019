package cui

import (
	"fmt"
	"time"

	"github.com/bitly/go-notify"
	"github.com/jroimartin/gocui"
	"github.com/nowa/ncov2019/util"
	"github.com/xeonx/timeago"
)

var (
	LastUpdatedAt time.Time
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

		c := make(chan interface{})
		notify.Start("_GETTING_DATA_", c)
		go func() {
			for {
				var m string
				data := <-c

				status := []string{"doing", "done"}
				if util.Contains(status, data.(string)) {
					g.Update(func(g *gocui.Gui) error {
						if data.(string) == "doing" {
							m = "Getting data from tencent..."
						} else if data.(string) == "done" {
							LastUpdatedAt = time.Now()
							m = fmt.Sprintf("Done. Last updated at %s", timeago.English.Format(LastUpdatedAt))
						}
						return UpdateView(g, "footer", m)
					})
				}
			}
		}()

		go func() {
			for range time.Tick(time.Second * 10) {
				var m string
				if !LastUpdatedAt.IsZero() && time.Now().Sub(LastUpdatedAt).Seconds() >= 10 {
					m = fmt.Sprintf("Last updated at %s", timeago.English.Format(LastUpdatedAt))
					g.Update(func(g *gocui.Gui) error {
						return UpdateView(g, "footer", m)
					})
				}
			}
		}()
	}

	return nil
}
