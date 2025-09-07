package duit_props

// BorderSide represents the style and properties of a single side of a border.
// It defines the color, width, style, and stroke alignment for a border edge.
//
// This is a Go implementation of Flutter's BorderSide class.
// For more information, see: https://api.flutter.dev/flutter/painting/BorderSide-class.html
type BorderSide struct {
	Color       *Color      `json:"color,omitempty"`
	Width       float32     `json:"width,omitempty"`
	Style       BorderStyle `json:"style,omitempty"`
	StrokeAlign float32     `json:"strokeAlign,omitempty"`
}

// Validate checks if the BorderSide properties are valid.
// It validates the color if it's not nil.
func (r *BorderSide) Validate() error {
	if r.Color != nil {
		if err := r.Color.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// SetColor sets the color of the border and returns the BorderSide for method chaining.
func (r *BorderSide) SetColor(color *Color) *BorderSide {
	r.Color = color
	return r
}

// SetWidth sets the width of the border and returns the BorderSide for method chaining.
func (r *BorderSide) SetWidth(width float32) *BorderSide {
	r.Width = width
	return r
}

// SetStyle sets the style of the border and returns the BorderSide for method chaining.
func (r *BorderSide) SetStyle(style BorderStyle) *BorderSide {
	r.Style = style
	return r
}

// SetStrokeAlign sets the stroke alignment of the border and returns the BorderSide for method chaining.
func (r *BorderSide) SetStrokeAlign(strokeAlign float32) *BorderSide {
	r.StrokeAlign = strokeAlign
	return r
}
