package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletSize  = 32
	bulletSpeed = 1
)

func newBullet(renderer *sdl.Renderer) *element {
	bullet := &element{}

	sr := newSpriteRenderer(bullet, renderer, "assets/images/player_bullet.bmp")
	bullet.addComponent(sr)

	mover := newBulletMover(bullet, bulletSpeed)
	bullet.addComponent(mover)

	bullet.active = false

	return bullet
}

var bulletPool []*element

func initBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		b := newBullet(renderer)
		bulletPool = append(bulletPool, b)
		elements = append(elements, b)
	}
}

func bulletFromPool() (*element, bool) {
	for _, b := range bulletPool {
		if !b.active {
			return b, true
		}
	}
	return nil, false
}

func (bullet *element) move() {
	pos := &bullet.position
	pos.x += bulletSpeed * math.Cos(bullet.rotation)
	pos.y += bulletSpeed * math.Sin(bullet.rotation)

	if pos.x > screenWidth || pos.x < 0 || pos.y > screenHeight || pos.y < 0 {
		bullet.active = false
	}
}
