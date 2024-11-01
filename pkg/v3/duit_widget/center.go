package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/pkg/duit_core"
)

func Center(attributes *duit_attributes.CenterAttributes, id string, controlled bool, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.Center, id, "", attributes, nil, controlled, 1, nil, child)
}
