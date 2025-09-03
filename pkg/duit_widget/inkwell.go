package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

func InkWell[TColor duit_props.Color, TAction duit_action.Action, TShape duit_props.ShapeBorder[TColor]](attributes *duit_attributes.InkwellAttributes[TColor, TAction, TShape], id string, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.InkWell, id, "", attributes, nil, true, 1, nil, child)
}
