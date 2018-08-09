package left

import (
	"github.com/jroimartin/gocui"
	"os/exec"
)

func Manager(g *gocui.Gui) error {
	x, y := g.Size()
	v, err := g.SetView("left", 2, 2, (x / 5) + 2, y - 2)
	if err != gocui.ErrUnknownView {
		return err
	}

	v.SelBgColor = gocui.ColorBlack
	v.SelFgColor = gocui.ColorWhite

	_, err = v.Write(getParentLs())
	if err != nil {
		return err
	}

	return nil
}

func getParentLs() []byte {
	cmd := exec.Command("ls", "..", "-a")
	data, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	return data
}
