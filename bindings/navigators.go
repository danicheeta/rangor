package bindings

import (
	"github.com/jroimartin/gocui"
	"os/exec"
	"fmt"
	"strings"
)

var (
	CurrentPath  = "/home/daniel"
	lastDirIndex int
)

func down(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		x, y := v.Cursor()
		v.SetCursor(x, y+1)

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
		v.SetCursor(x, y-1)

		rightView, _ := g.View("right")
		buf, _ := v.Line(y - 1)
		rightView.Clear()
		fmt.Fprint(rightView, getlsIn(buf))
	}
	return nil
}

func left(g *gocui.Gui, v *gocui.View) error {
	defer v.SetCursor(0, lastDirIndex)

	s := strings.Split(CurrentPath, `/`)
	CurrentPath = strings.Join(s[:len(s)-1], `/`)

	leftView, _ := g.View("left")
	leftBuffer := leftView.ViewBuffer()
	leftView.Clear()
	fmt.Fprint(leftView, getlsOut())

	lsView, _ := g.View("ls")
	lsBuffer := lsView.ViewBuffer()
	lsView.Clear()
	fmt.Fprint(lsView, leftBuffer)

	rightView, _ := g.View("right")
	rightView.Clear()
	fmt.Fprint(rightView, lsBuffer)

	return nil
}

func right(g *gocui.Gui, v *gocui.View) error {
	defer v.SetCursor(0, 0)

	_, y := v.Cursor()
	lastDirIndex = y
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
	leftView.Clear()
	fmt.Fprint(leftView, lsBuffwer)

	l, _ := lsView.Line(0)
	fmt.Fprint(rightView, getlsIn(l))

	return nil
}

func getlsIn(s string) string {
	cmd := exec.Command("ls", CurrentPath+"/"+s)
	data, err := cmd.Output()
	if err != nil {
		panic("lsIn: " + s)
	}

	return string(data)
}

func getlsOut() string {
	nodes := strings.Split(CurrentPath, `/`)
	beforePath := strings.Join(nodes[:len(nodes)-1], `/`)

	cmd := exec.Command("ls", beforePath)
	data, err := cmd.Output()
	if err != nil {
		panic("lsout: " + beforePath)
	}

	return string(data)
}
