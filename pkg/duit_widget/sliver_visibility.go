package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

/*
Example:

	SliverVisibility(
		&duit_attributes.SliverPaddingAttributes[duit_props.EdgeInsetsAll]{
			Visible: duit_utils.BoolPtr(true),
		},
		"",
		false,
		nil, //child
		nil, //replacementSliver
	)
*/
func SliverVisibility(attributes *duit_attributes.SliverVisibilityAttributes, id string, controlled bool, sliver *duit_core.DuitElementModel, replacementSliver *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	children := []*duit_core.DuitElementModel{sliver, replacementSliver}

	checkAttributes(attributes)

	return new(duit_core.DuitElementModel).CreateElement(duit_core.SliverVisibility, id, "", attributes, nil, controlled, 2, nil, children...)
}
