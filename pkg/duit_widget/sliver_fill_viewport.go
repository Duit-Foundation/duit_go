package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

/*
Example:
	SliverFillViewport(
		&duit_attributes.SliverFillViewportAttributes{
			ViewportFraction: 0.75,
		},
		"id",
		false,
		nil, //children
	)
*/
func SliverFillViewport(attributes *duit_attributes.SliverFillViewportAttributes, id string, controlled bool, children []*duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.SliverFillViewport, id, "", attributes, nil, controlled, 2, nil, children...)
}
