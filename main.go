package main

import rl "github.com/gen2brain/raylib-go/raylib"

// This is a test for github
// LMAO
func main() {

	rl.InitWindow(1920, 1080, "My Final Project :)")
	defer rl.CloseWindow()

	running := true

	// Allow ESC to be used
	rl.SetExitKey(0)

	// Starts the Audio
	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()

	rl.SetTargetFPS(60)

	// The Player starting info
	player := Entity{
		Position:  rl.NewVector2(0, 0),
		HitBox:    rl.NewRectangle(0, 0, 100, 50),
		Direction: rl.NewVector2(0, 0),
		Health:    10,
		Rotate:    0.0,
	}

	// The cam for the player
	cam := rl.Camera2D{
		Offset: rl.NewVector2(float32(rl.GetScreenWidth())/2,
			float32(rl.GetScreenHeight())/2),
		Target:   player.Position,
		Rotation: 0.0,
		Zoom:     1,
	}

	// Color Themes for their respective menus
	mainMenuTheme := ColorTheme{
		BaseColor:   rl.Purple,
		AccentColor: rl.DarkPurple,
		TextColor:   rl.White,
	}

	// Main Menu Buttons
	Start := Button{
		Position:   rl.NewVector2(1600, 500),
		Width:      300,
		Height:     125,
		Text:       "Start",
		TextSize:   50,
		ColorTheme: mainMenuTheme,
	}

	Quit := Button{
		Position:   rl.NewVector2(1600, 650),
		Width:      300,
		Height:     125,
		Text:       "Quit",
		TextSize:   50,
		ColorTheme: mainMenuTheme,
	}

	mode := MainMenu
	LoadMusic(mode)

	for !rl.WindowShouldClose() && running {
		switch mode {
		case MainMenu:
			rl.BeginDrawing()

			rl.ClearBackground(rl.Gray)

			rl.DrawText("MAIN MENU", 800, 400, 50, rl.Black)

			DrawButton(Start)
			DrawButton(Quit)

			if DetectMouseClick(&Start) {
				mode = Game
			}

			if DetectMouseClick(&Quit) {
				running = false
			}

			rl.EndDrawing()
		case Game:
			rl.BeginDrawing()
			rl.ClearBackground(rl.Green)

			rl.BeginMode2D(cam)

			// updates the camera
			cam.Target = player.Position

			// Draws the Player's hitbox
			rl.DrawRectangleRec(player.HitBox, rl.Blue)

			rl.EndMode2D()

			rl.DrawRectangle(30, 1000, 500, 50, rl.White)
			rl.DrawRectangle(30, 1000, int32(50*player.Health), 50, rl.Red)

			rl.EndDrawing()
		case GameMenu:
		}
	}

	EndMusic()
}
