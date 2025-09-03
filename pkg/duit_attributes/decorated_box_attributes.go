package duit_attributes

import (
	"errors"

	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

type DecoratedBoxAttributes[T duit_props.Color] struct {
	*ValueReferenceHolder
	*animations.AnimatedPropertyOwner
	*ThemeConsumer
	Decoration *duit_props.BoxDecoration[T] `json:"decoration,omitempty"`
}

func (r *DecoratedBoxAttributes[T]) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if r.Decoration == nil {
		//TODO: добавить валидацию для декораций позже
		return errors.New("decoration property is required ")
	}

	return nil
}
