package bindings

import (
	"github.com/jroimartin/gocui"
	"os/exec"
	"fmt"
)

var CurrentPath = "/home/daniel"

func down(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		x, y := v.Cursor()
		v.SetCursor(x, y + 1)

		rightView, _ := g.View("right")
		buf, _ := v.Line(y + 1)
		rightView.Clear()
		fmt.Fprint(rightView, getlsIn(buf))
	}

	return nil
}

func up(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		x, y := v.Cursor()
		v.SetCursor(x, y - 1)

		rightView, _ := g.View("right")
		buf, _ := v.Line(y - 1)
		rightView.Clear()
		fmt.Fprint(rightView, getlsIn(buf))
	}
	return nil
}

func left(g *gocui.Gui, v *gocui.View) error {
	return nil
}

func right(g *gocui.Gui, v *gocui.View) error {
	defer v.SetCursor(0, 0)

	_, y := v.Cursor()
	buf, _ := v.Line(y)
	CurrentPath = CurrentPath + "/" + buf

	rightView, _ := g.View("right")
	rightBuffer := rightView.ViewBuffer()
	rightView.Clear()

	lsView, _ := g.View("ls")
	lsBuffwer := lsView.ViewBuffer()
	lsView.Clear()
	fmt.Fprint(lsView, rightBuffer)


	leftView, _ := g.View("left")
	fmt.Fprint(leftView, lsBuffwer)

	l, _ := lsView.Line(0)
	fmt.Fprint(rightView, getlsIn(l))

	return nil
}

func getls() string {
	cmd := exec.Command("ls", CurrentPath)
	data, err := cmd.Output()
	if err != nil {
		panic(CurrentPath)
	}

	return string(data)
}

func getlsIn(s string) string {
	cmd := exec.Command("ls", CurrentPath + "/" + s)
	data, err := cmd.Output()
	if err != nil {
		panic(CurrentPath)
	}

	return string(data)
}