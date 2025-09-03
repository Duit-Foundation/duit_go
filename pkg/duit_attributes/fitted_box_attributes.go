package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

type FittedBoxAttributes struct {
	*ValueReferenceHolder
	*duit_props.AnimatedPropertyOwner
	*ThemeConsumer
	Fit          duit_props.BoxFit         `json:"fit,omitempty"`
	ClipBehavior duit_props.Clip           `json:"clipBehavior,omitempty"`
	Alignment    duit_props.Alignment `json:"alignment,omitempty"`
}

func (r *FittedBoxAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	return nil
}
