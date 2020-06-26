package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth = 600
	screenHeight = 800
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Println("Initializing SDL:", err)
		return
	}

	window, err := sdl.CreateWindow(
		"Gaming in Go - Episode 2", 
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	
	if err != nil {
		log.Println("Initializing window:", err)
		return
	}
	defer window.Destroy()
	
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Println("Initializing window:", err)
		return
	}
	defer renderer.Destroy()
	log.Println("starting window")
	for {
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		renderer.Present()
	}

}