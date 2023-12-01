package duit_widget

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_core"
)

func SizedBoxUiElement(attributes *duit_attributes.SizedBoxAttributes, id string, controlled bool, action *duit_core.Action) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.SizedBox, id, "", attributes, action, controlled)
}
