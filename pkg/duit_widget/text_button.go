package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

/*
Example:
	TextButton(
		duit_attributes.NewTextButtonAttributes().
			SetStyle(style),
		"id",
		action,
		child,
	)
*/
func TextButton(attributes *duit_attributes.TextButtonAttributes, id string, action any, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	checkAttributes(attributes)
	return new(duit_core.DuitElementModel).CreateElement(duit_core.TextButton, id, "", attributes, action, true, 1, nil, child)
}
