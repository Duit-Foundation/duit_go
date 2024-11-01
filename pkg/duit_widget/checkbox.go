package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_color"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
)

func CheckBox[T duit_color.Color](attributes *duit_attributes.CheckboxAttributes[T], id string, action *duit_core.Action) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.CheckBox, id, "", attributes, action, true, 0, nil)
}
