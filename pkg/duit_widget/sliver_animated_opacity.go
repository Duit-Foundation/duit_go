package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
)

/*
Example:

	SliverAnimatedOpacity(
		&duit_attributes.SliverOpacityAttributes{
			Opacity: 0.85,
			Duration: 500,
		},
		"",
		false,
		nil, //child
	)
*/
func SliverAnimatedOpacity(attributes *duit_attributes.SliverAnimatedpacityAttributes, id string, controlled bool, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.SliverAnimatedOpacity, id, "", attributes, nil, controlled, 1, nil, child)
}
