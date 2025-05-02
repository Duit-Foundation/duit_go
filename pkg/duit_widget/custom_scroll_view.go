package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
)

func CustomScrollView(attributes *duit_attributes.CustomScrollViewAttributes, id string, controlled bool, children []*duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.CustomScrollView, id, "", attributes, nil, controlled, 2, nil, children...)
}
