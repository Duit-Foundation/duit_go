package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_color"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

func CheckBox[TAction duit_action.Action, TColor duit_color.Color](attributes *duit_attributes.CheckboxAttributes[TColor], id string, action *TAction) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.CheckBox, id, "", attributes, action, true, 0, nil)
}
