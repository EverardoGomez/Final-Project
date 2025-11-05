package main

import rl "github.com/gen2brain/raylib-go/raylib"

// This file will hold all the structs that will make up the game world. I couldn't think of a good name for it so world it is :)

type Entity struct {
	Position  rl.Vector2
	HitBox    rl.Rectangle
	Direction rl.Vector2
	Item
}

type Circle struct {
	Position  rl.Vector2
	Radius    float32
	Direction rl.Vector2
}

type Item struct {
	Name   string
	Damage int
}

// Draws the world
func DrawWorld() {
	tempFloor := rl.Rectangle{
		X:      0,
		Y:      0,
		Width:  500,
		Height: 500,
	}

	rl.DrawRectanglePro(tempFloor, rl.Vector2Zero(), 60, rl.SkyBlue)
}
