package main

import (
	"log"
	"os"
	"time"

	"github.com/bitly/go-notify"
	"github.com/jroimartin/gocui"
	"github.com/nowa/ncov2019/cui"
	"github.com/nowa/ncov2019/model"
	"github.com/nowa/ncov2019/pkg/scraper"
)

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func main() {
	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(cui.Layout)
	cui.G = g

	go func() {
		for range time.Tick(time.Second * 10) {
			ncovdata, err := scraper.GetAllData()
			if err != nil {
				log.Fatal(err)
			}

			cd := model.ParseCountryData(ncovdata)
			notify.Post("_COUNTRY_DATA_UPDATED_", cd)
		}
	}()

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
		os.Exit(1)
	}
}
