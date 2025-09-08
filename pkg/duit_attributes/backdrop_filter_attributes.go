package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

type BackdropFilterAttributes struct {
	*ValueReferenceHolder
	*duit_props.AnimatedPropertyOwner
	*ThemeConsumer
	Filter    *duit_props.ImageFilter `json:"filter,omitempty"`
	BlendMode duit_props.BlendMode    `json:"blendMode,omitempty"`
}

func (r *BackdropFilterAttributes) Validate() error {
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
