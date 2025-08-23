package duit_attributes

import animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"

type IntrinsicWidthAttributes struct {
	*ValueReferenceHolder
	*animations.AnimatedPropertyOwner
	*ThemeConsumer
	StepWidth  float32 `json:"stepWidth,omitempty"`
	StepHeight float32 `json:"stepHeight,omitempty"`
}

func (r *IntrinsicWidthAttributes) Validate() error {
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
