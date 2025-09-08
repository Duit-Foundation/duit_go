package duit_props

// VisualDensity represents the visual density of a widget.
// It is used to control the density of a widget.
//
// This is a Go implementation of Flutter's VisualDensity class.
// For more information, see: https://api.flutter.dev/flutter/material/VisualDensity-class.html
type VisualDensity struct {
	Vertical   float32 `json:"vertical,omitempty"`
	Horizontal float32 `json:"horizontal,omitempty"`
}

// NewVisualDensity creates a new instance of VisualDensity
func NewVisualDensity() *VisualDensity {
	return &VisualDensity{}
}

// SetVertical sets the vertical density
func (r *VisualDensity) SetVertical(value float32) *VisualDensity {
	r.Vertical = value
	return r
}

// SetHorizontal sets the horizontal density
func (r *VisualDensity) SetHorizontal(value float32) *VisualDensity {
	r.Horizontal = value
	return r
}