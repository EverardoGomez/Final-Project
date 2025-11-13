package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// This file will handle all of the physics in my game; aka my bootleg physics engine

var deltaTime = rl.GetFrameTime()

// Moves the Entity by the given vector
func UpdateVector(e *Entity, v rl.Vector2) {
	e.Position.X += v.X * deltaTime
	e.Position.Y += v.Y * deltaTime

	e.HitBox.X += v.X * deltaTime
	e.HitBox.Y += v.Y * deltaTime
}

// Detects the Collision between two rects; returns false by default
func RectCollision(r1 rl.Rectangle, r2 rl.Rectangle) bool {
	if rl.CheckCollisionRecs(r1, r2) {
		return true
	}
	return false
}

// Detects the Collision between two circles; returns false by default
func CircCollision(c1 Circle, c2 Circle) bool {
	if rl.CheckCollisionCircles(c1.Position, c1.Radius, c2.Position, c2.Radius) {
		return true
	}
	return false
}

// converts degrees to radients
func Deg2Rad(d float32) float32 {
	return d * (math.Pi / 180)
}

// converts radients to a vector that points in the direction based off the angle
func Deg2Rotate(rad float32) rl.Vector2 {
	return rl.NewVector2(float32(math.Cos(float64(rad))), float32(math.Sin(float64(rad))))
}
