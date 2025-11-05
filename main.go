package main

import rl "github.com/gen2brain/raylib-go/raylib"

// This is a test for github
func main() {

	rl.InitWindow(1920, 1080, "My Final Project :)")
	defer rl.CloseWindow()
	rl.SetExitKey(0)

	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()

	rl.SetTargetFPS(60)

	cam := rl.NewCamera2D(
		rl.NewVector2(0, 0),
		rl.NewVector2(0, 0),
		0,
		1,
	)

	mode := MainMenu
	LoadMusic(mode)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.BeginMode2D(cam)

		rl.EndMode2D()
		rl.EndDrawing()
	}

	EndMusic()
}
