package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

/*
Example:
	ClipOval(
		duit_attributes.NewClipOvalAttributes().
			SetClipBehavior(duit_props.ClipHardEdge),
		"id",
		false,
		nil,
	)
*/
func ClipOval(attributes *duit_attributes.ClipOvalAttributes, id string, controlled bool, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	checkAttributes(attributes)
	return new(duit_core.DuitElementModel).CreateElement(duit_core.ClipOval, id, "", attributes, nil, controlled, 1, nil, child)
}