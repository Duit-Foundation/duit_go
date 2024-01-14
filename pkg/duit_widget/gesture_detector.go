package duit_widget

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_core"
)

func GestureDetectorUiElement(attributes *duit_attributes.GestureDetectorAttributes, id string) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.GestureDetector, id, "", attributes, nil, true, 1, nil)
}
