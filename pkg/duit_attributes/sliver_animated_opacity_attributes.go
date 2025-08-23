package duit_attributes

import (
	"errors"

	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type SliverAnimatedpacityAttributes struct {
	*ValueReferenceHolder
	*ThemeConsumer
	*animations.ImplicitAnimatable
	*SliverProps
	Opacity duit_utils.Tristate[float32] `json:"opacity,omitempty"`
}

func (r *SliverAnimatedpacityAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ImplicitAnimatable.Validate(); err != nil {
		return err
	}

	if r.Opacity == nil {
		return errors.New("opacity property is required")
	} else {
		return RangeValidator(*r.Opacity, 0, 1)
	}
}
