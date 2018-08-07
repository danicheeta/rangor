package main

import (
	"github.com/jroimartin/gocui"
	"github.com/danicheeta/ranger/src/managers"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		panic(err)
	}
	defer g.Close()

	g.SetManagerFunc(managers.Manager)

	addExitHandler(g)

	if err := g.MainLoop(); err != gocui.ErrQuit {
		println(err)
	}
}

func addExitHandler(g *gocui.Gui) {
	exit := func (*gocui.Gui, *gocui.View) error {return gocui.ErrQuit}
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, exit); err != nil {
		panic(err)
	}
}




