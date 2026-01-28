package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

/*
Example:

	Visibility(
		duit_attributes.NewVisibilityAttributes().SetVisible(true),
		"",
		false,
		nil, //child
		nil, //replacement
	)
*/
func Visibility(attributes *duit_attributes.VisibilityAttributes, id string, controlled bool, child *duit_core.DuitElementModel, replacement *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	children := []*duit_core.DuitElementModel{child, replacement}

	checkAttributes(attributes)

	return new(duit_core.DuitElementModel).CreateElement(duit_core.Visibility, id, "", attributes, nil, controlled, 2, nil, children...)
}
