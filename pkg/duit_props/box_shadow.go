package duit_props

// BoxShadow represents a shadow cast by a box.
//
// BoxShadow is used to create a shadow effect around a widget.
// It consists of color, blur radius, spread radius, blur style, and offset.
//
// For more information, see the Flutter documentation:
// https://api.flutter.dev/flutter/dart-ui/BoxShadow-class.html
type BoxShadow struct {
	Color        *Color    `json:"color,omitempty"`        // The color of the shadow
	BlurRadius   float32   `json:"blurRadius,omitempty"`   // The blur radius of the shadow
	SpreadRadius float32   `json:"spreadRadius,omitempty"` // The spread radius of the shadow
	BlurStyle    BlurStyle `json:"blurStyle,omitempty"`    // The blur style of the shadow
	Offset       *Offset   `json:"offset,omitempty"`       // The offset of the shadow
}

// NewBoxShadow creates a new BoxShadow with default values.
//
// Returns a pointer to a new BoxShadow instance with all fields set to their zero values.
func NewBoxShadow() *BoxShadow {
	return &BoxShadow{}
}

// SetColor sets the color of the shadow.
//
// Parameters:
//   - color: The color of the shadow
func (r *BoxShadow) SetColor(color *Color) *BoxShadow {
	r.Color = color
	return r
}

// SetBlurRadius sets the blur radius of the shadow.
//
// Parameters:
//   - radius: The blur radius of the shadow
func (r *BoxShadow) SetBlurRadius(radius float32) *BoxShadow {
	r.BlurRadius = radius
	return r
}

// SetSpreadRadius sets the spread radius of the shadow.
//
// Parameters:
//   - radius: The spread radius of the shadow
func (r *BoxShadow) SetSpreadRadius(radius float32) *BoxShadow {
	r.SpreadRadius = radius
	return r
}

// SetBlurStyle sets the blur style of the shadow.
//
// Parameters:
//   - style: The blur style of the shadow
func (r *BoxShadow) SetBlurStyle(style BlurStyle) *BoxShadow {
	r.BlurStyle = style
	return r
}

// SetOffset sets the offset of the shadow.
//
// Parameters:
//   - offset: The offset of the shadow
func (r *BoxShadow) SetOffset(offset *Offset) *BoxShadow {
	r.Offset = offset
	return r
}

func (r *BoxShadow) Validate() error {
	if r.Color != nil {
		if err := r.Color.Validate(); err != nil {
			return err
		}
	}

	return nil
}
