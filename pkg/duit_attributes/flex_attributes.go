package duit_attributes

import (
	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

type FlexAttributes struct {
	*ValueReferenceHolder
	*animations.AnimatedPropertyOwner
	*ThemeConsumer
	MainAxisAlignment  duit_props.MainAxisAlignment   `json:"mainAxisAlignment,omitempty"`
	MainAxisSize       duit_props.MainAxisSize        `json:"mainAxisSize,omitempty"`
	CrossAxisAlignment duit_props.CrossAxisAlignment `json:"crossAxisAlignment,omitempty"`
	TextDirection      duit_props.TextDirection `json:"textDirection,omitempty"`
	VerticalDirection  duit_props.VerticalDirection        `json:"verticalDirection,omitempty"`
	ClipBehavior       duit_props.Clip                     `json:"clipBehavior,omitempty"`
}

func (r *FlexAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	return nil
}
