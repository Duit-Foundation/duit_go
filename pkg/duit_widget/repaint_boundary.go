package duit_widget

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_core"
)

func RepaintBoundaryd(attributes *duit_attributes.RepaintBoundaryAttributes, id string, controlled bool, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.RepaintBoundary, id, "", attributes, nil, controlled, 1, nil, child)
}
