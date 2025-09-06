package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

//TODO: подумать над тем, чтобы добавить обычный BoxConstraints, чтобы было семантически корректно

type ConstrainedBoxAttributes struct {
	*ValueReferenceHolder
	*duit_props.AnimatedPropertyOwner
	*ThemeConsumer
	Constraints *duit_props.BoxConstraints `json:"constraints,omitempty"`
}

func (r *ConstrainedBoxAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if r.Constraints != nil {
		if err := r.Constraints.Validate(); err != nil {
			return err
		}
	} else {
		return errors.New("constraints property is required")
	}

	return nil
}
