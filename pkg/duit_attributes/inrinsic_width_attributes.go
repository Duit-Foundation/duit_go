package duit_attributes

import "github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"

type IntrinsicWidthAttributes struct {
	*ValueReferenceHolder
	*duit_props.AnimatedPropertyOwner
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
