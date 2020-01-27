package main

import (
	"log"
	"os"

	"github.com/jroimartin/gocui"
	"github.com/nowa/ncov2019/cui"
	"github.com/nowa/ncov2019/model"
	"github.com/nowa/ncov2019/pkg/scraper"
)

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func main() {
	log.Println("Getting all data...")
	ncovdata, err := scraper.GetAllData()
	if err != nil {
		log.Fatal(err)
	}

	_ = model.ParseData(ncovdata)

	log.Println("Total cities: ", len(ncovdata))

	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(cui.Layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
		os.Exit(1)
	}
}
