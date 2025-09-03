package duit_attributes

import "github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"

//TODO: подумать над тем, чтобы добавить обычный BoxConstraints, чтобы было семантически корректно

type ConstrainedBoxAttributes struct {
	*ValueReferenceHolder
	*duit_props.AnimatedPropertyOwner
	*ThemeConsumer
	MinWidth  float32 `json:"minWidth,omitempty"`
	MaxWidth  float32 `json:"maxWidth,omitempty"`
	MinHeight float32 `json:"minHeight,omitempty"`
	MaxHeight float32 `json:"maxHeight,omitempty"`
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

	return nil
}
