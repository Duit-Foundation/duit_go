package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
)

/*
Example:
	FlexibleSpaceBar(
		&duit_attributes.FlexibleSpaceBarAttributes[duit_edge_insets.EdgeInsetsAll]{
			Title: Text(&duit_attributes.TextAttributes[duit_color.Color]{
				Data: "Flexible Title",
			}, "title"),
			CenterTitle: true,
			ExpandedTitleScale: 1.5,
			CollapseMode: duit_attributes.CollapseModeParallax,
			StretchModes: []duit_attributes.StretchMode{
				duit_attributes.StretchModeZoomBackground,
			},
		},
		"flexibleSpaceBarId",
		true,
	)
*/
func FlexibleSpaceBar[TInsets duit_edge_insets.EdgeInsets](attributes *duit_attributes.FlexibleSpaceBarAttributes[TInsets], id string, controlled bool) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.FlexibleSpaceBar, id, "", attributes, nil, controlled, 0, nil)
} 