package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

/*
Example:
	SliverSafeArea(
		&duit_attributes.SliverSafeAreaAttributes[duit_props.EdgeInsetsAll]{
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
func SliverSafeArea[TInsets duit_props.EdgeInsets](attributes *duit_attributes.SliverSafeAreaAttributes[TInsets], id string, controlled bool, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.SliverSafeArea, id, "", attributes, nil, controlled, 1, nil, child)
} 