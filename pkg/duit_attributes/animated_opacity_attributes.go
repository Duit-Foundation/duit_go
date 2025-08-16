package duit_attributes

import (
	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type AnimatedOpacityAttributes struct {
	*ValueReferenceHolder
	*animations.ImplicitAnimatable
	*ThemeConsumer
	Opacity duit_utils.Tristate[float32] `json:"opacity,omitempty"`
}

func (a *AnimatedOpacityAttributes) Validate() error {
	if err := a.ImplicitAnimatable.Validate(); err != nil {
		return err
	}

	if err := a.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := a.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if a.Opacity != nil {
		if err := RangeValidator(*a.Opacity, 0, 1); err != nil {
			return err
		}
	}

	return nil
}