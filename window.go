package main

import "github.com/jroimartin/gocui"

type Window struct {
	*gocui.View
	lines       []string
	path        string
	cursorIndex int
}

type Coordinate struct {
	x0, y0, x1, y1 int
}

func NewWindow(g *gocui.Gui, name string, c Coordinate) (*Window, error) {
	w := &Window{}

	v, err := g.SetView(name, c.x0, c.y0, c.x1, c.y1)
	if err != gocui.ErrUnknownView {
		return nil, err
	}

	v.Wrap = true
	v.Highlight = true
	v.SelBgColor = gocui.ColorBlack
	v.SelFgColor = gocui.ColorWhite

	w.View = v

	return w, nil
}
