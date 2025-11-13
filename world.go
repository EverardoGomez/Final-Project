package main

import (
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// This file will hold all the structs that will make up the game world. I couldn't think of a good name for it so world it is :)

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
	left  rl.Rectangle
	right rl.Rectangle
	top   rl.Rectangle
}

// Draws the world
func DrawWorld() {
	tempFloor := rl.Rectangle{
		X:      0,
		Y:      0,
		Width:  1500,
		Height: 1500,
	}

	rl.DrawRectanglePro(tempFloor, rl.Vector2Zero(), 60, rl.SkyBlue)
}

// This function adds effects to items
func (i *Item) AddOnUseEffect(test func()) {
	i.OnUseEffect = append(i.OnUseEffect, test)
}

// Basic attakcs of items
func (i *Item) BasicSwing() {
	//hitBox := rl.NewRectangle(i.Pos.X, i.Pos.Y, 100, 50)

}

// Throws a sword slash that becomes a projectile
func (i *Item) ThrowSword() {

}

func (t *Tile) CreateTile(num int) {
	for i := 0; i < num; i++ {

	}
}

func DrawTiles(t *[]Tile) {
	for i := 0; i < len(*t); i++ {
		rl.DrawRectanglePro((*t)[i].top, rl.Vector2Zero(), 75, rl.Gray)
	}
}

// This function creates enemies
func CreateEnemies(num int) []Entity {
	enemies := []Entity{}
	for i := 0; i < num; i++ {
		enemy := Entity{
			Position:  rl.Vector2Zero(),
			HitBox:    rl.NewRectangle(0, 0, 100, 50),
			Direction: rl.NewVector2(0, 0),
			Health:    10,
		}
		enemies = append(enemies, enemy)
	}
	return enemies
}

// This function draws the slice of enemies
func SpawnEnemies(e []Entity) {
	for i := 0; i < len(e); i++ {
		spawnPoint := rl.NewVector2(float32(rand.IntN(1920)), float32(rand.IntN(1080)))
		e[i].Position = spawnPoint
		e[i].HitBox.X = spawnPoint.X
		e[i].HitBox.Y = spawnPoint.Y
		rl.DrawRectanglePro(e[i].HitBox, e[i].Position, 0, rl.Red)
	}
}
