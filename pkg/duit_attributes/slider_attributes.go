package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_color"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_gestures"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_material"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
)

type SliderAttributes[TAction duit_core.Action, TColor duit_color.Color] struct {
	ValueReferenceHolder
	Value                float32                                      `json:"value"`
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
	OverlayColor         *duit_material.MaterialStateProperty[TColor] `json:"overlayColor,omitempty"`
	Autofocus            bool                                         `json:"autofocus,omitempty"`
	Label                string                                       `json:"label,omitempty"`
	AllowedInteraction   duit_gestures.SliderInteraction              `json:"allowedInteraction,omitempty"`
}
