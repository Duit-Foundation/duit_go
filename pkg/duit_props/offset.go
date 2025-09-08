package duit_props

// Offset represents a 2D floating-point offset.
//
// An Offset is a distance from a particular point in 2D space.
// It can be used to represent a position, a displacement, or a size.
//
// The coordinate system has the origin at the top-left corner of the screen,
// with x increasing to the right and y increasing downward.
//
// For more information, see the Flutter documentation:
// https://api.flutter.dev/flutter/dart-ui/Offset-class.html
type Offset struct {
	Dx float32 `json:"dx"` // The horizontal component of the offset
	Dy float32 `json:"dy"` // The vertical component of the offset
}

// NewOffset creates a new Offset with the given horizontal and vertical components.
//
// Parameters:
//   - dx: The horizontal component of the offset
//   - dy: The vertical component of the offset
//
// Returns a pointer to a new Offset instance.
func NewOffset(dx, dy float32) *Offset {
	return &Offset{Dx: dx, Dy: dy}
}

