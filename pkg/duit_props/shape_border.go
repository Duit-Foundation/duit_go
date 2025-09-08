package duit_props

import "encoding/json"

type ShapeBorder struct {
	data any
}

func (r *ShapeBorder) Validate() error {
	if r == nil {
		return nil
	}

	return nil
}

func (r *ShapeBorder) MarshalJSON() ([]byte, error) { return json.Marshal(r.data) }

type roundedRectangleBorder struct {
	Type         ShapeBorderType `json:"type"`
	BorderRadius *BorderRadius   `json:"borderRadius,omitempty"`
	Side         *BorderSide     `json:"side,omitempty"`
}

func NewRoundedRectangleBorder(borderRadius *BorderRadius, side *BorderSide) *ShapeBorder {
	return &ShapeBorder{data: roundedRectangleBorder{
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

func NewBeveledRectangleBorder(borderRadius *BorderRadius, side *BorderSide) *ShapeBorder {
	return &ShapeBorder{data: beveledRectangleBorder{
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

func NewContinuousRectangleBorder(borderRadius *BorderRadius, side *BorderSide) *ShapeBorder {
	return &ShapeBorder{data: continuousRectangleBorder{
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

func NewCircleBorder(side *BorderSide, eccentricity float32) *ShapeBorder {
	return &ShapeBorder{data: circleBorder{
		Type:         ShapeBorderTypeCircle,
		Side:         side,
		Eccentricity: eccentricity,
	}}
}

type stadiumBorder struct {
	Type ShapeBorderType `json:"type"`
	Side *BorderSide     `json:"side,omitempty"`
}

func NewStadiumBorder(side *BorderSide) *ShapeBorder {
	return &ShapeBorder{data: stadiumBorder{
		Type: ShapeBorderTypeStadium,
		Side: side,
	}}
}
