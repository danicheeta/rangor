package main

import (
	"github.com/jroimartin/gocui"
	"github.com/danicheeta/ranger/bindings"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		panic(err)
	}
	defer g.Close()

	g.SetManagerFunc(Manager)

	addExitHandler(g)
	bindings.AddDefaultBindings(g)

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




