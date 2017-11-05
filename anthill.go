package main

import (
	"github.com/nsf/termbox-go"
	"math/rand"
  "time"
)

type Anthill struct {
	Ants  []Ant
	X     int
	Y     int
	Color termbox.Attribute
	World *World
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
	ah.Ants = append(ah.Ants, a)
}

func (ah *Anthill) Act() {
	// random chance to spawn a new ant
  r := rand.New(rand.NewSource(time.Now().UnixNano())) // have to re-seed each time, apparently :(
	rint := r.Intn(10)
	if rint > 1 {
		ah.SpawnAnt()
	}

	// make all its ants act
  ants := ah.Ants // cache ants here, since we're mutating ah.Ants as we go
	for i, _ := range ants {
		ants[i].Act() // can't use _,ant and ant.Move() - it doesn't move the ant???
	}
}

type Ant struct {
	X       int
	Y       int
	Anthill *Anthill
	World   *World
}

func (a *Ant) Act() {
	// move to a random block one tile away
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // have to re-seed each time, apparently :(
	x := r.Intn(3) - 1 // -1, 0, or 1
	y := r.Intn(3) - 1 // -1, 0, or 1
	a.Move(a.X+x, a.Y+y)
}

func (a *Ant) Die() {
	// remove from parent anthill
	for i, ant := range a.Anthill.Ants {
		if ant.X == a.X && ant.Y == a.Y {
      a.Anthill.Ants = append(a.Anthill.Ants[:i], a.Anthill.Ants[i+1:]...)
		}
	}
}

func (a *Ant) Move(x int, y int) {
	// check for out of bounds
	if x > a.World.Width || y > a.World.Height || x < 0 || y < 0 {
		return
	}

	// check for collision with other ants
	for _, anthill := range a.World.Anthills {
		for _, ant := range anthill.Ants {
			if ant.X == x && ant.Y == y {
				a.World.HandleCollision(&ant, a)
				return
			}
		}
	}

	// if no collision, move the ant
	a.X = x
	a.Y = y
}
