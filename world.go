package main

import (
	"math/rand/v2"

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

type Tile struct {
	X      float32
	Y      float32
	TileW  float32
	TileHm float32
	Depth  float32
	Type   Type
	Color  rl.Color
}

func makeTiles() [50][50]Tile {
	MaxHoles := 10
	SpawnTile := 10
	TrapTiles := 20
	t := [50][50]Tile{}

	for i := 0; i < len(t); i++ {
		for j := 0; j < len(t[i]); j++ {

			r := rand.IntN(100)

			if MaxHoles > 0 && r < 10 {
				t[i][j].Type = Hole
				MaxHoles--
			} else if SpawnTile > 0 && r < 25 {
				t[i][j].Type = Spawn
				SpawnTile--
			} else if TrapTiles > 0 && r < 35 {
				t[i][j].Type = Trap
				TrapTiles--
			} else {
				t[i][j].Type = Basic
			}

			t[i][j].TileW = 100
			t[i][j].TileHm = 75
			t[i][j].Depth = 50
		}
	}

	return t
}

func DrawTiles(t *[50][50]Tile) {
	originX := float32(600)
	originY := float32(52.6)

	for i := 0; i < 50; i++ {
		originX += 50
		originY += 37.35
		for j := 0; j < 50; j++ {
			t[i][j].X = originX - float32(50*j)
			t[i][j].Y = originY + 37.35*float32(j)
			DrawTile(&t[i][j])
		}
	}
}

// Draws a tile
func DrawTile(t *Tile) {
	color := rl.Color{}

	top := rl.NewVector2(t.X, t.Y)
	left := rl.NewVector2(t.X-(t.TileW/2), t.Y+(t.TileHm/2))
	right := rl.NewVector2(t.X+(t.TileW/2), t.Y+(t.TileHm/2))
	bottom := rl.NewVector2(t.X, t.Y+(t.TileHm))

	if t.Type == Basic {
		color = rl.Black
	}
	if t.Type == Trap {
		color = rl.Lime
	}
	if t.Type == Hole {
		color = rl.Blue
	}
	if t.Type == Spawn {
		color = rl.Purple
	}
	rl.DrawTriangleFan([]rl.Vector2{top, left, bottom, right}, color)

	leftBot := rl.NewVector2(left.X, left.Y+t.Depth)
	bottomBot := rl.NewVector2(bottom.X, bottom.Y+t.Depth)

	rl.DrawTriangleFan([]rl.Vector2{bottomBot, bottom, left, leftBot}, rl.LightGray)

	rightBot := rl.NewVector2(right.X, right.Y+t.Depth)

	rl.DrawTriangleFan([]rl.Vector2{right, bottom, bottomBot, rightBot}, rl.Gray)

}
