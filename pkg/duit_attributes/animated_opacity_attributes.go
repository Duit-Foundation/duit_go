package duit_attributes

import (
	"errors"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type AnimatedOpacityAttributes struct {
	*ValueReferenceHolder
	*duit_props.ImplicitAnimatable
	*ThemeConsumer
	Opacity duit_utils.Tristate[float32] `json:"opacity,omitempty"`
}

func (r *AnimatedOpacityAttributes) Validate() error {
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

	if r.Opacity != nil {
		if err := RangeValidator(*r.Opacity, 0, 1); err != nil {
			return err
		}
	}

	return nil
}