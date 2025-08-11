package duit_decoration

import (
	"encoding/json"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_color"
)

type ShapeBorderType string

const (
	RoundedRectangleBorderType    ShapeBorderType = "RoundedRectangleBorder"
	CircleBorderType              ShapeBorderType = "CircleBorder"
	StadiumBorderType             ShapeBorderType = "StadiumBorder"
	BeveledRectangleBorderType    ShapeBorderType = "BeveledRectangleBorder"
	ContinuousRectangleBorderType ShapeBorderType = "ContinuousRectangleBorder"
)

type RoundedRectangleBorder[TColor duit_color.Color] struct {
	Type         ShapeBorderType     `json:"type"`
	BorderRadius *BorderRadius       `json:"borderRadius,omitempty"`
	Side         *BorderSide[TColor] `json:"side,omitempty"`
}

func (s *RoundedRectangleBorder[TColor]) MarshalJSON() ([]byte, error) {
	s.Type = RoundedRectangleBorderType
	return json.Marshal(*s)
}

type BeveledRectangleBorder[TColor duit_color.Color] struct {
	Type         ShapeBorderType     `json:"type"`
	BorderRadius *BorderRadius       `json:"borderRadius,omitempty"`
	Side         *BorderSide[TColor] `json:"side,omitempty"`
}

func (s *BeveledRectangleBorder[TColor]) MarshalJSON() ([]byte, error) {
	s.Type = BeveledRectangleBorderType
	return json.Marshal(*s)
}

type ContinuousRectangleBorder[TColor duit_color.Color] struct {
	Type         ShapeBorderType     `json:"type"`
	BorderRadius *BorderRadius       `json:"borderRadius,omitempty"`
	Side         *BorderSide[TColor] `json:"side,omitempty"`
}

func (s *ContinuousRectangleBorder[TColor]) MarshalJSON() ([]byte, error) {
	s.Type = ContinuousRectangleBorderType
	return json.Marshal(*s)
}

type CircleBorder[TColor duit_color.Color] struct {
	Type ShapeBorderType     `json:"type"`
	Side *BorderSide[TColor] `json:"side,omitempty"`
}

func (s *CircleBorder[TColor]) MarshalJSON() ([]byte, error) {
	s.Type = CircleBorderType
	return json.Marshal(*s)
}

type StadiumBorder[TColor duit_color.Color] struct {
	Type ShapeBorderType     `json:"type"`
	Side *BorderSide[TColor] `json:"side,omitempty"`
}

func (s *StadiumBorder[TColor]) MarshalJSON() ([]byte, error) {
	s.Type = StadiumBorderType
	return json.Marshal(*s)
}

type ShapeBorder[TColor duit_color.Color] interface {
	RoundedRectangleBorder[TColor] | ContinuousRectangleBorder[TColor] | StadiumBorder[TColor] | CircleBorder[TColor] | BeveledRectangleBorder[TColor]
}
