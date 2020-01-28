package main

import (
	"log"
	"os"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/nowa/ncov2019/cui"
	// "github.com/nowa/ncov2019/model"
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
	g.Cursor = true
	cui.G = g

	// Need to find a better way to do the first request (channel?)
	time.AfterFunc(20*time.Millisecond, func() {
		scraper.GetAndParseData()
	})

	go func() {
		for range time.Tick(time.Minute * 5) {
			scraper.GetAndParseData()
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
