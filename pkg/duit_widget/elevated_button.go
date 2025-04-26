package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_color"
	duit_decoration "github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
)

func ElevatedButton[TColor duit_color.Color, TInsets duit_edge_insets.EdgeInsets, TShape duit_decoration.ShapeBorder[TColor], TAction duit_core.Action](attributes *duit_attributes.ElevatedButtonAttributes[TColor, TInsets, TShape, TAction], id string, action *TAction, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.ElevatedButton, id, "", attributes, action, true, 1, nil, child)
}
