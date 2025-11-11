package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

/*
Example:
	AnimatedPositionedDirectional(
		&duit_attributes.AnimatedPositionedAttributes{
			Top: 120,
			Start: 25,
			ImplicitAnimatable: duit_props.ImplicitAnimatable{
				Duration: 1000,
			},
		},
		"id",
		nil, //child
	)
*/
func AnimatedPositionedDirectional(attributes *duit_attributes.AnimatedPositionedDirectionalAttributes, id string, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	checkAttributes(attributes)
	return new(duit_core.DuitElementModel).CreateElement(duit_core.AnimatedPositionedDirectional, id, "", attributes, nil, true, 1, nil, child)
}