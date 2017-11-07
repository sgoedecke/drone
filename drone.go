package main

import (
	"github.com/nsf/termbox-go"
)

func main() {
	// initialize termbox
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetOutputMode(termbox.Output256) // set 256-color mode

	// initialize an event queue and poll eternally, sending events to a channel
	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	// set up world, anthills, ants
	w := World{}
	w.Width = 80
	w.Height = 40
	w.X = 1
	w.Y = 1

	a1 := Anthill{}
	a1.X = 5
	a1.Y = 5
	a1.Color = termbox.ColorRed
	a1.World = &w
	a1.Food = 5
	w.Anthills = append(w.Anthills, a1)

	a2 := Anthill{}
	a2.X = 76
	a2.Y = 30
	a2.Color = termbox.ColorGreen
	a2.World = &w
	a2.Food = 5
	w.Anthills = append(w.Anthills, a2)

	fs := FoodSource{}
	fs.X = 30
	fs.Y = 10
	fs.World = &w
	w.FoodSources = append(w.FoodSources, fs)

	fs2 := FoodSource{}
	fs2.X = 50
	fs2.Y = 30
	fs2.World = &w
	w.FoodSources = append(w.FoodSources, fs2)

	// kick off the act -> draw loop
	go w.Loop()

	// listen for q/ctrl+C etc.
	for {
		event := <-eventQueue
		if event.Type == termbox.EventKey {
			switch {
			case event.Ch == 'q' || event.Key == termbox.KeyEsc || event.Key == termbox.KeyCtrlC || event.Key == termbox.KeyCtrlD:
				return
			}
		}
	}
}
