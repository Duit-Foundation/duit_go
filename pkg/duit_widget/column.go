package duit_widget

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_core"
)

func ColumnUiElement(attributes *duit_attributes.FlexAttributes, id string, controlled bool, action *duit_core.Action) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.Column, id, "", attributes, action, controlled, 2)
}
