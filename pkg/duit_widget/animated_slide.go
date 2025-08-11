package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

/*
Example:
	AnimatedSlide(
		&duit_attributes.AnimatedSlideAttributes{
			Offset: &duit_flex.Offset{Dx: 0.5, Dy: 0.0},
			ImplicitAnimatable: duit_animations.ImplicitAnimatable{
				Duration: 1000,
			},
		},
		"id",
		nil, //child
	)
*/
func AnimatedSlide(attributes *duit_attributes.AnimatedSlideAttributes, id string, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.AnimatedSlide, id, "", attributes, nil, true, 1, nil, child)
} 