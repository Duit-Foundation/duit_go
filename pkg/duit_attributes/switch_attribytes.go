package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type SwitchAttributes[TColor duit_props.Color] struct {
	*ValueReferenceHolder
	*ThemeConsumer
	Value                 duit_utils.Tristate[bool]                     `json:"value,omitempty"`
	ActiveColor           TColor                                        `json:"activeColor,omitempty"`
	FocusColor            TColor                                        `json:"focusColor,omitempty"`
	HoverColor            TColor                                        `json:"hoverColor,omitempty"`
	ActiveTrackColor      TColor                                        `json:"activeTrackColor,omitempty"`
	InactiveTrackColor    TColor                                        `json:"inactiveTrackColor,omitempty"`
	OverlayColor          *duit_props.MaterialStateProperty[TColor]  `json:"overlayColor,omitempty"`
	TrackColor            *duit_props.MaterialStateProperty[TColor]  `json:"trackColor,omitempty"`
	ThumbColor            *duit_props.MaterialStateProperty[TColor]  `json:"thumbColor,omitempty"`
	TrackOutlineColor     *duit_props.MaterialStateProperty[TColor]  `json:"trackOutlineColor,omitempty"`
	TrackOutlineWidth     *duit_props.MaterialStateProperty[float32] `json:"trackOutlineWidth,omitempty"`
	SplashRadius          float32                                       `json:"splashRadius,omitempty"`
	MaterialTapTargetSize duit_props.MaterialTapTargetSize           `json:"materialTapTargetSize,omitempty"`
	Autofocus             duit_utils.Tristate[bool]                     `json:"autofocus,omitempty"`
}

func (r *SwitchAttributes[TColor]) Validate() error {
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
