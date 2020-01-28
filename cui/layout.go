package cui

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

var G *gocui.Gui

func Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if err := HeaderView(g, -1, -1, maxX, 1); err != nil {
		panic(err)
	}

	if err := FooterView(g, -1, maxY-2, maxX, maxY); err != nil {
		panic(err)
	}

	_ = ShowCountryView(g)

	return nil
}

func ShowCountryView(g *gocui.Gui) error {
	maxX, _ := g.Size()

	if err := CountryTitleView(g, 5, 5, maxX-5, 7); err != nil {
		panic(err)
	}

	return nil
}

func UpdateView(g *gocui.Gui, vn string, s string) error {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View(vn)
		if err != nil {
			log.Panicln(err)
			return err
		}

		v.Clear()
		v.SetCursor(0, 0)
		g.SetViewOnTop(v.Name())

		fmt.Fprintf(v, s)

		return nil
	})

	return nil
}
