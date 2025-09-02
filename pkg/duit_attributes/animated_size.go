package duit_attributes

import (
	"errors"

	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

type AnimatedSizeAttributes struct {
	*ValueReferenceHolder
	*animations.ImplicitAnimatable
	*ThemeConsumer
	ClipBehavior    duit_props.Clip           `json:"clipBehavior,omitempty"`
	Alignment       duit_props.Alignment `json:"alignment,omitempty"`
	ReverseDuration uint                     `json:"reverseDuration,omitempty"`
}

func (r *AnimatedSizeAttributes) Validate() error {
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
