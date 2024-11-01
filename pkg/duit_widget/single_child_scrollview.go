package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
)

func SingleChildScrollView[T duit_edge_insets.EdgeInsets](attributes *duit_attributes.SingleChildScrollViewAttributes[T], id string, controlled bool, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.SingleChildScrollView, id, "", attributes, nil, controlled, 1, nil, child)
}