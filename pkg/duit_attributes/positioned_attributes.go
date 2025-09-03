package duit_attributes

import "github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"

type PositionedAttributes struct {
	*ValueReferenceHolder
	*duit_props.AnimatedPropertyOwner
	*ThemeConsumer
	Left   float32 `json:"left,omitempty"`
	Top    float32 `json:"top,omitempty"`
	Right  float32 `json:"right,omitempty"`
	Bottom float32 `json:"bottom,omitempty"`
}

func (r *PositionedAttributes) Validate() error {
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
