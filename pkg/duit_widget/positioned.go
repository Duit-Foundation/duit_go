package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
)

func Positioned(attributes *duit_attributes.PositionedAttributes, id string, controlled bool, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.Positioned, id, "", attributes, nil, controlled, 1, nil, child)
}
