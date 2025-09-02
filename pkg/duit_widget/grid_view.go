package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

func GridView[TAction duit_action.Action, TInsets duit_props.EdgeInsets, U duit_attributes.GridViewAttributes[TInsets]](attributes *U, id string, controlled bool, action *TAction, children []*duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.GridView, id, "", attributes, action, controlled, 2, nil, children...)
}
