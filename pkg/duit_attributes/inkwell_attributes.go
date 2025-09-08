package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type InkwellAttributes struct {
	*ValueReferenceHolder
	*ThemeConsumer
	OnTap                any                          `json:"onTap,omitempty"`
	OnDoubleTap          any                          `json:"onDoubleTap,omitempty"`
	OnLongPress          any                          `json:"onLongPress,omitempty"`
	OnTapDown            any                          `json:"onTapDown,omitempty"`
	OnTapUp              any                          `json:"onTapUp,omitempty"`
	OnTapCancel          any                          `json:"onTapCancel,omitempty"`
	OnSecondaryTapDown   any                          `json:"onSecondaryTapDown,omitempty"`
	OnSecondaryTapCancel any                          `json:"onSecondaryTapCancel,omitempty"`
	OnSecondaryTap       any                          `json:"onSecondaryTap,omitempty"`
	OnSecondaryTapUp     any                          `json:"onSecondaryTapUp,omitempty"`
	FocusColor           *duit_props.Color            `json:"focusColor,omitempty"`
	HoverColor           *duit_props.Color            `json:"hoverColor,omitempty"`
	HighlightColor       *duit_props.Color            `json:"highlightColor,omitempty"`
	SplashColor          *duit_props.Color            `json:"splashColor,omitempty"`
	OverlayColor         *duit_props.WidgetStateColor `json:"overlayColor,omitempty"`
	Radius               float32                      `json:"radius,omitempty"`
	BorderRadius         *duit_props.BorderRadius     `json:"borderRadius,omitempty"`
	CustomBorder         *duit_props.ShapeBorder      `json:"customBorder,omitempty"`
	EnableFeedback       duit_utils.Tristate[bool]    `json:"enableFeedback,omitempty"`
	ExcludeFromSemantics duit_utils.Tristate[bool]    `json:"excludeFromSemantics,omitempty"`
	Autofocus            duit_utils.Tristate[bool]    `json:"autofocus,omitempty"`
	CanRequestFocus      duit_utils.Tristate[bool]    `json:"canRequestFocus,omitempty"`
	HoverDuration        uint                         `json:"hoverDuration,omitempty"`
}

func (r *InkwellAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	callbacks := []any{
		r.OnTap,
		r.OnDoubleTap,
		r.OnLongPress,
		r.OnTapDown,
		r.OnTapUp,
		r.OnTapCancel,
		r.OnSecondaryTapDown,
		r.OnSecondaryTapCancel,
		r.OnSecondaryTap,
		r.OnSecondaryTapUp,
	}

	for _, callback := range callbacks {
		if err := duit_action.CheckActionType(callback); err != nil {
			return err
		}
	}

	return nil
}
