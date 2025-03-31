package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_color"
	duit_decoration "github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_material"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
)

type InkwellAttributes[TColor duit_color.Color, TAction duit_core.Action, TShape duit_decoration.ShapeBorder[TColor]] struct {
	ValueReferenceHolder
	ThemeConsumer
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
	EnableFeedback       *bool                                        `json:"enableFeedback,omitempty"`
	ExcludeFromSemantics *bool                                        `json:"excludeFromSemantics,omitempty"`
	Autofocus            *bool                                        `json:"autofocus,omitempty"`
	CanRequestFocus      *bool                                        `json:"canRequestFocus,omitempty"`
	HoverDuration        uint                                         `json:"hoverDuration,omitempty"`
}
