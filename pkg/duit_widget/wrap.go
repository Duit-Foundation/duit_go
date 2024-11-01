package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/pkg/duit_core"
)

func Wrap(attributes *duit_attributes.WrapAttributes, id string, controlled bool, action *duit_core.Action, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.Wrap, id, "", attributes, action, controlled, 2, nil, child)
}
