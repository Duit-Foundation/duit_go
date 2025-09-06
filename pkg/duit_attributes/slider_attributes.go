package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type SliderAttributes struct {
	*ValueReferenceHolder
	*ThemeConsumer
	Value                duit_utils.Tristate[float32]                         `json:"value,omitempty"`
	Min                  float32                                              `json:"min,omitempty"`
	Max                  float32                                              `json:"max,omitempty"`
	Divisions            uint32                                               `json:"divisions,omitempty"`
	SecondaryTrackValue  float32                                              `json:"secondaryTrackValue,omitempty"`
	OnChanged            any                                                  `json:"onChanged,omitempty"`
	OnChangeStart        any                                                  `json:"onChangeStart,omitempty"`
	OnChangeEnd          any                                                  `json:"onChangeEnd,omitempty"`
	ActiveColor          *duit_props.Color                                    `json:"activeColor,omitempty"`
	InactiveColor        *duit_props.Color                                    `json:"inactiveColor,omitempty"`
	ThumbColor           *duit_props.Color                                    `json:"thumbColor,omitempty"`
	SecondaryActiveColor *duit_props.Color                                    `json:"secondaryActiveColor,omitempty"`
	OverlayColor         *duit_props.MaterialStateProperty[*duit_props.Color] `json:"overlayColor,omitempty"`
	Autofocus            duit_utils.Tristate[bool]                            `json:"autofocus,omitempty"`
	Label                string                                               `json:"label,omitempty"`
	AllowedInteraction   duit_props.SliderInteraction                         `json:"allowedInteraction,omitempty"`
}

func (r *SliderAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if r.Value == nil {
		return errors.New("value property is required")
	}

	if r.Min >= r.Max {
		return errors.New("min must be less than max")
	}

	actions := []any{
		r.OnChanged,
		r.OnChangeStart,
		r.OnChangeEnd,
	}
	
	for _, action := range actions {
		if err := duit_action.CheckActionType(action); err != nil {
			return err
		}
	}

	return nil
}
