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

/////////////////////////////////////////////////////////////////////////////////////////////// STRUCTS ////////////////////////////////////////////////////////////////////////////////////////

type Entity struct {
	Position  rl.Vector2
	HitBox    rl.Rectangle
	Direction rl.Vector2
	Health    int
	Speed     int
	Rotate    float32
	Inventory []Item
}

type Circle struct {
	Position  rl.Vector2
	Radius    float32
	Direction rl.Vector2
}

type Item struct {
	Name        string
	Damage      int
	Pos         *rl.Vector2
	OnUseEffect []func()
	SpriteRenderer
}

type SpriteRenderer struct {
	Sprite   rl.Texture2D
	Color    rl.Color
	Position *rl.Vector2
	Rect     rl.Rectangle
	Angle    float32
	Scale    float32
}

/////////////////////////////////////////////////////////////////////////////////////////////// SPRITE CODE ////////////////////////////////////////////////////////////////////////////////////////

func NewSpriteRenderer(newSprite rl.Texture2D, newColor rl.Color, newPosition *rl.Vector2) SpriteRenderer {
	sr := SpriteRenderer{
		Sprite:   newSprite,
		Color:    newColor,
		Position: newPosition,
		Angle:    0,
		Scale:    2,
	}
	return sr
}

func (sr *SpriteRenderer) Draw() {
	pos := *sr.Position

	pw := sr.Scale * float32(sr.Sprite.Width)
	ph := sr.Scale * float32(sr.Sprite.Height)

	sr.Rect = rl.NewRectangle(pos.X, pos.Y, pw, ph)

	sourceRect := rl.NewRectangle(0, 0, float32(sr.Sprite.Width), float32(sr.Sprite.Height))
	destRect := rl.NewRectangle(pos.X, pos.Y, float32(sr.Sprite.Width)*sr.Scale, float32(sr.Sprite.Height)*sr.Scale)
	origin := rl.Vector2Scale(rl.NewVector2(float32(sr.Sprite.Width)/2, float32(sr.Sprite.Height)/2), sr.Scale)

	rl.DrawTexturePro(sr.Sprite, sourceRect, destRect, origin, sr.Angle, sr.Color)
}

//////////////////////////////////////////////////////////////////////////////////////////// ENEMY CODE //////////////////////////////////////////////////////////////////////////////////////////////
