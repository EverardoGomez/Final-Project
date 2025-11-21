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
	playerSpeed := 200

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

	level := makeTiles()

	leftBorder := rl.Rectangle{X: level[0][0].X + 650, Y: level[0][0].Y + 100, Width: 600 + (50 * 49), Height: 52.6}
	RightBorder := rl.Rectangle{X: level[49][49].X + 650 + (50 * 49), Y: level[49][49].Y + 2050, Width: 600 + (50 * 49), Height: 52.6}

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

			DrawTiles((*[50][50]Tile)(&level))

			// Draws the Player's hitbox
			rl.DrawRectangleRec(player.HitBox, rl.Blue)

			rl.DrawRectanglePro(leftBorder, rl.Vector2Zero(), 143, rl.Pink)
			rl.DrawRectanglePro(RightBorder, rl.Vector2Zero(), 143, rl.Pink)

			RectCollision(player.HitBox, leftBorder)

			// Player Movement
			if rl.IsKeyDown(rl.KeyD) {
				dx := player.Position.X + (float32(playerSpeed) * rl.GetFrameTime())
				player.Position.X = dx
				BoundaryCheck(&player.Position, &level)
				player.HitBox.X = dx
				player.Direction.X = 1
				cam.Target = player.Position
			}
			if rl.IsKeyReleased(rl.KeyD) {
				player.Direction.X = 0
			}
			if rl.IsKeyDown(rl.KeyA) {
				dx := player.Position.X - (float32(playerSpeed) * rl.GetFrameTime())
				player.Position.X = dx
				BoundaryCheck(&player.Position, &level)
				player.HitBox.X = dx
				player.Direction.X = -1
				cam.Target = player.Position
			}
			if rl.IsKeyReleased(rl.KeyA) {
				player.Direction.X = 0
			}
			if rl.IsKeyDown(rl.KeyW) {
				dy := player.Position.Y - (float32(playerSpeed) * rl.GetFrameTime())
				player.Position.Y = dy
				BoundaryCheck(&player.Position, &level)
				player.HitBox.Y = dy
				player.Direction.Y = -1
				cam.Target = player.Position
			}
			if rl.IsKeyReleased(rl.KeyW) {
				player.Direction.Y = 0
			}
			if rl.IsKeyDown(rl.KeyS) {
				dy := player.Position.Y + (float32(playerSpeed) * rl.GetFrameTime())
				player.Position.Y = dy
				BoundaryCheck(&player.Position, &level)
				player.HitBox.Y = dy
				player.Direction.Y = 1
				cam.Target = player.Position
			}
			if rl.IsKeyReleased(rl.KeyS) {
				player.Direction.Y = 0
			}
			// The DASH CODE
			if rl.IsKeyPressed(rl.KeyLeftShift) {
				player.Position.X += 16000 * rl.GetFrameTime() * player.Direction.X
				player.Position.Y += 16000 * rl.GetFrameTime() * player.Direction.Y
				player.HitBox.X += 16000 * rl.GetFrameTime() * player.Direction.X
				player.HitBox.Y += 16000 * rl.GetFrameTime() * player.Direction.Y
				BoundaryCheck(&player.Position, &level)
				cam.Target = player.Position
			}

			rl.EndMode2D()

			rl.DrawRectangle(30, 1000, 500, 50, rl.White)
			rl.DrawRectangle(30, 1000, int32(50*player.Health), 50, rl.Red)

			rl.EndDrawing()
		case GameMenu:
		}
	}

	EndMusic()
}
