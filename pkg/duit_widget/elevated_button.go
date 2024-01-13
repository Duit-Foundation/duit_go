package duit_widget

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_core"
)

func ElevatedButtonUiElement(attributes *duit_attributes.ElevatedButtonAttributes, id string, action *duit_core.Action) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.ElevatedButton, id, "", attributes, action, true, 1, nil)
}
