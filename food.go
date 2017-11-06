package main

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
