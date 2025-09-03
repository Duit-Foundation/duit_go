package duit_attributes

import (
	"errors"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type SliderAttributes[TAction duit_action.Action, TColor duit_props.Color] struct {
	*ValueReferenceHolder
	*ThemeConsumer
	Value                duit_utils.Tristate[float32]                 `json:"value,omitempty"`
	Min                  float32                                      `json:"min,omitempty"`
	Max                  float32                                      `json:"max,omitempty"`
	Divisions            uint32                                       `json:"divisions,omitempty"`
	SecondaryTrackValue  float32                                      `json:"secondaryTrackValue,omitempty"`
	OnChanged            *TAction                                     `json:"onChanged,omitempty"`
	OnChangeStart        *TAction                                     `json:"onChangeStart,omitempty"`
	OnChangeEnd          *TAction                                     `json:"onChangeEnd,omitempty"`
	ActiveColor          TColor                                       `json:"activeColor,omitempty"`
	InactiveColor        TColor                                       `json:"inactiveColor,omitempty"`
	ThumbColor           TColor                                       `json:"thumbColor,omitempty"`
	SecondaryActiveColor TColor                                       `json:"secondaryActiveColor,omitempty"`
	OverlayColor         *duit_props.MaterialStateProperty[TColor] `json:"overlayColor,omitempty"`
	Autofocus            duit_utils.Tristate[bool]                    `json:"autofocus,omitempty"`
	Label                string                                       `json:"label,omitempty"`
	AllowedInteraction   duit_props.SliderInteraction              `json:"allowedInteraction,omitempty"`
}

func (r *SliderAttributes[TAction, TColor]) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if r.Value == nil {
		return errors.New("value property is required")
	}

	return nil
}
