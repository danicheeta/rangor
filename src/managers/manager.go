package managers

import (
	"github.com/jroimartin/gocui"
	"github.com/danicheeta/ranger/src/managers/front"
	"github.com/danicheeta/ranger/src/managers/left"
	"github.com/danicheeta/ranger/src/managers/right"
)

func Manager(g *gocui.Gui) error {
	g.Cursor = true

	front.Manager(g)
	left.Manager(g)
	right.Manager(g)

	g.SetCurrentView("ls")

	return nil
}
