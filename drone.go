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

	for i := 0; i < 50; i++ {
		f := Food{}
		f.X = 80 - i
		f.Y = 35 + i/10
		f.World = &w
		w.Food = append(w.Food, f)
	}

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
