package duit_attributes

import (
	"errors"

	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type RotatedBoxAttributes struct {
	*ValueReferenceHolder
	*ThemeConsumer
	*animations.AnimatedPropertyOwner
	QuarterTurns duit_utils.Tristate[int] `json:"quarterTurns,omitempty"`
}

func (r *RotatedBoxAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if r.QuarterTurns == nil {
		return errors.New("quarterTurns property is required")
	}

	return nil
}
