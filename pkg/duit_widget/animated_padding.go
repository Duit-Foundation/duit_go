package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

/*
Example:
	AnimatedPadding(
		&duit_attributes.AnimatedPaddingAttributes[duit_edge_insets.EdgeInsetsAll]{
			Padding: 12.0,
			ImplicitAnimatable: duit_animations.ImplicitAnimatable{
				Duration: 1000,
			},
		},
		"id",
		nil, //child
	)
*/
func AnimatedPadding[TInsets duit_edge_insets.EdgeInsets](attributes *duit_attributes.AnimatedPaddingAttributes[TInsets], id string, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.AnimatedPadding, id, "", attributes, nil, true, 1, nil, child)
}