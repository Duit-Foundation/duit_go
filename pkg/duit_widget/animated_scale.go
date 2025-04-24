package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
)

/*
Example:
	AnimatedScale(
		&duit_attributes.AnimatedScaleAttributes{
			Scale: 2.0,
			ImplicitAnimatable: duit_animations.ImplicitAnimatable{
				Duration: 1000,
			},
		},
		"id",
		nil, //child
	)
*/
func AnimatedScale(attributes *duit_attributes.AnimatedScaleAttributes, id string, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.AnimatedScale, id, "", attributes, nil, true, 1, nil, child)
}