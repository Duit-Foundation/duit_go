package duit_props

import "encoding/json"

// RoundedRectangleBorder represents a rounded rectangle border
// Note: Uses any type for Color to avoid import cycles
type RoundedRectangleBorder[TColor Color] struct {
	Type         ShapeBorderType     `json:"type"`
	BorderRadius *BorderRadius       `json:"borderRadius,omitempty"`
	Side         *BorderSide[TColor] `json:"side,omitempty"`
}

func (s *RoundedRectangleBorder[TColor]) MarshalJSON() ([]byte, error) {
	s.Type = ShapeBorderTypeRoundedRectangle
	return json.Marshal(*s)
}

// BeveledRectangleBorder represents a beveled rectangle border
// Note: Uses any type for Color to avoid import cycles
type BeveledRectangleBorder[TColor Color] struct {
	Type         ShapeBorderType     `json:"type"`
	BorderRadius *BorderRadius       `json:"borderRadius,omitempty"`
	Side         *BorderSide[TColor] `json:"side,omitempty"`
}

func (s *BeveledRectangleBorder[TColor]) MarshalJSON() ([]byte, error) {
	s.Type = ShapeBorderTypeBeveledRectangle
	return json.Marshal(*s)
}

// ContinuousRectangleBorder represents a continuous rectangle border
// Note: Uses any type for Color to avoid import cycles
type ContinuousRectangleBorder[TColor Color] struct {
	Type         ShapeBorderType     `json:"type"`
	BorderRadius *BorderRadius       `json:"borderRadius,omitempty"`
	Side         *BorderSide[TColor] `json:"side,omitempty"`
}

func (s *ContinuousRectangleBorder[TColor]) MarshalJSON() ([]byte, error) {
	s.Type = ShapeBorderTypeContinuousRectangle
	return json.Marshal(*s)
}

// CircleBorder represents a circle border
// Note: Uses any type for Color to avoid import cycles
type CircleBorder[TColor Color] struct {
	Type ShapeBorderType     `json:"type"`
	Side *BorderSide[TColor] `json:"side,omitempty"`
}

func (s *CircleBorder[TColor]) MarshalJSON() ([]byte, error) {
	s.Type = ShapeBorderTypeCircle
	return json.Marshal(*s)
}

// StadiumBorder represents a stadium border
// Note: Uses any type for Color to avoid import cycles
type StadiumBorder[TColor Color] struct {
	Type ShapeBorderType     `json:"type"`
	Side *BorderSide[TColor] `json:"side,omitempty"`
}

func (s *StadiumBorder[TColor]) MarshalJSON() ([]byte, error) {
	s.Type = ShapeBorderTypeStadium
	return json.Marshal(*s)
}

// ShapeBorder represents a shape border interface
// Note: Uses any type for Color to avoid import cycles
type ShapeBorder[TColor Color] interface {
	RoundedRectangleBorder[TColor] | ContinuousRectangleBorder[TColor] | StadiumBorder[TColor] | CircleBorder[TColor] | BeveledRectangleBorder[TColor]
}
