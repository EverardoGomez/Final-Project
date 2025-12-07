package main

import rl "github.com/gen2brain/raylib-go/raylib"

type AnimationFSM struct {
	Current    Animation
	Animations map[string]Animation
}

func NewAnimationFSM() AnimationFSM {
	return AnimationFSM{Animations: make(map[string]Animation)}
}

func (a *AnimationFSM) AddAnimation(anim Animation) {
	a.Animations[anim.Name] = anim
}

func (a *AnimationFSM) ChangeState(name string) {
	if name == a.Current.Name {
		return
	}

	a.Current = a.Animations[name]
	a.Current.CurrentIndex = 0
}

func (a *AnimationFSM) DrawFSM(pos rl.Vector2, rotate float32, color rl.Color) {
	a.Current.UpdateAnimation()
	a.Current.DrawAnimation(pos, rotate, color)
}
