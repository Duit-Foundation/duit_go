package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	duit_decoration "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

func Card[TColor duit_props.Color, TInsets duit_edge_insets.EdgeInsets, TShape duit_decoration.ShapeBorder[TColor]] (attributes *duit_attributes.CardAttributes[TColor, TInsets, TShape], id string, controlled bool, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.Card, id, "", attributes, nil, controlled, 1, nil, child)
}
