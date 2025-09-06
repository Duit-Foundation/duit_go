package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type CardAttributes struct {
	*ValueReferenceHolder
	*duit_props.AnimatedPropertyOwner
	*ThemeConsumer
	Color              *duit_props.Color         `json:"color,omitempty"`
	ShadowColor        *duit_props.Color         `json:"shadowColor,omitempty"`
	Elevation          float32                   `json:"elevation,omitempty"`
	BorderOnForeground duit_utils.Tristate[bool] `json:"borderOnForeground,omitempty"`
	SemanticContainer  duit_utils.Tristate[bool] `json:"semanticContainer,omitempty"`
	Margin             *duit_props.EdgeInsets    `json:"margin,omitempty"`
	Shape              *duit_props.ShapeBorder   `json:"shape,omitempty"`
}

// TODO: влупить валиадцию вложенных структур
func (r *CardAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if r.Shape != nil {
		if err := r.Shape.Validate(); err != nil {
			return err
		}
	}

	if r.Margin != nil {
		if err := r.Margin.Validate(); err != nil {
			return err
		}
	}

	return nil
}
