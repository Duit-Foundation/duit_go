package duit_attributes

import (
	"errors"

	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type TextAttributes[T duit_props.Color] struct {
	*ValueReferenceHolder
	*animations.AnimatedPropertyOwner
	*ThemeConsumer
	Data           string                             `json:"data"`
	SemanticsLabel string                             `json:"semanticsLabel,omitempty"`
	TextAlign      duit_props.TextAlign     `json:"textAlign,omitempty"`
	TextDirection  duit_props.TextDirection `json:"textDirection,omitempty"`
	Overflow       duit_props.TextOverflow  `json:"overflow,omitempty"`
	SoftWrap       duit_utils.Tristate[bool]          `json:"softWrap,omitempty"`
	MaxLines       int                                `json:"maxLines,omitempty"`
	Style          *duit_props.TextStyle[T] `json:"style,omitempty"`
}

func (r *TextAttributes[T]) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if len(r.Data) == 0 {
		return errors.New("data property is required")
	}

	return nil
}
