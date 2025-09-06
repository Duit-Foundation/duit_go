package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

type ColoredBoxAttributes struct {
	*ValueReferenceHolder
	*duit_props.AnimatedPropertyOwner
	*ThemeConsumer
	Color *duit_props.Color `json:"color,omitempty"`
}

func (r *ColoredBoxAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if r.Color == nil {
		return errors.New("color property is required")
	} else {
		if err := r.Color.Validate(); err != nil {
			return err
		}
	}

	return nil
}
