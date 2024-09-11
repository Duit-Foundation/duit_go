package duit_widget

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_edge_insets"
	"github.com/lesleysin/duit_go/pkg/duit_core"
)

func ListView[TInsets duit_edge_insets.EdgeInsets, U duit_attributes.ListViewAttributes[TInsets]](attributes *U, id string, controlled bool, action *duit_core.Action, children []*duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.ListView, id, "", attributes, action, controlled, 2, nil, children...)
}
