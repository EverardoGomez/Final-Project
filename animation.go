package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Animation struct {
	Name         string
	SpriteSheet  rl.Texture2D
	CurrentIndex int
	MaxIndex     int
	Timer        float32
	SwitchTime   float32
	Loop         bool
}

func NewAnimation(name string, text rl.Texture2D, numSprites int, switchTime float32) Animation {
	animation := Animation{
		Name:        name,
		SpriteSheet: text,
		MaxIndex:    numSprites,
		SwitchTime:  switchTime,
		Loop:        true,
	}
	return animation
}

func (a Animation) DrawAnimation(pos rl.Vector2, rotate float32, color rl.Color) {
	rect1 := rl.NewRectangle(float32(32*a.CurrentIndex), 0, 32, 32)
	rect2 := rl.NewRectangle(pos.X, pos.Y, 50, 100)
	rl.DrawTexturePro(a.SpriteSheet, rect1, rect2, rl.NewVector2(0, 0), rotate, color)
}

func (a *Animation) UpdateAnimation() {
	a.Timer += rl.GetFrameTime()

	if a.Timer > a.SwitchTime {
		a.Timer = 0
		a.CurrentIndex++
	}

	if a.CurrentIndex > a.MaxIndex {
		if a.Loop {
			a.CurrentIndex = 0
		} else {
			a.CurrentIndex = a.MaxIndex
		}
	}
}
