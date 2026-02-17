package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

/*
Example:
	Icon(
		duit_attributes.NewIconAttributes().
			SetIcon(duit_props.NewIcon("Icons.add")).
			SetSize(24),
		"id",
		false,
	)
*/
func Icon(attributes *duit_attributes.IconAttributes, id string, controlled bool) *duit_core.DuitElementModel {
	checkAttributes(attributes)
	return new(duit_core.DuitElementModel).CreateElement(duit_core.Icon, id, "", attributes, nil, controlled, 0, nil)
}
