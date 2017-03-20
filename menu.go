package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type Menu struct {
	Groups []string
}

func (m *Menu) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	v, err := g.SetView("menu", 1, 0, maxX/6, maxY-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	v.Title = " Groups & Repos "
	v.Frame = true

	fmt.Fprintln(v, "Woop")

	return nil
}
