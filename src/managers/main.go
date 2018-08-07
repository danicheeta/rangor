package managers

import (
	"github.com/jroimartin/gocui"
	"os/exec"
)

func Manager(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView("colors", 2, 2, maxX -2, maxY - 2)
	if err != gocui.ErrUnknownView {
		return err
	}

	v.Highlight = true
	v.SelBgColor = gocui.ColorBlack
	v.SelFgColor = gocui.ColorWhite

	_, err = v.Write(getls())
	if err != nil {
		return err
	}

	return nil
}

func getls() []byte {
	cmd := exec.Command("ls", "-a")
	data, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	return data
}
