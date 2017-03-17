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

	if v, err := g.SetView("menu", 1, 0, maxX/6, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		fmt.Fprintln(v, "MENU")
	}

	return nil
}
