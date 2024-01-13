package duit_widget

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_core"
)

func ColumnUiElement(attributes *duit_attributes.FlexAttributes, id string, controlled bool) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.Column, id, "", attributes, nil, controlled, 2, nil)
}
