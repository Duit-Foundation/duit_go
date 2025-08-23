package duit_attributes

import animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"

type SizedBoxAttributes struct {
	*ValueReferenceHolder
	*animations.AnimatedPropertyOwner
	*ThemeConsumer
	Width  float32 `json:"width,omitempty"`
	Height float32 `json:"height,omitempty"`
}

func (r *SizedBoxAttributes) Validate() error {
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
