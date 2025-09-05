package duit_props

import "encoding/json"

// RoundedRectangleBorder represents a rounded rectangle border
// Note: Uses any type for Color to avoid import cycles
type RoundedRectangleBorder struct {
	Type         ShapeBorderType `json:"type"`
	BorderRadius *BorderRadius   `json:"borderRadius,omitempty"`
	Side         *BorderSide     `json:"side,omitempty"`
}

func (s *RoundedRectangleBorder) MarshalJSON() ([]byte, error) {
	s.Type = ShapeBorderTypeRoundedRectangle
	return json.Marshal(*s)
}

// BeveledRectangleBorder represents a beveled rectangle border
// Note: Uses any type for Color to avoid import cycles
type BeveledRectangleBorder struct {
	Type         ShapeBorderType `json:"type"`
	BorderRadius *BorderRadius   `json:"borderRadius,omitempty"`
	Side         *BorderSide     `json:"side,omitempty"`
}

func (s *BeveledRectangleBorder) MarshalJSON() ([]byte, error) {
	s.Type = ShapeBorderTypeBeveledRectangle
	return json.Marshal(*s)
}

// ContinuousRectangleBorder represents a continuous rectangle border
// Note: Uses any type for Color to avoid import cycles
type ContinuousRectangleBorder struct {
	Type         ShapeBorderType `json:"type"`
	BorderRadius *BorderRadius   `json:"borderRadius,omitempty"`
	Side         *BorderSide     `json:"side,omitempty"`
}

func (s *ContinuousRectangleBorder) MarshalJSON() ([]byte, error) {
	s.Type = ShapeBorderTypeContinuousRectangle
	return json.Marshal(*s)
}

// CircleBorder represents a circle border
// Note: Uses any type for Color to avoid import cycles
type CircleBorder struct {
	Type ShapeBorderType `json:"type"`
	Side *BorderSide     `json:"side,omitempty"`
}

func (s *CircleBorder) MarshalJSON() ([]byte, error) {
	s.Type = ShapeBorderTypeCircle
	return json.Marshal(*s)
}

// StadiumBorder represents a stadium border
// Note: Uses any type for Color to avoid import cycles
type StadiumBorder struct {
	Type ShapeBorderType `json:"type"`
	Side *BorderSide     `json:"side,omitempty"`
}

func (s *StadiumBorder) MarshalJSON() ([]byte, error) {
	s.Type = ShapeBorderTypeStadium
	return json.Marshal(*s)
}

// ShapeBorder represents a shape border interface
// Note: Uses any type for Color to avoid import cycles
type ShapeBorder interface {
	RoundedRectangleBorder | ContinuousRectangleBorder | StadiumBorder | CircleBorder | BeveledRectangleBorder
}
