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

	plr, err := newPlayer(renderer)
	if err != nil {
		log.Println("creating player:", err)
		return
	}

	// enemies
	var enemies []basicEnemy

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			x := (float64(i) / 5) * screenWidth + (basicEnemySize / 2)
			y := float64(j * basicEnemySize) + basicEnemySize / 2

			enemy, err := newBasicEnemy(renderer, x, y)
			if err != nil {
				log.Println("creating basic enemy:", err)
				return
			}
			enemies = append(enemies, enemy)
		}
	}


	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		plr.draw(renderer)
		plr.update()

		for _, enemy := range enemies {
			enemy.draw(renderer)
		}

		renderer.Present()
	}

}