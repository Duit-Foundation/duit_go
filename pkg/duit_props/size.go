package duit_props

import (
	"encoding/json"
	"errors"
)

// Size represents a 2D size with width and height dimensions.
//
// Size is used to define dimensions of widgets and other UI elements.
// It supports different construction methods:
//   - NewSize: creates size with explicit width and height
//   - NewSizeSquare: creates square size with equal dimensions
//   - NewSizeFrom: creates size based on a value and main axis
//
// Reference: https://api.flutter.dev/flutter/dart-ui/Size-class.html
type Size struct {
	data any
}

type sizeWH struct {
	Width  float32 `json:"width"`
	Height float32 `json:"height"`
}

type sizeSqare = uint

type sizeFrom struct {
	Value    float32 `json:"value"`
	MainAxis Axis    `json:"mainAxis"`
}

// NewSize creates a new Size with the given width and height
// https://api.flutter.dev/flutter/dart-ui/Size/Size.html
func NewSize(width, height float32) *Size {
	return &Size{data: sizeWH{Width: width, Height: height}}
}

// NewSizeSquare creates a new Size with the given value
// https://api.flutter.dev/flutter/dart-ui/Size/Size.square.html
func NewSizeSquare(value uint) *Size {
	return &Size{data: sizeSqare(value)}
}

// NewSizeFrom creates a new Size with the given value and main axis
// https://api.flutter.dev/flutter/dart-ui/Size/Size.fromHeight.html 
// https://api.flutter.dev/flutter/dart-ui/Size/Size.fromWidth.html
func NewSizeFrom(value float32, mainAxis Axis) *Size {
	return &Size{data: sizeFrom{Value: value, MainAxis: mainAxis}}
}

func (r *Size) Validate() error {
	switch r.data.(type) {
	case sizeWH:
		if r.data.(sizeWH).Width <= 0 || r.data.(sizeWH).Height <= 0 {
			return errors.New("size must be greater or equal to 0")
		}
	case sizeSqare:
		if r.data.(sizeSqare) <= 0 {
			return errors.New("size must be greater or equal to 0")
		}
	case sizeFrom:
		if r.data.(sizeFrom).Value <= 0 {
			return errors.New("size must be greater or equal to 0")
		}
	}

	return nil
}

func (r *Size) MarshalJSON() ([]byte, error) { return json.Marshal(r.data) }
