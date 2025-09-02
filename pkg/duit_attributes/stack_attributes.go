package duit_attributes

import (
	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_flex"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_text_properties"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

type StackAttributes struct {
	*ValueReferenceHolder
	*animations.AnimatedPropertyOwner
	*ThemeConsumer
	TextDirection duit_text_properties.TextDirection `json:"textDirection,omitempty"`
	ClipBehavior  duit_props.Clip                     `json:"clipBehavior,omitempty"`
	Alignment     duit_props.Alignment           `json:"alignment,omitempty"`
	Fit           duit_flex.StackFit                 `json:"fit,omitempty"`
}

func (r *StackAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	return nil
}
