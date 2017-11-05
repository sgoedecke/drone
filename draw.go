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

	// draw walls
	for y := 0; y < w.Height; y++ {
		termbox.SetCell(w.X, w.Y+y, 35, termbox.ColorDefault, termbox.ColorDefault)         //#
		termbox.SetCell(w.X+w.Width, w.Y+y, 35, termbox.ColorDefault, termbox.ColorDefault) //#
	}

	for x := 0; x < w.Width; x++ {
		termbox.SetCell(w.X+x, w.Y, 35, termbox.ColorDefault, termbox.ColorDefault)          //#
		termbox.SetCell(w.X+x, w.Y+w.Height, 35, termbox.ColorDefault, termbox.ColorDefault) //#
	}
	_ = termbox.Flush()
}

func (ah *Anthill) Draw() {
	// draw ants
	for _, ant := range ah.Ants {
		ant.Draw()
	}

	// draw self
	termbox.SetCell(ah.X+ah.World.X, ah.Y+ah.World.Y, 64, ah.Color, termbox.ColorDefault) //@
}

func (a *Ant) Draw() {
	if a.Dead {
		return
	}

	termbox.SetCell(a.X+a.World.X, a.Y+a.World.Y, 35, a.Anthill.Color, termbox.ColorDefault) //#
}
