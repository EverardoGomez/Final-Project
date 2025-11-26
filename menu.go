package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// This will handle the UI functions my game

// The Button Struct; Everything I need Button Wise is contained here
type Button struct {
	Position     rl.Vector2
	Width        int
	Height       int
	Text         string
	TextSize     int
	MouseClikced bool
	MouseOver    bool
	MousePressed bool
	ColorTheme
	OnClickButton []func()
}

// Used for anything color related for menus
type ColorTheme struct {
	BaseColor   rl.Color
	AccentColor rl.Color
	TextColor   rl.Color
}

// These enums handle the game state
type Mode int

const (
	MainMenu Mode = 0
	Game     Mode = 1
	GameMenu Mode = 2
)

// This function adds on click events to the given button
func (b *Button) AddOnCLickButton(test func()) {
	b.OnClickButton = append(b.OnClickButton, test)
}

// Detects if the button is over a button
func IsMouseOver(b *Button) {
	mousePos := rl.GetMousePosition()
	b.MouseOver = false

	if mousePos.X < b.Position.X || mousePos.X > float32(b.Width)+b.Position.X {
		return
	}
	if mousePos.Y < b.Position.Y || mousePos.Y > float32(b.Height)+b.Position.Y {
		return
	}
	b.MouseOver = true
}

// This Function draws the button
func DrawButton(b *Button) {
	color := b.BaseColor
	IsMouseOver(b)

	if b.MouseOver {
		color = b.AccentColor
	}

	rl.DrawRectangle(int32(b.Position.X), int32(b.Position.Y), int32(b.Width), int32(b.Height), color)
	rl.DrawText(b.Text, int32(b.Position.X), int32(b.Position.Y), int32(b.TextSize), b.TextColor)
}

// Detects Mouse Clicks on buttons
func DetectMouseClick(b *Button) bool {
	IsMouseOver(b)

	if b.MouseOver {
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			return true
		}
	}
	return false
}
