package main

import (
	"log"

	"github.com/jroimartin/gocui"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	// kanan, bawah
	if v, err := g.SetView("hello", maxX/2-25, maxY/2-4, maxX/2+25, maxY/2+3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "login"
	}

	if v, err := g.SetView("username", maxX/2-17, maxY/2-3, maxX/2+17, maxY/2-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "username"
	}

	if v, err := g.SetView("password", maxX/2-17, maxY/2, maxX/2+17, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "password"
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
