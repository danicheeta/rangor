package managers

import (
	"github.com/jroimartin/gocui"
	"fmt"
	"os/exec"
)

func Manager(g *gocui.Gui) error {
	x, y := g.Size()
	if v, err := g.SetView("hello", 2, 2, x - 2, y - 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		data := getls()
		fmt.Fprintln(v, string(data))
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