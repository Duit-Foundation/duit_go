package duit_props

// BoxDecoration is a description of how to paint a box.
//
// Ref: https://api.flutter.dev/flutter/painting/BoxDecoration-class.html
type BoxDecoration struct {
	Color        *Color           `json:"color,omitempty"`
	Boder        *BorderSide     `json:"border,omitempty"`
	BorderRadius float32         `json:"borderRadius,omitempty"`
	Shape        BoxShape        `json:"shape,omitempty"`
	BoxShadow    []BoxShadow     `json:"boxShadow,omitempty"`
	Gradient     *LinearGradient `json:"gradient,omitempty"`
}
