package main

import rl "github.com/gen2brain/raylib-go/raylib"

// This file will handle the audio parts of the game cause why not *idk if this will stay or not*

var Music = rl.LoadMusicStream("Audio\\Music\\MenuMusic.mp3")

func LoadMusic(m Mode) {
	switch m {
	case MainMenu:
		Music = rl.LoadMusicStream("Audio\\Music\\MenuMusic.mp3")
		rl.PlayMusicStream(Music)
	}
}

func UpdateMusic() {
	rl.UpdateMusicStream(Music)
}

func EndMusic() {
	rl.UnloadMusicStream(Music)
}
