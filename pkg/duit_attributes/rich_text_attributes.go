package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_color"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_text_properties"
)

type RichTextAttributes[T duit_color.Color] struct {
	ValueReferenceHolder
	ThemeConsumer
	TextSpan       *duit_text_properties.TextSpan[T]   `json:"textSpan"`
	Style          *duit_text_properties.TextStyle[T]  `json:"style,omitempty"`
	TextAlign      duit_text_properties.TextAlign      `json:"textAlign,omitempty"`
	MaxLines       int                                 `json:"maxLines,omitempty"`
	SoftWrap       bool                                `json:"softWrap,omitempty"`
	Overflow       duit_text_properties.TextOverflow   `json:"overflow,omitempty"`
	SemanticsLabel string                              `json:"semanticsLabel,omitempty"`
	SelectionColor T                                   `json:"selectionColor,omitempty"`
	TextDirection  duit_text_properties.TextDirection  `json:"textDirection,omitempty"`
	TextScaler     *duit_text_properties.TextScaler    `json:"textScaleFactor,omitempty"`
	TextWidthBasis duit_text_properties.TextWidthBasis `json:"textWidthBasis,omitempty"`
}
