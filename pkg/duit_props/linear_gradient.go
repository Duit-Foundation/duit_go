package duit_props

import (
	"errors"
	"math"
)

// LinearGradient creates a linear gradient with colors distributed along a line.
//
// LinearGradient is used to create smooth color transitions between multiple colors
// along a straight line. It supports customizable color stops, rotation angle,
// and alignment points for precise control over the gradient appearance.
//
// Flutter documentation: https://api.flutter.dev/flutter/painting/LinearGradient-class.html
type LinearGradient struct {
	Colors        []*Color  `json:"colors"`                  // The colors to be used in the gradient
	RotationAngle float32   `json:"rotationAngle,omitempty"` // Rotation angle in radians
	Stops         []float32 `json:"stops,omitempty"`         // Relative positions of colors along the gradient line
	Begin         Alignment `json:"begin,omitempty"`         // The start point of the gradient
	End           Alignment `json:"end,omitempty"`           // The end point of the gradient
}

// NewLinearGradient creates a new empty LinearGradient structure with default values
func NewLinearGradient() *LinearGradient {
	return &LinearGradient{}
}

// SetColors sets the colors to be used in the gradient
func (r *LinearGradient) SetColors(colors []*Color) *LinearGradient {
	r.Colors = colors
	return r
}

// AddColor adds a single color to the gradient
func (r *LinearGradient) AddColor(color *Color) *LinearGradient {
	r.Colors = append(r.Colors, color)
	return r
}

// SetRotationAngle sets the rotation angle of the gradient in radians
func (r *LinearGradient) SetRotationAngle(rad float32) *LinearGradient {
	r.RotationAngle = rad
	return r
}

// SetRotationAngleDegrees sets the rotation angle of the gradient in degrees
func (r *LinearGradient) SetRotationAngleDegrees(deg float32) *LinearGradient {
	r.RotationAngle = deg * (math.Pi / 180.0)
	return r
}

// SetStops sets the relative positions of colors along the gradient line
func (r *LinearGradient) SetStops(stops []float32) *LinearGradient {
	r.Stops = stops
	return r
}

// AddStop adds a single stop position to the gradient
func (r *LinearGradient) AddStop(stop float32) *LinearGradient {
	r.Stops = append(r.Stops, stop)
	return r
}

// SetBegin sets the start point of the gradient
func (r *LinearGradient) SetBegin(begin Alignment) *LinearGradient {
	r.Begin = begin
	return r
}

// SetEnd sets the end point of the gradient
func (r *LinearGradient) SetEnd(end Alignment) *LinearGradient {
	r.End = end
	return r
}

func (r *LinearGradient) Validate() error {
	if r.Colors == nil {
		return errors.New("colors are required")
	} else {
		for _, color := range r.Colors {
			if color == nil {
				return errors.New("color in gradient colors array must not be nil")
			}
			if err := color.Validate(); err != nil {
				return err
			}
		}
	}

	if len(r.Begin) == 0 {
		return errors.New("begin is required")
	}

	if len(r.End) == 0 {
		return errors.New("end is required")
	}

	return nil
}
