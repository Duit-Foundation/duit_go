package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

func CheckBox(attributes *duit_attributes.CheckboxAttributes, id string, action any) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.CheckBox, id, "", attributes, action, true, 0, nil)
}
