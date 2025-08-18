package duit_attributes

import (
	"errors"

	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type AnimatedPositionedAttributes struct {
	*ValueReferenceHolder
	*animations.ImplicitAnimatable
	*ThemeConsumer
	Left   duit_utils.Tristate[float32] `json:"left,omitempty"`
	Top    duit_utils.Tristate[float32] `json:"top,omitempty"`
	Right  duit_utils.Tristate[float32] `json:"right,omitempty"`
	Bottom duit_utils.Tristate[float32] `json:"bottom,omitempty"`
}

func (r *AnimatedPositionedAttributes) Validate() error {
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

	return nil
}
