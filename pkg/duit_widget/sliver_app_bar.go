package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_color"
	duit_decoration "github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
)

/*
Example:
	SliverAppBar(
		&duit_attributes.SliverAppBarAttributes[duit_color.Color, duit_edge_insets.EdgeInsetsAll, duit_decoration.RoundedRectangleBorder]{
			Title: Text(&duit_attributes.TextAttributes[duit_color.Color]{
				Data: "Sliver App Bar",
			}, "title"),
			Floating: true,
			Pinned: true,
			ExpandedHeight: 200.0,
		},
		"sliverAppBarId",
		true,
	)
*/
func SliverAppBar[TColor duit_color.Color, TInsets duit_edge_insets.EdgeInsets, TShape duit_decoration.ShapeBorder[TColor]](attributes *duit_attributes.SliverAppBarAttributes[TColor, TInsets, TShape], id string, controlled bool) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.SliverAppBar, id, "", attributes, nil, controlled, 0, nil)
} 