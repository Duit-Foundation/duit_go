package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

func CarouselView[TShape duit_props.ShapeBorder](attributes *duit_attributes.CarouselViewAttributes[TShape], id string, controlled bool, children []*duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.CarouselView, id, "", attributes, nil, controlled, 2, nil, children...)
}
