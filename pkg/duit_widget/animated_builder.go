package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

/*
Example usage:

	AnimatedBuilder(
		&duit_attributes.AnimatedBuilderAttributes{
			PersistentId: "1", //override it for use builder inside component layout
			Tweens: []any{
				duit_props.Tween(
					"height", //animated property name
					0.0, //begin
					100.0, //end
					3000, //duratuin
					duit_props.AnimationTriggerOnEnter, //trigger
					duit_props.AnimationMethodForward, //method
					false, //repeat
					"", //linear curve by default
				),
			},
		},
		"", //id
		nil, //child
	)
*/
func AnimatedBuilder(attributes *duit_attributes.AnimatedBuilderAttributes, id string, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.AnimatedBuilder, id, "", attributes, nil, true, 1, nil, child)
}
