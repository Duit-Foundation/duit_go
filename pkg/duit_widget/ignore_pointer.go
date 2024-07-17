package duit_widget

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_core"
)

func IgnorePointer(attributes *duit_attributes.IgnorePointerAttributes, id string, controlled bool) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.IgnorePointer, id, "", attributes, nil, controlled, 1, nil)
}
