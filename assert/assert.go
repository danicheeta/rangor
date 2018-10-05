package assert

import (
	"github.com/jroimartin/gocui"
	"fmt"
)

func Nil(g *gocui.Gui, err error) {
	if err != nil {
		v, _ := g.SetView("alarm", 1, 1, 300, 10)
		fmt.Fprint(v, err)
	}
}
