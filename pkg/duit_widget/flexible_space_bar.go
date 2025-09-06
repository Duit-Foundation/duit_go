package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

/*
Example:

	FlexibleSpaceBar(
		&duit_attributes.FlexibleSpaceBarAttributes[duit_props.EdgeInsetsAll]{
			CenterTitle: true,
			ExpandedTitleScale: 1.5,
			CollapseMode: duit_attributes.CollapseModeParallax,
			StretchModes: []duit_attributes.StretchMode{
				duit_attributes.StretchModeZoomBackground,
			},
		},
		"flexibleSpaceBarId",
		true,
		Text(&duit_attributes.TextAttributes[duit_props.Color]{
			Data: "Flexible Title",
		}, "title"),
		Text(&duit_attributes.TextAttributes[duit_props.Color]{
			Data: "Flexible Background",
		}, "background"),
	)
*/
func FlexibleSpaceBar(attributes *duit_attributes.FlexibleSpaceBarAttributes, id string, controlled bool, title *duit_core.DuitElementModel, background *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	children := []*duit_core.DuitElementModel{title, background}

	checkAttributes(attributes)

	return new(duit_core.DuitElementModel).CreateElement(duit_core.FlexibleSpaceBar, id, "", attributes, nil, controlled, 2, nil, children...)
}
