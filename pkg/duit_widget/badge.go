package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

func Badge(attributes *duit_attributes.BadgeAttributes, id string, controlled bool, child *duit_core.DuitElementModel, label *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	children := []*duit_core.DuitElementModel{child, label}
	checkAttributes(attributes)
	return new(duit_core.DuitElementModel).CreateElement(duit_core.Badge, id, "", attributes, nil, controlled, 2, nil, children...)
}
