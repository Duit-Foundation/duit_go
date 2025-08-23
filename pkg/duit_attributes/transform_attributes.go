package duit_attributes

import (
	"errors"

	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_transform"
)

type TransfromAttributes[T duit_transform.Transform] struct {
	*ValueReferenceHolder
	*animations.AnimatedPropertyOwner
	*ThemeConsumer
	Type duit_transform.TransfromType `json:"type"`
	Data *T                           `json:"data"`
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
