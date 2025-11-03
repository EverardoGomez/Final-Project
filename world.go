package main

import rl "github.com/gen2brain/raylib-go/raylib"

// This file will hold all the structs that will make up the game world. I couldn't think of a good name for it so world it is :)

type Entity struct {
	Position rl.Vector2
	HitBox   rl.Rectangle
}

type Circle struct {
	Position rl.Vector2
	Radius   float32
}
