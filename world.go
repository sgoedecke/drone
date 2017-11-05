package main

import (
	"time"
)

func (w *World) Loop() {
	for {
		w.Act()
		w.Draw() // draw the world
		time.Sleep(100 * time.Millisecond)
	}
}

type World struct {
	Anthills []Anthill
	Width    int
	Height   int
}

func (w *World) HandleCollision(ant1 *Ant, ant2 *Ant) {
	if ant1.Anthill.Color == ant2.Anthill.Color {
		return // ants are from the same anthill, no problem.
	} else {
		// ants declare war
		ant1.Die()
		ant2.Die()
	}
}

func (w *World) Act() {
	for i,_ := range w.Anthills {
		w.Anthills[i].Act()
	}
}
