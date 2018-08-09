package right

import (
	"github.com/jroimartin/gocui"
	"os/exec"
	"github.com/danicheeta/ranger/src/bindings"
)

func Manager(g *gocui.Gui) error {
	x, y := g.Size()
	v, err := g.SetView("right", (x * 4 / 5) - 2, 2, x -2, y - 2)
	if err != gocui.ErrUnknownView {
		return err
	}
	
	v.SelBgColor = gocui.ColorBlack
	v.SelFgColor = gocui.ColorWhite

	lsView, _ := g.View("ls")
	lsFirstLine := lsView.BufferLines()[0]

	_, err = v.Write(getls(lsFirstLine))
	if err != nil {
		return err
	}

	return nil
}

func getls(s string) []byte {
	cmd := exec.Command("ls", bindings.CurrentPath + "/" + s)
	data, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	return data
}
