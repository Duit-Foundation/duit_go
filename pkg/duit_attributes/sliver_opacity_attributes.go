package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type SliverOpacityAttributes struct {
	*ValueReferenceHolder
	*ThemeConsumer
	*duit_props.AnimatedPropertyOwner
	*SliverProps
	Opacity duit_utils.Tristate[float32] `json:"opacity,omitempty"`
}

func (r *SliverOpacityAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if r.Opacity == nil {
		return errors.New("opacity property is required")
	} else {
		return duit_utils.RangeValidator(*r.Opacity, 0, 1)
	}
}
