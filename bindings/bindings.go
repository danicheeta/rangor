package bindings

import (
	"github.com/jroimartin/gocui"
)

var keymaps = map[rune]func(*gocui.Gui, *gocui.View) error{
	'j': down,
	'k': up,
	'h': left,
	'l': right,
}

func AddDefaultBindings(g *gocui.Gui) {
	for i := range keymaps {
		if err := g.SetKeybinding("", i, gocui.ModNone, keymaps[i]); err != nil {
			panic(err)
		}
	}
}



