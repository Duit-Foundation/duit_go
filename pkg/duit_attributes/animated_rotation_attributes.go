package duit_attributes

import (
	"errors"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type AnimatedRotationAttributes struct {
	*ValueReferenceHolder
	*ThemeConsumer
	*duit_props.ImplicitAnimatable
	Turns         duit_utils.Tristate[float32] `json:"turns,omitempty"`
	Alignment     duit_props.Alignment     `json:"alignment,omitempty"`
	FilterQuality duit_props.FilterQuality  `json:"filterQuality,omitempty"`
}

func (r *AnimatedRotationAttributes) Validate() error {
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

	if r.Turns == nil {
		return errors.New("turns property is required")
	}

	return nil
}
