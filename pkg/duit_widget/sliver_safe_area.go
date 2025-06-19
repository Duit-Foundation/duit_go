package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
)

/*
Example:
	SliverSafeArea(
		&duit_attributes.SliverSafeAreaAttributes[duit_edge_insets.EdgeInsetsAll]{
			Left:  true,
			Top:   true,
			Right: true,
			Bottom: true,
		},
		"",
		false,
		nil, //child
	)
*/
func SliverSafeArea[TInsets duit_edge_insets.EdgeInsets](attributes *duit_attributes.SliverSafeAreaAttributes[TInsets], id string, controlled bool, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.SliverSafeArea, id, "", attributes, nil, controlled, 1, nil, child)
} 