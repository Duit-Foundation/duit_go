package duit_attributes

import (
	"errors"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

type RichTextAttributes[T duit_props.Color] struct {
	*ValueReferenceHolder
	*ThemeConsumer
	TextSpan       *duit_props.TextSpan[T]   `json:"textSpan"`
	Style          *duit_props.TextStyle[T]  `json:"style,omitempty"`
	TextAlign      duit_props.TextAlign      `json:"textAlign,omitempty"`
	MaxLines       int                                 `json:"maxLines,omitempty"`
	SoftWrap       bool                                `json:"softWrap,omitempty"`
	Overflow       duit_props.TextOverflow   `json:"overflow,omitempty"`
	SemanticsLabel string                              `json:"semanticsLabel,omitempty"`
	SelectionColor T                                   `json:"selectionColor,omitempty"`
	TextDirection  duit_props.TextDirection  `json:"textDirection,omitempty"`
	TextScaler     *duit_props.TextScaler    `json:"textScaleFactor,omitempty"`
	TextWidthBasis duit_props.TextWidthBasis `json:"textWidthBasis,omitempty"`
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
