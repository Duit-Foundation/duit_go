package duit_props

// BoxDecoration is a description of how to paint a box.
//
// Ref: https://api.flutter.dev/flutter/painting/BoxDecoration-class.html
type BoxDecoration[T Color] struct {
	Color        T                  `json:"color,omitempty"`
	Boder        *BorderSide[T]     `json:"border,omitempty"`
	BorderRadius float32            `json:"borderRadius,omitempty"`
	Shape        BoxShape           `json:"shape,omitempty"`
	BoxShadow    []BoxShadow[T]     `json:"boxShadow,omitempty"`
	Gradient     *LinearGradient[T] `json:"gradient,omitempty"`
}
