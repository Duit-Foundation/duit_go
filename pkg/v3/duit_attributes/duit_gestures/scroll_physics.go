package duit_gestures

type ScrollPhysics string

const (
	AlwaysScrollableScrollPhysics ScrollPhysics = "alwaysScrollableScrollPhysics"
	BouncingScrollPhysics         ScrollPhysics = "bouncingScrollPhysics"
	ClampingScrollPhysics         ScrollPhysics = "clampingScrollPhysics"
	FixedExtentScrollPhysics      ScrollPhysics = "fixedExtentScrollPhysics"
	NeverScrollableScrollPhysics  ScrollPhysics = "neverScrollableScrollPhysics"
)
