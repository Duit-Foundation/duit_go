package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

func AbsorbPointer(attributes *duit_attributes.AbsorbPointerAttributes, id string, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	checkAttributes(attributes)
	return new(duit_core.DuitElementModel).CreateElement(duit_core.AbsorbPointer, id, "", attributes, nil, false, 1, nil, child)
}