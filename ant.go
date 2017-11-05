package main

import (
	"math/rand"
	"time"
)

type Ant struct {
	X       int
	Y       int
	Anthill *Anthill
	World   *World
	Dead    bool
}

func (a *Ant) Act() {
	if a.Dead {
		return
	}

	// move to a random block one tile away
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // have to re-seed each time, apparently :(
	x := r.Intn(3) - 1                                   // -1, 0, or 1
	y := r.Intn(3) - 1                                   // -1, 0, or 1
	a.Move(a.X+x, a.Y+y)
}

func (a *Ant) Die() {
	// set flag
	a.Dead = true
}

func (a *Ant) Move(x int, y int) {
	// check for out of bounds
	if x > a.World.Width || y > a.World.Height || x < 0 || y < 0 {
		return
	}

	// check for collision with other ants
	for _, anthill := range a.World.Anthills {
		for _, ant := range anthill.Ants {
			if ant.X == x && ant.Y == y && !ant.Dead {
				a.World.HandleCollision(&ant, a)
				return
			}
		}
	}

	// if no collision, move the ant
	a.X = x
	a.Y = y
}
