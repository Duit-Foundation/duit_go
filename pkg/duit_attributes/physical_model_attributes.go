package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type PhysicalModelAttributes[TColor duit_props.Color] struct {
	*ValueReferenceHolder
	*duit_animations.AnimatedPropertyOwner
	*ThemeConsumer
	Elevation    duit_utils.Tristate[float32]  `json:"elevation,omitempty"`
	Color        TColor                        `json:"color,omitempty"`
	ShadowColor  TColor                        `json:"shadowColor,omitempty"`
	ClipBehavior duit_props.Clip                `json:"clipBehavior,omitempty"`
	BorderRadius *duit_props.BorderRadius `json:"borderRadius,omitempty"`
	Shape        duit_props.BoxShape      `json:"shape,omitempty"`
}

func (r *PhysicalModelAttributes[TColor]) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if r.Elevation == nil {
		return errors.New("elevation property is required")
	}

	return nil
}