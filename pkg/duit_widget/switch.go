package duit_widget

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_color"
	"github.com/lesleysin/duit_go/pkg/duit_core"
)

func SwitchUiElement[T duit_color.Color](attributes *duit_attributes.SwitchAttributes[T], id string, action *duit_core.Action) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.Switch, id, "", attributes, action, true, 0, nil)
}
