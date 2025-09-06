package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

func ElevatedButton[TInsets duit_props.EdgeInsets, TShape duit_props.ShapeBorder, TAction duit_action.Action](attributes *duit_attributes.ElevatedButtonAttributes[TInsets, TShape, TAction], id string, action *TAction, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.ElevatedButton, id, "", attributes, action, true, 1, nil, child)
}
