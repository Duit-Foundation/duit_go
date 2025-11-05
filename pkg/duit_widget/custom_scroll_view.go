package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

func CustomScrollView(attributes *duit_attributes.CustomScrollViewAttributes, id string, controlled bool, slivers []*duit_core.DuitElementModel) *duit_core.DuitElementModel {
	checkAttributes(attributes)
	return new(duit_core.DuitElementModel).CreateElement(duit_core.CustomScrollView, id, "", attributes, nil, controlled, 2, nil, slivers...)
}
