package duit_attributes

import (
	animations "github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_color"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_text_properties"
)

type TextAttributes[T duit_color.Color] struct {
	ValueReferenceHolder
	animations.AnimatedPropertyOwner
	ThemeConsumer
	Data           string                             `json:"data"`
	SemanticsLabel string                             `json:"semanticsLabel,omitempty"`
	TextAlign      duit_text_properties.TextAlign     `json:"textAlign,omitempty"`
	TextDirection  duit_text_properties.TextDirection `json:"textDirection,omitempty"`
	Overflow       duit_text_properties.TextOverflow  `json:"overflow,omitempty"`
	SoftWrap       bool                               `json:"softWrap,omitempty"`
	MaxLines       int                                `json:"maxLines,omitempty"`
	Style          *duit_text_properties.TextStyle[T] `json:"style,omitempty"`
}
