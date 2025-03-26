package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_color"
	duit_decoration "github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
)

func AppBar[TColor duit_color.Color, TInsets duit_edge_insets.EdgeInsets, TShape duit_decoration.ShapeBorder[TColor]](attributes *duit_attributes.AppBarAttributes[TColor, TInsets, TShape], id string) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.AppBar, id, "", attributes, nil, true, 0, nil)
}
