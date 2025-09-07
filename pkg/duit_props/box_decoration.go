package duit_props

// BoxDecoration describes how to paint a box.
// It provides a way to paint a rectangular area with rounded corners, borders, shadows, and gradients.
//
// Used to decorate containers and other widgets with visual styling including
// background colors, borders, shadows, and gradients.
//
// Flutter documentation: https://api.flutter.dev/flutter/painting/BoxDecoration-class.html
type BoxDecoration struct {
	Color               *Color          `json:"color,omitempty"`
	Border              *BorderSide     `json:"border,omitempty"`
	Gradient            *LinearGradient `json:"gradient,omitempty"`
	BoxShadow           []BoxShadow     `json:"boxShadow,omitempty"`
	BorderRadius        *BorderRadius `json:"borderRadius,omitempty"`
	Shape               BoxShape        `json:"shape,omitempty"`
	BackgroundBlendMode BlendMode       `json:"backgroundBlendMode,omitempty"`
}

// NewBoxDecoration creates a new empty BoxDecoration structure with default values
func NewBoxDecoration() *BoxDecoration {
	return &BoxDecoration{}
}

// SetColor sets the background color of the box
func (r *BoxDecoration) SetColor(color *Color) *BoxDecoration {
	r.Color = color
	return r
}

// SetBorder sets the border around the box
func (r *BoxDecoration) SetBorder(border *BorderSide) *BoxDecoration {
	r.Border = border
	return r
}

// SetBorderRadius sets the radius of the rounded corners
func (r *BoxDecoration) SetBorderRadius(radius *BorderRadius) *BoxDecoration {
	r.BorderRadius = radius
	return r
}

// SetShape sets the shape of the box
func (r *BoxDecoration) SetShape(shape BoxShape) *BoxDecoration {
	r.Shape = shape
	return r
}

// SetBoxShadow sets the shadows cast by the box
func (r *BoxDecoration) SetBoxShadow(shadows []BoxShadow) *BoxDecoration {
	r.BoxShadow = shadows
	return r
}

// AddBoxShadow adds a single shadow to the box
func (r *BoxDecoration) AddBoxShadow(shadow BoxShadow) *BoxDecoration {
	r.BoxShadow = append(r.BoxShadow, shadow)
	return r
}

// SetGradient sets the gradient fill for the box
func (r *BoxDecoration) SetGradient(gradient *LinearGradient) *BoxDecoration {
	r.Gradient = gradient
	return r
}

// SetBackgroundBlendMode sets the blend mode for the background
func (r *BoxDecoration) SetBackgroundBlendMode(blendMode BlendMode) *BoxDecoration {
	r.BackgroundBlendMode = blendMode
	return r
}

// Validate checks if the BoxDecoration is valid
func (r *BoxDecoration) Validate() error {
	// Validate border if present
	if r.Border != nil {
		if err := r.Border.Validate(); err != nil {
			return err
		}
	}

	// Validate box shadows if present
	// for _, _ := range r.BoxShadow {
	// 	//TODO: box shadow validation
	// }

	// Validate gradient if present
	if r.Gradient != nil {
		//TODO: gradient validation
	}

	return nil
}
