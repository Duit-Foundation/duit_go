package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

/*
Example:
	AnimatedPositioned(
		&duit_attributes.AnimatedPositionedAttributes{
			Top: 0,
			Right: 0,
			ImplicitAnimatable: duit_animations.ImplicitAnimatable{
				Duration: 1000,
			},
		},
		"id",
		nil, //child
	)
*/
func AnimatedPositioned(attributes *duit_attributes.AnimatedPositionedAttributes, id string, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.AnimatedPositioned, id, "", attributes, nil, true, 1, nil, child)
}