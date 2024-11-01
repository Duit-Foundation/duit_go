package duit_decoration

import "github.com/Duit-Foundation/duit_go/pkg/duit_attributes/duit_color"

type Border string
type BorderStyle string

const (
	Outline   Border      = "outline"
	Underline Border      = "underline"
	Solid     BorderStyle = "solid"
	None      BorderStyle = "none"
)

type BorderSide[T duit_color.Color] struct {
	Color T           `json:"color,omitempty"`
	Width float32     `json:"width,omitempty"`
	Style BorderStyle `json:"style,omitempty"`
}

type InputBorderOptions[T duit_color.Color] struct {
	GapPadding   float32        `json:"gapPadding,omitempty"`
	BorderRadius float32        `json:"borderRadius,omitempty"`
	BorderSide   *BorderSide[T] `json:"borderSide,omitempty"`
}

type InputBorder[T duit_color.Color] struct {
	Type    Border                 `json:"type"`
	Options *InputBorderOptions[T] `json:"options,omitempty"`
}
