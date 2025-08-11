package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
)

func Offstage(attributes *duit_attributes.OffstageAttributes, id string, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.Offstage, id, "", attributes, nil, false, 1, nil, child)
}