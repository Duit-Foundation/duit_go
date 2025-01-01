package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_color"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
)

func Switch[TAction duit_core.Action, TColor duit_color.Color](attributes *duit_attributes.SwitchAttributes[TColor], id string, action *TAction) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.Switch, id, "", attributes, action, true, 0, nil)
}
