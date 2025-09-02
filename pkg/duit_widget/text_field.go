package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

func TextField[TAction duit_action.Action, TInsets duit_props.EdgeInsets, TColor duit_props.Color](attributes *duit_attributes.TextFieldAttributes[TInsets, TColor], id string, action *TAction) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.TextField, id, "", attributes, action, true, 0, nil)
}
