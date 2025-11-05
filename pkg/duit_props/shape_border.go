package duit_props

import "encoding/json"

// ShapeBorder represents the shape of a border for widgets.
// It defines how to draw the outline of a widget and provides methods
// for creating different border shapes like rounded rectangles, circles,
// and stadiums. The border can have different styles and properties
// depending on the specific shape type used.
//
// This is a Go implementation of Flutter's ShapeBorder class.
// For more information, see: https://api.flutter.dev/flutter/painting/ShapeBorder-class.html
type ShapeBorder struct {
	data any
}

func (r *ShapeBorder) Validate() error {
	if r == nil {
		return nil
	}
	switch r.data.(type) {
	case *roundedRectangleBorder:
		return r.data.(*roundedRectangleBorder).Validate()
	case *beveledRectangleBorder:
		return r.data.(*beveledRectangleBorder).Validate()
	case *continuousRectangleBorder:
		return r.data.(*continuousRectangleBorder).Validate()
	case *circleBorder:
		return r.data.(*circleBorder).Validate()
	case *stadiumBorder:
		return r.data.(*stadiumBorder).Validate()
	}
	return nil
}

func (r *ShapeBorder) MarshalJSON() ([]byte, error) { return json.Marshal(r.data) }

type roundedRectangleBorder struct {
	Type         ShapeBorderType `json:"type"`
	BorderRadius *BorderRadius   `json:"borderRadius,omitempty"`
	Side         *BorderSide     `json:"side,omitempty"`
}

func (r *roundedRectangleBorder) Validate() error {
	if r.BorderRadius != nil {
		if err := r.BorderRadius.Validate(); err != nil {
			return err
		}
	}
	if r.Side != nil {
		if err := r.Side.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// https://api.flutter.dev/flutter/painting/RoundedRectangleBorder-class.html
func NewRoundedRectangleBorder(borderRadius *BorderRadius, side *BorderSide) *ShapeBorder {
	return &ShapeBorder{data: &roundedRectangleBorder{
		Type:         ShapeBorderTypeRoundedRectangle,
		BorderRadius: borderRadius,
		Side:         side,
	}}
}

type beveledRectangleBorder struct {
	Type         ShapeBorderType `json:"type"`
	BorderRadius *BorderRadius   `json:"borderRadius,omitempty"`
	Side         *BorderSide     `json:"side,omitempty"`
}

func (r *beveledRectangleBorder) Validate() error {
	if r.BorderRadius != nil {
		if err := r.BorderRadius.Validate(); err != nil {
			return err
		}
	}
	if r.Side != nil {
		if err := r.Side.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// https://api.flutter.dev/flutter/painting/BeveledRectangleBorder-class.html
func NewBeveledRectangleBorder(borderRadius *BorderRadius, side *BorderSide) *ShapeBorder {
	return &ShapeBorder{data: &beveledRectangleBorder{
		Type:         ShapeBorderTypeBeveledRectangle,
		BorderRadius: borderRadius,
		Side:         side,
	}}
}

type continuousRectangleBorder struct {
	Type         ShapeBorderType `json:"type"`
	BorderRadius *BorderRadius   `json:"borderRadius,omitempty"`
	Side         *BorderSide     `json:"side,omitempty"`
}

func (r *continuousRectangleBorder) Validate() error {
	if r.BorderRadius != nil {
		if err := r.BorderRadius.Validate(); err != nil {
			return err
		}
	}
	if r.Side != nil {
		if err := r.Side.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// https://api.flutter.dev/flutter/painting/ContinuousRectangleBorder-class.html
func NewContinuousRectangleBorder(borderRadius *BorderRadius, side *BorderSide) *ShapeBorder {
	return &ShapeBorder{data: &continuousRectangleBorder{
		Type:         ShapeBorderTypeContinuousRectangle,
		BorderRadius: borderRadius,
		Side:         side,
	}}
}

type circleBorder struct {
	Type         ShapeBorderType `json:"type"`
	Side         *BorderSide     `json:"side,omitempty"`
	Eccentricity float32         `json:"eccentricity,omitempty"`
}

func (r *circleBorder) Validate() error {
	if r.Side != nil {
		if err := r.Side.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// https://api.flutter.dev/flutter/painting/CircleBorder-class.html
func NewCircleBorder(side *BorderSide, eccentricity float32) *ShapeBorder {
	return &ShapeBorder{data: &circleBorder{
		Type:         ShapeBorderTypeCircle,
		Side:         side,
		Eccentricity: eccentricity,
	}}
}

type stadiumBorder struct {
	Type ShapeBorderType `json:"type"`
	Side *BorderSide     `json:"side,omitempty"`
}

func (r *stadiumBorder) Validate() error {
	if r.Side != nil {
		if err := r.Side.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// https://api.flutter.dev/flutter/painting/StadiumBorder-class.html
func NewStadiumBorder(side *BorderSide) *ShapeBorder {
	return &ShapeBorder{data: &stadiumBorder{
		Type: ShapeBorderTypeStadium,
		Side: side,
	}}
}
