package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

/*
Example:

	Tooltip(
		duit_attributes.NewTooltipAttributes().SetMessage("Tap to open details"),
		"tooltip_1",
		false,
		nil,
	)
*/
func Tooltip(attributes *duit_attributes.TooltipAttributes, id string, controlled bool, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	checkAttributes(attributes)
	return new(duit_core.DuitElementModel).CreateElement(duit_core.Tooltip, id, "", attributes, nil, controlled, 1, nil, child)
}
