package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

func CarouselView[TColor duit_props.Color, TInsets duit_props.EdgeInsets, TShape duit_props.ShapeBorder[TColor]](attributes *duit_attributes.CarouselViewAttributes[TColor, TInsets, TShape], id string, controlled bool, children []*duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.CarouselView, id, "", attributes, nil, controlled, 2, nil, children...)
}
