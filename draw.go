package main

import (
	"github.com/nsf/termbox-go"
)

func (w *World) Draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	// draw anthills
	for _, hill := range w.Anthills {
		hill.Draw()
	}
	_ = termbox.Flush()
}

func (ah *Anthill) Draw() {
	// draw ants
	for _, ant := range ah.Ants {
		ant.Draw()
	}

	// draw self
	termbox.SetCell(ah.X, ah.Y, 64, ah.Color, termbox.ColorDefault) //@
}

func (a *Ant) Draw() {
	termbox.SetCell(a.X, a.Y, 35, a.Anthill.Color, termbox.ColorDefault) //#
}
