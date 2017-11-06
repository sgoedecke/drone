package main

import (
	"math"
	"math/rand"
	"time"
)

type Ant struct {
	X       int
	Y       int
	Anthill *Anthill
	World   *World
	Dead    bool // every time an an anthill acts, it finishes by removing dead ants from its list
	HasFood bool // if has food, head for anthill
	Focused bool // if focused, head for target
	targetX int
	targetY int
}

func (a *Ant) Act() {
	if a.Dead {
		return
	}

	// if the ant is roaming, look for food/pheromones. otherwise drop one
	if !a.Focused && !a.HasFood {

		// look for pheromones
		for _, p := range a.Anthill.Pheromones {
			if a.X == p.X && a.Y == p.Y {
				a.Focused = true
				a.targetX = p.targetX
				a.targetY = p.targetY
			}
		}

		// check for collision with food
		var newfood []Food
		for i := range a.World.Food {
			food := a.World.Food[i] // get direct reference to food
			if a.X == food.X && a.Y == food.Y {
				a.Eat(&food)
			} else {
				newfood = append(newfood, food)
			}
		}
		a.World.Food = newfood

	} else if a.Focused {
		// check for collision with target
		if a.X == a.targetX && a.Y == a.targetY {
			a.Focused = false
		}
	} else if a.HasFood {
		if math.Abs(float64(a.X-a.Anthill.X)) < 2.0 && math.Abs(float64(a.Y-a.Anthill.Y)) < 2.0 { // ant is adj to anthill
			a.HasFood = false
			a.Focused = false
		} else {
			a.DropPheromone()
		}
	}

	if a.HasFood {
		// move back to anthill
		a.MoveTo(a.Anthill.X, a.Anthill.Y)
	} else if a.Focused {
		a.MoveTo(a.targetX, a.targetY)
	} else {
		a.Roam()
	}
}

func (a *Ant) Roam() {
	// move to a random block one tile away
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // have to re-seed each time, apparently :(
	x := r.Intn(3) - 1                                   // -1, 0, or 1
	y := r.Intn(3) - 1                                   // -1, 0, or 1
	a.Move(a.X+x, a.Y+y)
}

func (a *Ant) MoveTo(x int, y int) {
	// move towards target
	var tX int
	var tY int
	if x > a.X {
		tX = a.X + 1
	} else if x < a.X {
		tX = a.X - 1
	} else {
		tX = a.X
	}

	if y > a.Y {
		tY = a.Y + 1
	} else if y < a.Y {
		tY = a.Y - 1
	} else {
		tY = a.Y
	}

	a.Move(tX, tY)
}

func (a *Ant) Die() {
	// set flag so ant gets cleaned up later
	a.Dead = true
}

func (a *Ant) Eat(f *Food) {
	a.Focused = false
	a.HasFood = true
	a.targetX = f.X
	a.targetY = f.Y
}

func (a *Ant) DropPheromone() {
	p := Pheromone{}
	p.X = a.X
	p.Y = a.Y
	p.targetX = a.targetX
	p.targetY = a.targetY
	p.Anthill = a.Anthill

	a.Anthill.Pheromones = append(a.Anthill.Pheromones, p)
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
