package duit_props

import (
	"encoding/json"
)

// BorderRadius represents the border radius of a rectangle.
// It can specify different radius values for each corner of the rectangle.
//
// This is a Go implementation of Flutter's BorderRadius class.
// For more information, see: https://api.flutter.dev/flutter/painting/BorderRadius-class.html
type BorderRadius struct {
	data any
}

// borderRadiusAll represents a border radius where all corners have the same radius.
type borderRadiusAll struct {
	Radius *Radius `json:"radius"`
}

// borderRadiusOnly represents a border radius where each corner can have a different radius.
type borderRadiusOnly struct {
	TopLeft     *Radius `json:"topLeft,omitempty"`
	TopRight    *Radius `json:"topRight,omitempty"`
	BottomLeft  *Radius `json:"bottomLeft,omitempty"`
	BottomRight *Radius `json:"bottomRight,omitempty"`
}

// borderRadiusVertical represents a border radius where top and bottom corners can have different radii.
type borderRadiusVertical struct {
	Top    *Radius `json:"top,omitempty"`
	Bottom *Radius `json:"bottom,omitempty"`
}

// borderRadiusHorizontal represents a border radius where left and right corners can have different radii.
type borderRadiusHorizontal struct {
	Left  *Radius `json:"left,omitempty"`
	Right *Radius `json:"right,omitempty"`
}

// NewBorderRadiusAll creates a new BorderRadius where all corners have the same radius.
func NewBorderRadiusAll(radius *Radius) *BorderRadius {
	return &BorderRadius{data: borderRadiusAll{Radius: radius}}
}

// NewBorderRadiusOnly creates a new BorderRadius where each corner can have a different radius.
// Any of the parameters can be nil to indicate no radius for that corner.
func NewBorderRadiusOnly(topLeft, topRight, bottomLeft, bottomRight *Radius) *BorderRadius {
	return &BorderRadius{data: borderRadiusOnly{
		TopLeft:     topLeft,
		TopRight:    topRight,
		BottomLeft:  bottomLeft,
		BottomRight: bottomRight,
	}}
}

// NewBorderRadiusVertical creates a new BorderRadius where top and bottom corners can have different radii.
// Either parameter can be nil to indicate no radius for that side.
func NewBorderRadiusVertical(top, bottom *Radius) *BorderRadius {
	return &BorderRadius{data: borderRadiusVertical{
		Top:    top,
		Bottom: bottom,
	}}
}

// NewBorderRadiusHorizontal creates a new BorderRadius where left and right corners can have different radii.
// Either parameter can be nil to indicate no radius for that side.
func NewBorderRadiusHorizontal(left, right *Radius) *BorderRadius {
	return &BorderRadius{data: borderRadiusHorizontal{
		Left:  left,
		Right: right,
	}}
}

// Validate checks if all radius values in the BorderRadius are valid.
// It recursively validates all Radius values contained within the BorderRadius.
func (r *BorderRadius) Validate() error {
	if r == nil {
		return nil
	}

	switch data := r.data.(type) {
	case borderRadiusAll:
		if data.Radius != nil {
			return data.Radius.Validate()
		}
	case borderRadiusOnly:
		if data.TopLeft != nil {
			if err := data.TopLeft.Validate(); err != nil {
				return err
			}
		}
		if data.TopRight != nil {
			if err := data.TopRight.Validate(); err != nil {
				return err
			}
		}
		if data.BottomLeft != nil {
			if err := data.BottomLeft.Validate(); err != nil {
				return err
			}
		}
		if data.BottomRight != nil {
			if err := data.BottomRight.Validate(); err != nil {
				return err
			}
		}
	case borderRadiusVertical:
		if data.Top != nil {
			if err := data.Top.Validate(); err != nil {
				return err
			}
		}
		if data.Bottom != nil {
			if err := data.Bottom.Validate(); err != nil {
				return err
			}
		}
	case borderRadiusHorizontal:
		if data.Left != nil {
			if err := data.Left.Validate(); err != nil {
				return err
			}
		}
		if data.Right != nil {
			if err := data.Right.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *BorderRadius) MarshalJSON() ([]byte, error) { return json.Marshal(r.data) }
