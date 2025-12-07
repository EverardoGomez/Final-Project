package main

import (
	"math"
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// This file will hold all the structs that will make up the game world. I couldn't think of a good name for it so world it is :)
type Type int

const (
	Fodder = iota
	Ranged
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
	AnimationFSM
	Circle
	Type
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
	Color       rl.Color
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

func fodder(e *Entity) Entity {
	// 1. Random angle between 0 and 2π
	angle := rand.Float32() * (2 * math.Pi)

	// 2. Convert angle to direction vector
	dir := rl.NewVector2(
		float32(math.Cos(float64(angle))),
		float32(math.Sin(float64(angle))),
	)

	// 3. Multiply by the spawn radius
	spawnOffset := rl.Vector2Scale(dir, e.Radius)

	// 4. Final spawn position
	spawnPos := rl.NewVector2(
		e.Position.X+spawnOffset.X,
		e.Position.Y+spawnOffset.Y,
	)

	// Create enemy
	enemy := Entity{
		Speed:    100,
		Health:   10,
		Position: spawnPos,
		Type:     Fodder,
	}

	return enemy
}

func ranged(e *Entity) Entity {
	// 1. Random angle between 0 and 2π
	angle := rand.Float32() * (2 * math.Pi)

	// 2. Convert angle to direction vector
	dir := rl.NewVector2(
		float32(math.Cos(float64(angle))),
		float32(math.Sin(float64(angle))),
	)

	// 3. Multiply by the spawn radius
	spawnOffset := rl.Vector2Scale(dir, e.Radius)

	// 4. Final spawn position
	spawnPos := rl.NewVector2(
		e.Position.X+spawnOffset.X,
		e.Position.Y+spawnOffset.Y,
	)

	// Create enemy
	enemy := Entity{
		Speed:    100,
		Health:   10,
		Position: spawnPos,
		Type:     Ranged,
	}

	return enemy
}

func createEnemies(num int, e *Entity) []Entity {
	enemyList := []Entity{}

	for i := 0; i < num; i++ {

		var enemy Entity

		// 50/50 chance
		if rand.IntN(2) == 0 {
			enemy = fodder(e)
		} else {
			enemy = ranged(e)
		}

		// Make sure this enemy doesn't spawn too close to another
		valid := true
		for j := 0; j < len(enemyList); j++ {
			dx := enemy.Position.X - enemyList[j].Position.X
			dy := enemy.Position.Y - enemyList[j].Position.Y
			dist2 := dx*dx + dy*dy

			// Minimum distance between spawns (adjust as needed)
			minDist := float32(30.0)

			if dist2 < minDist*minDist {
				valid = false
				break
			}
		}

		if valid {
			enemyList = append(enemyList, enemy)
		} else {
			i-- // try again for this enemy index
		}
	}

	return enemyList
}

//////////////////////////////////////////////////////////////////////////////////////////// WEAPON CODE /////////////////////////////////////////////////////////////////////////////////////////////

func BasicSword(e *Entity) {
	entity := Entity{
		HitBox: rl.NewRectangle(e.HitBox.X, e.HitBox.Y, 80, 50),
		Rotate: e.Rotate,
	}

	rl.DrawRectanglePro(
		entity.HitBox,
		rl.NewVector2((e.HitBox.Width/2)+15, (e.HitBox.Height/2)+30),
		entity.Rotate,
		rl.Red,
	)
}
