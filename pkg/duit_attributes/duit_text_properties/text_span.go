package duit_text_properties

import "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_color"

type TextSpan[T duit_color.Color] struct {
	Text     string         `json:"text,omitempty"`
	Style    *TextStyle[T]  `json:"style,omitempty"`
	SpellOut bool           `json:"spellOut,omitempty"`
	Children []*TextSpan[T] `json:"children,omitempty"`
}
