package main

import (
	"github.com/nsf/termbox-go"
	"time"
)

type World struct {
	Anthills    []Anthill
	Food        []Food
	FoodSources []FoodSource
	Width       int
	Height      int
	X           int
	Y           int
}

func (w *World) Loop() {
	for {
		w.Act()
		w.Draw() // draw the world
		time.Sleep(100 * time.Millisecond)
	}
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
	for i, _ := range w.Anthills {
		w.Anthills[i].Act()
	}
	for i, _ := range w.FoodSources {
		w.FoodSources[i].Act()
	}
}

func (w *World) SpawnAnthill(x int, y int, color termbox.Attribute) {
	a := Anthill{}
	a.X = x
	a.Y = y
	a.Color = color
	a.World = w
	a.Food = 15
	w.Anthills = append(w.Anthills, a)
}

func (w *World) SpawnFoodSource(x int, y int) {
	fs := FoodSource{}
	fs.X = x
	fs.Y = y
	fs.World = w
	w.FoodSources = append(w.FoodSources, fs)
}

func (w *World) SpawnFood(x int, y int) {
	f := Food{}
	f.X = x
	f.Y = y
	f.World = w
	w.Food = append(w.Food, f)
}
