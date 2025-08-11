package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_color"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

func TextField[TAction duit_core.Action, TInsets duit_edge_insets.EdgeInsets, TColor duit_color.Color](attributes *duit_attributes.TextFieldAttributes[TInsets, TColor], id string, action *TAction) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.TextField, id, "", attributes, action, true, 0, nil)
}
