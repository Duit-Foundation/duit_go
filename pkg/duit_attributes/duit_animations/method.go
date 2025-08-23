package duit_animations

type AnimationMethod = uint8

const (
	Forward AnimationMethod = iota
	Repeat
	Reverse
	Toggle
)
