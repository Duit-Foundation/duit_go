package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

/*
Example:

	SliverAppBar(
		&duit_attributes.SliverAppBarAttributes[duit_props.Color, duit_props.EdgeInsetsAll, duit_props.RoundedRectangleBorder]{
			Floating: true,
			Pinned: true,
			ExpandedHeight: 200.0,
		},
		"sliverAppBarId",
		true,
		Text(&duit_attributes.TextAttributes[duit_props.Color]{
				Data: "Sliver App Bar",
			}, "title"),
		nil, //leading
		nil, //flexibleSpace
		nil, //bottom
		[]*duit_core.DuitElementModel{
			Text(&duit_attributes.TextAttributes[duit_props.Color]{
				Data: "Action 1",
			}, "action1"),
		},
	)
*/
func SliverAppBar[TInsets duit_props.EdgeInsets, TShape duit_props.ShapeBorder](attributes *duit_attributes.SliverAppBarAttributes[TInsets, TShape], id string, controlled bool, title, leading, flexibleSpace, bottom *duit_core.DuitElementModel, actions []*duit_core.DuitElementModel) *duit_core.DuitElementModel {
	children := []*duit_core.DuitElementModel{title, leading, flexibleSpace, bottom}

	if len(actions) > 0 {
		children = append(children, actions...)
	}

	checkAttributes(attributes)

	return new(duit_core.DuitElementModel).CreateElement(duit_core.SliverAppBar, id, "", attributes, nil, controlled, 2, nil, children...)
}
