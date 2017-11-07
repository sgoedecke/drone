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

	w.SpawnAnthill(10, 10, termbox.ColorGreen)
	w.SpawnAnthill(70, 30, termbox.ColorRed)

	w.SpawnFoodSource(30, 30)
	w.SpawnFoodSource(20, 20)
	w.SpawnFoodSource(35, 35)
	w.SpawnFoodSource(50, 10)

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
