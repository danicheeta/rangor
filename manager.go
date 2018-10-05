package main

import (
	"github.com/jroimartin/gocui"
	"os/exec"
	"strings"
	"github.com/danicheeta/ranger/assert"
)

var (
	CurrentPath  = "/home/daniel"
	lastDirIndex int
	windows      []*Window
)

type windowIndex int

const (
	LSWindow     windowIndex = iota
	BeforeWindow
	AfterWindow
)

func Manager(g *gocui.Gui) error {
	x, y := g.Size()
	g.Cursor = true

	_, err := g.SetView("main-frame", 2, 2, x-2, y-2)
	if err != gocui.ErrUnknownView {
		return err
	}

	frontCoordination := Coordinate{(x / 5) + 2, 2, (x * 4 / 5) - 2, y - 2}
	lsWindow, err := NewWindow(g, "front", frontCoordination)
	if err != nil {
		panic(err)
	}

	beforeCoordination := Coordinate{2, 2, (x / 5) + 2, y - 2}
	beforeWindow, err := NewWindow(g, "left", beforeCoordination)
	if err != nil {
		panic(err)
	}

	afterCoordination := Coordinate{(x * 4 / 5) - 2, 2, x - 2, y - 2}
	afterWindow, err := NewWindow(g, "right", afterCoordination)
	if err != nil {
		panic(err)
	}

	_, err = lsWindow.Write(getls())
	assert.Nil(g, err)
	l, _ := lsWindow.Line(0)

	_, err = beforeWindow.Write(getlsOut())
	assert.Nil(g, err)
	_, err = afterWindow.Write(getlsIn(l))
	assert.Nil(g, err)

	windows = []*Window{
		lsWindow,
		beforeWindow,
		afterWindow,
	}

	return nil
}

func getWindow(s windowIndex) *Window {
	return windows[s]
}

func getlsIn(s string) []byte {
	cmd := exec.Command("ls", CurrentPath+"/"+s)
	data, err := cmd.Output()
	if err != nil {
		panic("lsIn: " + s)
	}

	return data
}

func getlsOut() []byte {
	nodes := strings.Split(CurrentPath, `/`)
	beforePath := strings.Join(nodes[:len(nodes)-1], `/`)

	cmd := exec.Command("ls", beforePath)
	data, err := cmd.Output()
	if err != nil {
		panic("lsout: " + beforePath)
	}

	return data
}

func getls() []byte {
	cmd := exec.Command("ls", CurrentPath)
	data, err := cmd.Output()
	if err != nil {
		panic("getLs: " + err.Error())
	}

	return data
}
