package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type AnimatedPhysicalModelAttributes struct {
	*ValueReferenceHolder
	*duit_props.ImplicitAnimatable
	*ThemeConsumer
	Elevation          duit_utils.Tristate[float32] `json:"elevation"`
	Color              *duit_props.Color            `json:"color,omitempty"`
	ShadowColor        *duit_props.Color            `json:"shadowColor,omitempty"`
	ClipBehavior       duit_props.Clip              `json:"clipBehavior,omitempty"`
	BorderRadius       *duit_props.BorderRadius     `json:"borderRadius,omitempty"`
	Shape              duit_props.BoxShape          `json:"shape,omitempty"`
	AnimateColor       duit_utils.Tristate[bool]    `json:"animateColor,omitempty"`
	AnimateShadowColor duit_utils.Tristate[bool]    `json:"animateShadowColor,omitempty"`
}

func (r *AnimatedPhysicalModelAttributes) Validate() error {
	if r.ImplicitAnimatable != nil {
		if err := r.ImplicitAnimatable.Validate(); err != nil {
			return err
		}
	} else {
		return errors.New("implicitAnimatable property is required on implicit animated widgets")
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if r.Elevation != nil {
		if *r.Elevation < 0 {
			return errors.New("elevation must be greater than 0")
		}
	} else {
		return errors.New("elevation property is required")
	}

	if r.Color != nil {
		if err := r.Color.Validate(); err != nil {
			return err
		}
	} else {
		return errors.New("color property is required")
	}

	if r.ShadowColor != nil {
		if err := r.ShadowColor.Validate(); err != nil {
			return err
		}
	} else {
		return errors.New("shadowColor property is required")
	}

	if r.BorderRadius != nil {
		if err := r.BorderRadius.Validate(); err != nil {
			return err
		}
	}

	return nil
}
