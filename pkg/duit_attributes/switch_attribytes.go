package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type SwitchAttributes struct {
	*ValueReferenceHolder
	*ThemeConsumer
	Value                 duit_utils.Tristate[bool]        `json:"value,omitempty"`
	ActiveColor           *duit_props.Color                `json:"activeColor,omitempty"`
	FocusColor            *duit_props.Color                `json:"focusColor,omitempty"`
	HoverColor            *duit_props.Color                `json:"hoverColor,omitempty"`
	ActiveTrackColor      *duit_props.Color                `json:"activeTrackColor,omitempty"`
	InactiveTrackColor    *duit_props.Color                `json:"inactiveTrackColor,omitempty"`
	OverlayColor          *duit_props.WidgetStateColor     `json:"overlayColor,omitempty"`
	TrackColor            *duit_props.WidgetStateColor     `json:"trackColor,omitempty"`
	ThumbColor            *duit_props.WidgetStateColor     `json:"thumbColor,omitempty"`
	TrackOutlineColor     *duit_props.WidgetStateColor     `json:"trackOutlineColor,omitempty"`
	TrackOutlineWidth     *duit_props.WidgetStateFloat32   `json:"trackOutlineWidth,omitempty"`
	SplashRadius          float32                          `json:"splashRadius,omitempty"`
	MaterialTapTargetSize duit_props.MaterialTapTargetSize `json:"materialTapTargetSize,omitempty"`
	Autofocus             duit_utils.Tristate[bool]        `json:"autofocus,omitempty"`
}

func (r *SwitchAttributes) Validate() error {
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
