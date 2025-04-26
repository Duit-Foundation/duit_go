package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_clip"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_color"
	duit_decoration "github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_material"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
)

type ElevatedButtonAttributes[TColor duit_color.Color, TInsets duit_edge_insets.EdgeInsets, TShape duit_decoration.ShapeBorder[TColor], TAction duit_core.Action] struct {
	ValueReferenceHolder
	ThemeConsumer
	Autofocus    bool                                                `json:"autofocus,omitempty"`
	ClipBehavior duit_clip.Clip                                      `json:"clipBehavior,omitempty"`
	Style        *duit_material.ButtonStyle[TColor, TInsets, TShape] `json:"style,omitempty"`
	OnLongPress  *TAction                                            `json:"onLongPress,omitempty"`
}
