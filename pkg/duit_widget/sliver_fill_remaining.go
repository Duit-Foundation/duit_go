package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

/*
Example:

	SliverFillRemaining(
		&duit_attributes.SliverFillRemainingAttributes{
			FillOverscroll: true,
		},
		"",
		false,
		nil, //child
	)
*/
func SliverFillRemaining(attributes *duit_attributes.SliverFillRemainingAttributes, id string, controlled bool, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	checkAttributes(attributes)
	return new(duit_core.DuitElementModel).CreateElement(duit_core.SliverFillRemaining, id, "", attributes, nil, controlled, 1, nil, child)
}
