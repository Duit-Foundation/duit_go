package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

/*
Example:
	SliverPadding(
		&duit_attributes.SliverPaddingAttributes{
			Padding: &duit_props.EdgeInsetsAll{
				Top:    12.0,
				Right:  12.0,
				Bottom: 12.0,
				Left:   12.0,
			},
		},
		"",
		false,
		nil, //child
	)
*/
func SliverPadding(attributes *duit_attributes.SliverPaddingAttributes, id string, controlled bool, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	checkAttributes(attributes)
	return new(duit_core.DuitElementModel).CreateElement(duit_core.SliverPadding, id, "", attributes, nil, controlled, 1, nil, child)
}
