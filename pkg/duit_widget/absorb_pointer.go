package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes"
)

func AbsorbPointer(attributes *duit_attributes.AbsorbPointerAttributes, id string, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.AbsorbPointer, id, "", attributes, nil, false, 1, nil, child)
}