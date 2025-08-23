package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_flex"
)

type AnimatedSlideAttributes struct {
	*ValueReferenceHolder
	*duit_animations.ImplicitAnimatable
	*ThemeConsumer
	Offset *duit_flex.Offset `json:"offset,omitempty"`
}

func (r *AnimatedSlideAttributes) Validate() error {
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

	if r.Offset == nil {
		return errors.New("offset property is required")
	}

	return nil
}
