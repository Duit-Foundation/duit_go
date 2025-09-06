package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

type ContainerAttributes struct {
	*ValueReferenceHolder
	*duit_props.AnimatedPropertyOwner
	*ThemeConsumer
	Width                float32                    `json:"width,omitempty"`
	Height               float32                    `json:"height,omitempty"`
	Color                *duit_props.Color          `json:"color,omitempty"`
	ClipBehavior         duit_props.Clip            `json:"clipBehavior,omitempty"`
	Decoration           *duit_props.BoxDecoration  `json:"decoration,omitempty"`
	ForegroundDecoration *duit_props.BoxDecoration  `json:"foregroundDecoration,omitempty"`
	Padding              *duit_props.EdgeInsets     `json:"padding,omitempty"`
	Margin               *duit_props.EdgeInsets     `json:"margin,omitempty"`
	Alignment            duit_props.Alignment       `json:"alignment,omitempty"`
	TransformAlignment   duit_props.Alignment       `json:"transformAlignment,omitempty"`
	Constraints          *duit_props.BoxConstraints `json:"constraints,omitempty"`
}

func (r *ContainerAttributes) Validate() error {
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
