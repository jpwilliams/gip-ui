package main

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/jroimartin/gocui"
)

type Body struct {
	View         *gocui.View
	haveRunSetup bool
	Groups       []string
}

func (b *Body) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	var err error
	b.View, err = g.SetView("body", maxX/6+1, 0, (maxX/6)*5, maxY-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			panic(err)
		}
	}

	b.View.Autoscroll = true
	b.View.Title = " Logs - service-attachments "
	b.View.Frame = true
	b.View.Wrap = true

	if !b.haveRunSetup {
		b.haveRunSetup = true
		b.setup(g)
	}

	return nil
}

func (b *Body) setup(g *gocui.Gui) {
	b.updateLoop(g)
}

func (b *Body) updateLoop(g *gocui.Gui) {
	go func() {
		for {
			cmd := exec.Command("gip", "view", "work", "-a", "1.week")
			out, _ := cmd.Output()

			g.Execute(func(g *gocui.Gui) error {
				v, err := g.View("body")
				if err != nil {
					panic(err)
				}

				v.Clear()
				fmt.Fprintln(v, string(out))
				return nil
			})
			time.Sleep(time.Second * 5)
		}
	}()
}
