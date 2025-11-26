package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// This file will hold all the structs that will make up the game world. I couldn't think of a good name for it so world it is :)

type Type int

const (
	Basic = iota
	Trap
	Hole
	Spawn
)

type Entity struct {
	Position  rl.Vector2
	HitBox    rl.Rectangle
	Direction rl.Vector2
	Health    int
	Rotate    float32
	Item
}

type Circle struct {
	Position  rl.Vector2
	Radius    float32
	Direction rl.Vector2
}

type Item struct {
	Name        string
	Damage      int
	Pos         rl.Vector2
	OnUseEffect []func()
}
