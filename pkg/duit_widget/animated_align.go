package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

/*
Example:
	AnimatedAlign(
		&duit_attributes.AnimatedAlignAttributes{
			Alignment: duit_alignment.Center,
			ImplicitAnimatable: duit_animations.ImplicitAnimatable{
				Duration: 1000,
			},
		},
		"id",
		nil, //child
	)
*/
func AnimatedAlign(attributes *duit_attributes.AnimatedAlignAttributes, id string, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.AnimatedAlign, id, "", attributes, nil, true, 1, nil, child)
}