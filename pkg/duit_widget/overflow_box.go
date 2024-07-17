package duit_widget

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_core"
)

func OverflowBox(attributes *duit_attributes.OverflowBoxAttributes, id string, controlled bool) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.OverflowBox, id, "", attributes, nil, controlled, 1, nil)
}
