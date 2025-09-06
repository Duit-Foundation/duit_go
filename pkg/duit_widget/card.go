package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

func Card[TInsets duit_props.EdgeInsets, TShape duit_props.ShapeBorder](attributes *duit_attributes.CardAttributes[TInsets, TShape], id string, controlled bool, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.Card, id, "", attributes, nil, controlled, 1, nil, child)
}
