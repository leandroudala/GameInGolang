package main

import (
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 800
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Println("Initializing SDL:", err)
		return
	}
	// creating window
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

	// creating renderer
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Println("Initializing window:", err)
		return
	}
	defer renderer.Destroy()

	// creating player
	plr := newPlayer(renderer)

	// enemies
	var enemies []*element

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			x := (float64(i)/5)*screenWidth + (basicEnemySize / 2)
			y := float64(j*basicEnemySize) + basicEnemySize/2

			enemy := newBasicEnemy(renderer, vector{x: x, y: y})

			enemies = append(enemies, enemy)
		}
	}

	initBulletPool(renderer)
	// starting screen
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		err = plr.onDraw(renderer)
		if err != nil {
			fmt.Println("drawing player:", err)
			return
		}
		err = plr.onUpdate()
		if err != nil {
			fmt.Println("updating player:", err)
			return
		}

		for _, enemy := range enemies {
			enemy.onDraw(renderer)
		}

		for _, bul := range bulletPool {
			bul.draw(renderer)
			bul.update()
		}

		renderer.Present()
	}

}
