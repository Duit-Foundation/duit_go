package duit_material

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_color"
	duit_decoration "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_flex"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_text_properties"
)

type ButtonStyle[TColor duit_color.Color, TInsets duit_edge_insets.EdgeInsets, TShape duit_decoration.ShapeBorder[TColor]] struct {
	TextStyle         MaterialStateProperty[duit_text_properties.TextStyle[TColor]] `json:"textStyle,omitempty"`
	BackgroundColor   MaterialStateProperty[TColor]                                 `json:"backgroundColor,omitempty"`
	ForegroundColor   MaterialStateProperty[TColor]                                 `json:"foregroundColor,omitempty"`
	OverlayColor      MaterialStateProperty[TColor]                                 `json:"overlayColor,omitempty"`
	ShadowColor       MaterialStateProperty[TColor]                                 `json:"shadowColor,omitempty"`
	IconColor         MaterialStateProperty[TColor]                                 `json:"iconColor,omitempty"`
	Elevation         MaterialStateProperty[float32]                                `json:"elevation,omitempty"`
	IconSize          MaterialStateProperty[float32]                                `json:"iconSize,omitempty"`
	Padding           MaterialStateProperty[TInsets]                                `json:"padding,omitempty"`
	MinimumSize       MaterialStateProperty[duit_flex.Size]                         `json:"minimumSize,omitempty"`
	MaximumSize       MaterialStateProperty[duit_flex.Size]                         `json:"maximumSize,omitempty"`
	Side              MaterialStateProperty[duit_decoration.BorderSide[TColor]]     `json:"side,omitempty"`
	Shape             MaterialStateProperty[TShape]                                 `json:"shape,omitempty"`
	VisualDensity     MaterialStateProperty[duit_flex.VisualDensity]                `json:"visualDensity,omitempty"`
	TapTargetSize     MaterialTapTargetSize                                         `json:"materialTapTargetSize,omitempty"`
	AnimationDuration uint                                                          `json:"animationDuration,omitempty"`
}
