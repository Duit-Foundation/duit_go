package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
)

func ElevatedButton[TAction duit_core.Action](attributes *duit_attributes.ElevatedButtonAttributes, id string, action *TAction, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.ElevatedButton, id, "", attributes, action, true, 1, nil, child)
}
