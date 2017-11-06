package main

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"time"
)

type Anthill struct {
	Ants       []Ant
	Pheromones []Pheromone
	X          int
	Y          int
	Color      termbox.Attribute
	World      *World
}

func (ah *Anthill) SpawnAnt() {

	for _, ant := range ah.Ants {
		if ant.X == ah.X && ant.Y == ah.Y {
			return
		}
	}

	// spawn an ant directly over the anthill
	a := Ant{}
	a.X = ah.X
	a.Y = ah.Y
	a.Anthill = ah
	a.World = ah.World
	a.Dead = false
	ah.Ants = append(ah.Ants, a)
}

func (ah *Anthill) Act() {
	// random chance to spawn a new ant
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // have to re-seed each time, apparently :(
	rint := r.Intn(10)
	if rint > 8 {
		ah.SpawnAnt()
	}

	// make all its ants act
	ants := ah.Ants // cache ants here, since we're mutating ah.Ants as we go - see below explanation
	for i, _ := range ants {
		ants[i].Act() // can't use _,ant and ant.Move() - it doesn't move the ant
		// the problem here is that `range` copies values. so you can't mutate _ without doing another lookup
		// need to figure out a better solution for iterating over an array & removing values at the same time
	}

	// we can now purge dead ants, since we've taken all the actions we can and our index woes are over
	var newants []Ant
	for i, _ := range ants {
		if !ants[i].Dead {
			newants = append(newants, ants[i])
		}
	}
	ah.Ants = newants

	// age all pheromones and purge old pheromones
	var np []Pheromone
	for i := range ah.Pheromones {
		p := ah.Pheromones[i]
		p.Age = p.Age + 1
		if p.Age < 30 { // pheromone age
			np = append(np, p)
		}
	}
	ah.Pheromones = np
}
