package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	duit_decoration "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

func ElevatedButton[TColor duit_props.Color, TInsets duit_props.EdgeInsets, TShape duit_decoration.ShapeBorder[TColor], TAction duit_action.Action](attributes *duit_attributes.ElevatedButtonAttributes[TColor, TInsets, TShape, TAction], id string, action *TAction, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.ElevatedButton, id, "", attributes, action, true, 1, nil, child)
}
