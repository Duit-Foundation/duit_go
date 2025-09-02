package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	duit_decoration "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type ElevatedButtonAttributes[TColor duit_props.Color, TInsets duit_props.EdgeInsets, TShape duit_decoration.ShapeBorder[TColor], TAction duit_action.Action] struct {
	*ValueReferenceHolder
	*ThemeConsumer
	Autofocus    duit_utils.Tristate[bool]                           `json:"autofocus,omitempty"`
	ClipBehavior duit_props.Clip                                      `json:"clipBehavior,omitempty"`
	Style        *duit_props.ButtonStyle[TColor, TInsets, TShape] `json:"style,omitempty"`
	OnLongPress  *TAction                                            `json:"onLongPress,omitempty"`
}

func (r *ElevatedButtonAttributes[TColor, TInsets, TShape, TAction]) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	return nil
}
