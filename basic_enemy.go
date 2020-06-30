package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	basicEnemySize = 105
)

func newBasicEnemy(renderer *sdl.Renderer, position vector) *element {
	basicEnemy := &element{}

	basicEnemy.position = position
	basicEnemy.rotation = 180

	// sr := newSpriteRenderer(basicEnemy, renderer, "assets/images/basic_enemy.bmp")
	// basicEnemy.addComponent(sr)
	idleSequence, err := newSequence(renderer, "sprites/basic_enemy/idle", 10, true)
	if err != nil {
		panic(fmt.Errorf("creating idle sequence: %v", err))
	}
	destroySequence, err := newSequence(renderer, "sprites/basic_enemy/destroy", 20, false)
	if err != nil {
		panic(fmt.Errorf("creating destroy sequence: %v", err))
	}

	sequences := map[string]*sequence{
		"idle":    idleSequence,
		"destroy": destroySequence,
	}

	animator := newAnimator(basicEnemy, sequences, "idle")
	basicEnemy.addComponent(animator)

	vtb := newVulnerableToBullets(basicEnemy)
	basicEnemy.addComponent(vtb)

	col := circle{
		center: basicEnemy.position,
		radius: 38,
	}
	basicEnemy.collisions = append(basicEnemy.collisions, col)

	basicEnemy.active = true

	return basicEnemy
}
