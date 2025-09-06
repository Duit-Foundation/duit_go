package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type ElevatedButtonAttributes[TInsets duit_props.EdgeInsets, TShape duit_props.ShapeBorder, TAction duit_action.Action] struct {
	*ValueReferenceHolder
	*ThemeConsumer
	Autofocus    duit_utils.Tristate[bool]                `json:"autofocus,omitempty"`
	ClipBehavior duit_props.Clip                          `json:"clipBehavior,omitempty"`
	Style        *duit_props.ButtonStyle[TInsets, TShape] `json:"style,omitempty"`
	OnLongPress  *TAction                                 `json:"onLongPress,omitempty"`
}

func (r *ElevatedButtonAttributes[TInsets, TShape, TAction]) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if r.OnLongPress != nil {
		if err := duit_action.CheckActionType(r.OnLongPress); err != nil {
			return err
		}
	}

	return nil
}
