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

	cam := rl.NewCamera2D(
		rl.NewVector2(0, 0),
		rl.NewVector2(0, 0),
		0,
		1,
	)

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
			rl.BeginMode2D(cam)

			rl.ClearBackground(rl.Blue)

			rl.EndMode2D()
			rl.EndDrawing()
		case GameMenu:
		}
	}

	EndMusic()
}
