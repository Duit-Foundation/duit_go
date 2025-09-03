package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

type TransfromAttributes[T duit_props.Transform] struct {
	*ValueReferenceHolder
	*duit_props.AnimatedPropertyOwner
	*ThemeConsumer
	Type duit_props.TransformType `json:"type"`
	Data *T                       `json:"data"`
}

func (r *TransfromAttributes[T]) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if len(r.Type) == 0 {
		return errors.New("type property is required")
	}

	if r.Data == nil {
		return errors.New("data property is required")
	}

	return nil
}
