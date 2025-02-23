package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
)

func IntrinsicWidth(attributes *duit_attributes.IntrinsicWidthAttributes, id string, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.IntrinsicWidth, id, "", attributes, nil, false, 1, nil, child)
}
