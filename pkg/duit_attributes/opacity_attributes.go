package duit_attributes

import (
	"errors"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type OpacityAttributes struct {
	*ValueReferenceHolder
	*duit_props.AnimatedPropertyOwner
	*ThemeConsumer
	Opacity duit_utils.Tristate[float32] `json:"opacity"`
}

func (r *OpacityAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if r.Opacity == nil {
		return errors.New("opacity property is required")
	} else {
		return RangeValidator(*r.Opacity, 0.0, 1.1)
	}
}
