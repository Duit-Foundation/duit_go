package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
)

/*
Example:
	SliverVisibility(
		&duit_attributes.SliverPaddingAttributes[duit_edge_insets.EdgeInsetsAll]{
			Visible: duit_utils.BoolPtr(true),
		},
		"",
		false,
		nil, //child
	)
*/
func SliverVisibility(attributes *duit_attributes.SliverVisibilityAttributes, id string, controlled bool, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.SliverVisibility, id, "", attributes, nil, controlled, 1, nil, child)
}
