package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	duit_decoration "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_material"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type InkwellAttributes[TColor duit_props.Color, TAction duit_action.Action, TShape duit_decoration.ShapeBorder[TColor]] struct {
	*ValueReferenceHolder
	*ThemeConsumer
	OnTap                *TAction                                     `json:"onTap,omitempty"`
	OnDoubleTap          *TAction                                     `json:"onDoubleTap,omitempty"`
	OnLongPress          *TAction                                     `json:"onLongPress,omitempty"`
	OnTapDown            *TAction                                     `json:"onTapDown,omitempty"`
	OnTapUp              *TAction                                     `json:"onTapUp,omitempty"`
	OnTapCancel          *TAction                                     `json:"onTapCancel,omitempty"`
	OnSecondaryTapDown   *TAction                                     `json:"onSecondaryTapDown,omitempty"`
	OnSecondaryTapCancel *TAction                                     `json:"onSecondaryTapCancel,omitempty"`
	OnSecondaryTap       *TAction                                     `json:"onSecondaryTap,omitempty"`
	OnSecondaryTapUp     *TAction                                     `json:"onSecondaryTapUp,omitempty"`
	FocusColor           TColor                                       `json:"focusColor,omitempty"`
	HoverColor           TColor                                       `json:"hoverColor,omitempty"`
	HighlightColor       TColor                                       `json:"highlightColor,omitempty"`
	SplashColor          TColor                                       `json:"splashColor,omitempty"`
	OverlayColor         *duit_material.MaterialStateProperty[TColor] `json:"overlayColor,omitempty"`
	Radius               float32                                      `json:"radius,omitempty"`
	BorderRadius         *duit_decoration.BorderRadius                `json:"borderRadius,omitempty"`
	CustomBorder         *TShape                                      `json:"customBorder,omitempty"`
	EnableFeedback       duit_utils.Tristate[bool]                    `json:"enableFeedback,omitempty"`
	ExcludeFromSemantics duit_utils.Tristate[bool]                    `json:"excludeFromSemantics,omitempty"`
	Autofocus            duit_utils.Tristate[bool]                    `json:"autofocus,omitempty"`
	CanRequestFocus      duit_utils.Tristate[bool]                    `json:"canRequestFocus,omitempty"`
	HoverDuration        uint                                         `json:"hoverDuration,omitempty"`
}

func (r *InkwellAttributes[TColor, TAction, TShape]) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	callbacks := []*TAction{
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
