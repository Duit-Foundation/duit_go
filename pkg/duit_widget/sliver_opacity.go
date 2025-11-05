package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

/*
Example:

	SliverOpacity(
		&duit_attributes.SliverOpacityAttributes{
			Opacity: 0.85,
		},
		"",
		false,
		nil, //child
	)
*/
func SliverOpacity(attributes *duit_attributes.SliverOpacityAttributes, id string, controlled bool, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	checkAttributes(attributes)
	return new(duit_core.DuitElementModel).CreateElement(duit_core.SliverOpacity, id, "", attributes, nil, controlled, 1, nil, child)
}
