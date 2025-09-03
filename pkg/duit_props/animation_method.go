package duit_props

type AnimationMethod = uint8

const (
	AnimationMethodForward AnimationMethod = iota
	AnimationMethodRepeat
	AnimationMethodReverse
	AnimationMethodToggle
)
