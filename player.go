package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type player struct {
	tex *sdl.Texture
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

	return
}

func (p *player) draw(renderer *sdl.Renderer) {
	renderer.Copy(
		p.tex, 
		&sdl.Rect{X: 0, Y: 0, W: 105, H: 105}, 
		&sdl.Rect{X: 0, Y: 0, W: 105, H: 105})
}