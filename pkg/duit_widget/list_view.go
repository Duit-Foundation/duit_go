package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

func ListView[TAction duit_core.Action, TInsets duit_edge_insets.EdgeInsets, U duit_attributes.ListViewAttributes[TInsets]](attributes *U, id string, controlled bool, action *TAction, children []*duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.ListView, id, "", attributes, action, controlled, 2, nil, children...)
}
