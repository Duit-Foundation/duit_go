package duit_attributes

import (
	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	duit_decoration "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type CardAttributes[TColor duit_props.Color, TInsets duit_props.EdgeInsets, TShape duit_decoration.ShapeBorder[TColor]] struct {
	*ValueReferenceHolder
	*animations.AnimatedPropertyOwner
	*ThemeConsumer
	Color              TColor                    `json:"color,omitempty"`
	ShadowColor        TColor                    `json:"shadowColor,omitempty"`
	Elevation          float32                   `json:"elevation,omitempty"`
	BorderOnForeground duit_utils.Tristate[bool] `json:"borderOnForeground,omitempty"`
	SemanticContainer  duit_utils.Tristate[bool] `json:"semanticContainer,omitempty"`
	Margin             TInsets                   `json:"margin,omitempty"`
	Shape              *TShape                   `json:"shape,omitempty"`
}
//TODO: влупить валиадцию вложенных структур
func (r *CardAttributes[TColor, TInsets, TShape]) Validate() error {
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
