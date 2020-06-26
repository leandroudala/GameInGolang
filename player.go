package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed = 0.1
	playerSize = 105
)

type player struct {
	tex *sdl.Texture
	x, y float64
}

func newPlayer(renderer *sdl.Renderer) (p player, err error) {
	img, err := sdl.LoadBMP("assets/images/player.bmp")
	if err != nil {
		return player{}, fmt.Errorf("loading player sprite: %v", err)
	}
	defer img.Free()

	p.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return player{}, fmt.Errorf("loading player sprite: %v", err)
	}

	p.x = screenWidth / 2
	p.y = screenHeight - playerSize / 2.0

	return
}

func (p *player) draw(renderer *sdl.Renderer) {
	// Converting player coordinates to top left of sprite
	x := p.x - playerSize / 2.0
	y := p.y - playerSize / 2.0

	renderer.Copy(
		p.tex, 
		&sdl.Rect{X: 0, Y: 0, W: 105, H: 105}, 
		&sdl.Rect{X: int32(x), Y: int32(y), W: 105, H: 105})
}

func (p *player) update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		// move player to left
		p.x -= playerSpeed
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		// move player to left
		p.x += playerSpeed
	}

	if keys[sdl.SCANCODE_UP] == 1{
		p.y -= playerSpeed
	} else if keys[sdl.SCANCODE_DOWN] == 1 {
		p.y += playerSpeed
	}
}