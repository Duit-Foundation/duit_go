package duit_props

import (
	"encoding/json"
	"errors"
	"strconv"
)

// Radius represents the radius of a circle or ellipse.
// It can be either circular (single value) or elliptical (two values for x and y axes).
//
// This is a Go implementation of Flutter's Radius class.
// For more information, see: https://api.flutter.dev/flutter/dart-ui/Radius-class.html
type Radius struct {
	data any
}

// NewRadiusElliptical creates a new elliptical radius with separate x and y values.
// The value parameter should contain [x, y] values where both must be greater than 0.
func NewRadiusElliptical(value [2]float32) *Radius {
	return &Radius{data: value}
}

// NewRadiusCircular creates a new circular radius with a single value.
// The value must be greater than 0.
func NewRadiusCircular(value float32) *Radius {
	return &Radius{data: value}
}

// Validate checks if the radius values are valid.
// For circular radius, the value must be greater than 0.
// For elliptical radius, both x and y values must be greater than 0.
func (r *Radius) Validate() error {
	switch r.data.(type) {
	case [2]float32:
		for i, v := range r.data.([2]float32) {
			if v <= 0 {
				return errors.New("radius value at index " + strconv.Itoa(i) + " must be greater than 0")
			}
		}

		return nil
	case float32:
		if r.data.(float32) <= 0 {
			return errors.New("radius must be greater than 0")
		}

		return nil
	}

	return nil
}

func (r *Radius) MarshalJSON() ([]byte, error) { return json.Marshal(r.data) }