package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type Body struct {
	Groups []string
}

func (b *Body) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("body", maxX/6+1, 0, (maxX/6)*5, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		fmt.Fprintln(v, "BODY")
	}

	return nil
}
