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

	if v, err := g.SetView("hello", maxX/6*5+1, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		fmt.Fprintln(v, "Hello world!")
	}

	return nil
}
