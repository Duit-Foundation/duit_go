package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

/*
Example:
	SliverPadding(
		&duit_attributes.SliverPaddingAttributes[duit_edge_insets.EdgeInsetsAll]{
			Padding: 12.0,
		},
		"",
		false,
		nil, //child
	)
*/
func SliverPadding[T duit_edge_insets.EdgeInsets](attributes *duit_attributes.SliverPaddingAttributes[T], id string, controlled bool, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.SliverPadding, id, "", attributes, nil, controlled, 1, nil, child)
}
