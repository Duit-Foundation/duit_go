package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_text_properties"
)

type RichTextAttributes[T duit_props.Color] struct {
	*ValueReferenceHolder
	*ThemeConsumer
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

func (r *RichTextAttributes[T]) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if r.TextSpan == nil {
		return errors.New("textSpan property is required")
	}

	return nil
}
