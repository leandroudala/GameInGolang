package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	basicEnemySize = 105
)

type basicEnemy struct {
	tex *sdl.Texture
	x, y float64
}

func newBasicEnemy(renderer *sdl.Renderer, x, y float64) (be basicEnemy, err error){
	img, err := sdl.LoadBMP("assets/images/basic_enemy.bmp")
	if err != nil {
		return basicEnemy{}, fmt.Errorf("loading basic enemy sprite: %v", err)
	}
	defer img.Free()
	be.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return basicEnemy{}, fmt.Errorf("creating basic enemy texture: %v", err)
	}

	be.x = x
	be.y = y

	return be, nil
}

func (be *basicEnemy) draw(renderer *sdl.Renderer) {
	x := be.x - basicEnemySize / 2.0
	y := be.y - basicEnemySize / 2.0

	renderer.CopyEx(
		be.tex,
		&sdl.Rect{X: 0, Y:0, W: basicEnemySize, H: basicEnemySize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: basicEnemySize, H: basicEnemySize},
		180, // angle
		&sdl.Point{X: basicEnemySize / 2, Y: basicEnemySize /2},
		sdl.FLIP_NONE)
}