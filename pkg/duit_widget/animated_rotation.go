package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

/*
Example:
	AnimatedRotation(
		&duit_attributes.AnimatedRotationAttributes{
			Turns: 3.14,
			ImplicitAnimatable: duit_props.ImplicitAnimatable{
				Duration: 1000,
			},
		},
		"id",
		nil, //child
	)
*/
func AnimatedRotation(attributes *duit_attributes.AnimatedRotationAttributes, id string, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	checkAttributes(attributes)
	return new(duit_core.DuitElementModel).CreateElement(duit_core.AnimatedRotation, id, "", attributes, nil, true, 1, nil, child)
}