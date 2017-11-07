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

	// draw FoodSources
	for _, fs := range w.FoodSources {
		fs.Draw()
	}

	// draw anthills & ants
	for ai, hill := range w.Anthills {
		hill.Draw()

		// // draw anthill stats
		// termbox.SetCell(w.X, w.Y+w.Height+ai+2, 64, hill.Color, termbox.ColorDefault) // anthill
		// for i := 0; i < hill.Food; i++ {
		// 	termbox.SetCell(w.X+i+1, w.Y+w.Height+ai+2, 176, hill.Color, termbox.ColorDefault) // big dot
		// }
	}

	// draw walls
	for y := 0; y < w.Height+1; y++ {
		termbox.SetCell(w.X-1, w.Y+y, 9608, termbox.ColorDefault, termbox.ColorDefault)         //#
		termbox.SetCell(w.X+w.Width+1, w.Y+y, 9608, termbox.ColorDefault, termbox.ColorDefault) //#
	}

	for x := 0; x < w.Width+3; x++ {
		termbox.SetCell(w.X+x-1, w.Y-1, 9608, termbox.ColorDefault, termbox.ColorDefault)          //#
		termbox.SetCell(w.X+x-1, w.Y+w.Height+1, 9608, termbox.ColorDefault, termbox.ColorDefault) //#
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
	termbox.SetCell(ah.X+ah.World.X, ah.Y+ah.World.Y, 11044, ah.Color, termbox.ColorDefault) //@
}

func (a *Ant) Draw() {
	if a.Dead {
		return
	}

	termbox.SetCell(a.X+a.World.X, a.Y+a.World.Y, 35, a.Anthill.Color, termbox.ColorDefault) //#
}

func (f *Food) Draw() {
	termbox.SetCell(f.X+f.World.X, f.Y+f.World.Y, 176, termbox.ColorWhite, termbox.ColorDefault) // big dot
}

func (p *Pheromone) Draw() {
	termbox.SetCell(p.X+p.Anthill.World.X, p.Y+p.Anthill.World.Y, 184, p.Anthill.Color, termbox.ColorDefault) //.
}

func (fs *FoodSource) Draw() {
	termbox.SetCell(fs.X+fs.World.X, fs.Y+fs.World.Y, 177, termbox.ColorWhite, termbox.ColorDefault) //tree symbol
}
