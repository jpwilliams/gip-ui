package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type Filter struct {
	Groups []string
}

func (f *Filter) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	v, err := g.SetView("filter", maxX/6*5+1, 0, maxX-1, maxY-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	v.Autoscroll = true
	v.Title = " Filters "
	v.Frame = true

	fmt.Fprintln(v, "Hello world!")

	return nil
}
