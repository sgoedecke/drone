package main

import (
	"github.com/nsf/termbox-go"
)

func (w *World) Draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	// draw Food
	for _, food := range w.Food {
		food.Draw()
	}

	// draw anthills & ants
	for _, hill := range w.Anthills {
		hill.Draw()
	}

	// draw walls
	for y := 0; y < w.Height; y++ {
		termbox.SetCell(w.X-1, w.Y+y, 35, termbox.ColorDefault, termbox.ColorDefault)       //#
		termbox.SetCell(w.X+w.Width, w.Y+y, 35, termbox.ColorDefault, termbox.ColorDefault) //#
	}

	for x := 0; x < w.Width; x++ {
		termbox.SetCell(w.X+x, w.Y-1, 35, termbox.ColorDefault, termbox.ColorDefault)        //#
		termbox.SetCell(w.X+x, w.Y+w.Height, 35, termbox.ColorDefault, termbox.ColorDefault) //#
	}

	_ = termbox.Flush()
}

func (ah *Anthill) Draw() {
	// draw pheromones
	for _, p := range ah.Pheromones {
		p.Draw()
	}

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

func (f *Food) Draw() {
	termbox.SetCell(f.X+f.World.X, f.Y+f.World.Y, 184, termbox.ColorWhite, termbox.ColorDefault) //.
}

func (p *Pheromone) Draw() {
	termbox.SetCell(p.X+p.Anthill.World.X, p.Y+p.Anthill.World.Y, 184, p.Anthill.Color, termbox.ColorDefault) //.
}
