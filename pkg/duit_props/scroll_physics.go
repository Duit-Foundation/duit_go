package duit_props

type ScrollPhysics = string

const (
	ScrollPhysicsAlwaysScrollable ScrollPhysics = "alwaysScrollableScrollPhysics"
	ScrollPhysicsBouncing         ScrollPhysics = "bouncingScrollPhysics"
	ScrollPhysicsClamping         ScrollPhysics = "clampingScrollPhysics"
	ScrollPhysicsFixedExtent      ScrollPhysics = "fixedExtentScrollPhysics"
	ScrollPhysicsNeverScrollable  ScrollPhysics = "neverScrollableScrollPhysics"
)
