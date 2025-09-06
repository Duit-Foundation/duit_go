package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

type AnimatedPaddingAttributes struct {
	*ValueReferenceHolder
	*duit_props.ImplicitAnimatable
	*ThemeConsumer
	Padding *duit_props.EdgeInsets `json:"padding,omitempty"`
}

func (r *AnimatedPaddingAttributes) Validate() error {
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

	if r.Padding != nil {
		if err := r.Padding.Validate(); err != nil {
			return err
		}
	} else {
		return errors.New("padding property is required")
	}

	return nil
}
