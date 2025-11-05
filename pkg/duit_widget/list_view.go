package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

func ListView(attributes *duit_attributes.ListViewAttributes, id string, controlled bool, action any, children []*duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.ListView, id, "", attributes, action, controlled, 2, nil, children...)
}
