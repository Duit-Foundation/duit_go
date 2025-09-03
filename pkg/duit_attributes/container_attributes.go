package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

type ContainerAttributes[TInsets duit_props.EdgeInsets, TColor duit_props.Color] struct {
	*ValueReferenceHolder
	*duit_props.AnimatedPropertyOwner
	*ThemeConsumer
	Width                float32                                `json:"width,omitempty"`
	Height               float32                                `json:"height,omitempty"`
	Color                TColor                                 `json:"color,omitempty"`
	ClipBehavior         duit_props.Clip                         `json:"clipBehavior,omitempty"`
	Decoration           *duit_props.BoxDecoration[TColor] `json:"decoration,omitempty"`
	ForegroundDecoration *duit_props.BoxDecoration[TColor] `json:"foregroundDecoration,omitempty"`
	Padding              TInsets                                `json:"padding,omitempty"`
	Margin               TInsets                                `json:"margin,omitempty"`
	Alignment            duit_props.Alignment               `json:"alignment,omitempty"`
	TransformAlignment   duit_props.Alignment               `json:"transformAlignment,omitempty"`
	Constraints          *duit_props.BoxConstraints              `json:"constraints,omitempty"`
}

func (r *ContainerAttributes[TInsets, TColor]) Validate() error {
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
