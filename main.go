package main

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// This is a test for github
// LMAO
func main() {

	rl.InitWindow(1920, 1080, "My Final Project :)")
	defer rl.CloseWindow()

	// Runtime Check
	running := true

	// Allow ESC to be used
	rl.SetExitKey(0)

	// Starts the Audio
	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()

	rl.SetTargetFPS(60)

	// The Player starting info
	player := Entity{
		Position:     rl.NewVector2(0, 0),
		HitBox:       rl.NewRectangle(0, 0, 50, 50),
		Direction:    rl.NewVector2(0, 0),
		Health:       10,
		Speed:        500,
		Rotate:       0.0,
		AnimationFSM: NewAnimationFSM(),
	}

	Idle1Text := rl.LoadTexture("Assets\\sword.png")
	Idle1 := NewAnimation("Idle1", Idle1Text, 1, 12)
	player.AnimationFSM.AddAnimation(Idle1)
	player.AnimationFSM.ChangeState("Idle1")

	melee1 := Item{Name: "Sword", Damage: 1, Color: rl.Red}

	player.Inventory = append(player.Inventory, melee1)

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

	// Menu Buttons
	Start := Button{
		Position:   rl.NewVector2(1600, 500),
		Width:      300,
		Height:     125,
		Text:       "Start",
		TextSize:   50,
		ColorTheme: mainMenuTheme,
	}

	Continue := Button{
		Position:   rl.NewVector2(1600, 500),
		Width:      300,
		Height:     125,
		Text:       "Continue",
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

	dead := false
	mode := MainMenu
	LoadMusic(mode)

	for !rl.WindowShouldClose() && running {
		switch mode {
		case MainMenu:
			rl.BeginDrawing()

			rl.ClearBackground(rl.Gray)

			rl.DrawText("MAIN MENU", 800, 400, 50, rl.Black)

			DrawButton(&Start)
			DrawButton(&Quit)

			if DetectMouseClick(&Start) {
				fmt.Println(mode)
				mode = Game

			}

			if DetectMouseClick(&Quit) {
				running = false
			}

			rl.EndDrawing()
		case Game:
			rl.BeginDrawing()
			rl.ClearBackground(rl.Green)

			if player.Health <= 0 {
				dead = true
			}
			if rl.IsKeyPressed(rl.KeyEscape) || dead {
				mode = GameMenu
			}

			if rl.IsKeyDown(rl.KeyW) {
				player.Position.Y -= float32(player.Speed) * rl.GetFrameTime()
				player.Direction.Y = -1
			}
			if rl.IsKeyDown(rl.KeyS) {
				player.Position.Y += float32(player.Speed) * rl.GetFrameTime()
				player.Direction.Y = 1
			}
			if rl.IsKeyDown(rl.KeyD) {
				player.Position.X += float32(player.Speed) * rl.GetFrameTime()
				player.Direction.X = 1
			}
			if rl.IsKeyDown(rl.KeyA) {
				player.Position.X -= float32(player.Speed) * rl.GetFrameTime()
				player.Direction.X = -1
			}

			if rl.IsKeyPressed(rl.KeyLeftShift) {
				player.Position.X += 150 * 2 * player.Direction.X
				player.Position.Y += 150 * 2 * player.Direction.Y
			}

			player.HitBox.X = player.Position.X
			player.HitBox.Y = player.Position.Y

			// updates the camera
			cam.Target = player.Position

			// Gets the player to rotate with the mouse direction
			mouseWorld := rl.GetScreenToWorld2D(rl.GetMousePosition(), cam)
			dx := mouseWorld.X - player.Position.X
			dy := mouseWorld.Y - player.Position.Y
			angle := float32(math.Atan2(float64(dy), float64(dx))) * (180.0 / math.Pi)
			player.Rotate = angle

			rl.BeginMode2D(cam)

			rl.DrawRectangle(100, 100, 100, 100, rl.LightGray)

			if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
				BasicSword(&player)
			}

			// Draws the Player's hitbox
			rl.DrawRectanglePro(
				player.HitBox,
				rl.NewVector2(player.HitBox.Width/2, player.HitBox.Height/2),
				player.Rotate,
				rl.Blue,
			)
			player.AnimationFSM.DrawFSM(player.Position, player.Rotate, player.Inventory[0].Color)

			rl.EndMode2D()

			rl.DrawRectangle(30, 1000, 500, 50, rl.White)
			rl.DrawRectangle(30, 1000, int32(50*player.Health), 50, rl.Red)

			rl.EndDrawing()
		case GameMenu:
			rl.BeginDrawing()

			rl.ClearBackground(rl.Gray)

			if dead {
				rl.DrawText("DEAD", 800, 400, 50, rl.Black)
			} else {
				rl.DrawText("PAUSED", 800, 400, 50, rl.Black)
			}

			DrawButton(&Continue)
			DrawButton(&Quit)

			if DetectMouseClick(&Continue) {
				fmt.Println(mode)
				mode = Game

			}

			if DetectMouseClick(&Quit) {
				running = false
			}

			rl.EndDrawing()
		}
	}

	EndMusic()
}
