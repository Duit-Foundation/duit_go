package duit_attributes

import (
	"errors"

	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_painting"
)

type BackdropFilterAttributes[T duit_painting.ImageFilter] struct {
	*ValueReferenceHolder
	*animations.AnimatedPropertyOwner
	*ThemeConsumer
	Filter    *T                      `json:"filter,omitempty"`
	BlendMode duit_painting.BlendMode `json:"blendMode,omitempty"`
}

func (r *BackdropFilterAttributes[T]) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if r.Filter == nil {
		//TODO: добавить валидацию для фильтров позже
		return errors.New("filter property is required")
	}

	if r.BlendMode == "" {
		return errors.New("blendMode property is required")
	}

	return nil
}
