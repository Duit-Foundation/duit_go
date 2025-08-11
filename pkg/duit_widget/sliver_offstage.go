package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

/*
Example:
	SliverOffstage(
		&duit_attributes.SliverOffstageAttributes{
			Offstage: duit_utils.BoolPtr(true),
		},
		"",
		false,
		nil, //child
	)
*/
func SliverOffstage(attributes *duit_attributes.SliverOffstageAttributes, id string, controlled bool, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.SliverOffstage, id, "", attributes, nil, controlled, 1, nil, child)
} 