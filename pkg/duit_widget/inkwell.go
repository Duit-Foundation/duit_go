package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_color"
	duit_decoration "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

func InkWell[TColor duit_color.Color, TAction duit_action.Action, TShape duit_decoration.ShapeBorder[TColor]](attributes *duit_attributes.InkwellAttributes[TColor, TAction, TShape], id string, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.InkWell, id, "", attributes, nil, true, 1, nil, child)
}
