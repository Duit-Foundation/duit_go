package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	duit_decoration "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

func AppBar[TColor duit_props.Color, TInsets duit_props.EdgeInsets, TShape duit_decoration.ShapeBorder[TColor]](
	attributes *duit_attributes.AppBarAttributes[TColor, TInsets, TShape],
	id string,
	controlled bool,
	title, leading, flexibleSpace, bottom *duit_core.DuitElementModel,
	actions []*duit_core.DuitElementModel,
) *duit_core.DuitElementModel {
	children := []*duit_core.DuitElementModel{title, leading, flexibleSpace, bottom}

	if len(actions) > 0 {
		children = append(children, actions...)
	}

	return new(duit_core.DuitElementModel).CreateElement(duit_core.AppBar, id, "", attributes, nil, controlled, 2, nil, children...)
}
