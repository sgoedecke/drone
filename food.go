package main

import (
	"math/rand"
	"time"
)

type Food struct {
	X     int
	Y     int
	World *World
}

type Pheromone struct {
	X       int
	Y       int
	targetX int
	targetY int
	Anthill *Anthill
	Age     int
}

type FoodSource struct {
	X     int
	Y     int
	World *World
}

func (fs *FoodSource) Act() {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // have to re-seed each time, apparently :(

	if r.Intn(10) < 8 { // high chance to spawn nothing
		return
	}

	// random chance to spawn a new food tile
	x := r.Intn(7) - 3 + fs.X
	y := r.Intn(7) - 3 + fs.Y
	spotTaken := false
	for _, f := range fs.World.Food {
		if f.X == x && f.Y == y {
			spotTaken = true
		}
	}
	if !spotTaken {
		fs.World.SpawnFood(x, y)
	}
}
