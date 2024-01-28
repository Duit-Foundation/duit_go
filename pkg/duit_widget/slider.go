package duit_widget

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_color"
	"github.com/lesleysin/duit_go/pkg/duit_core"
)

func SliderUiElement[T duit_color.Color](attributes *duit_attributes.SliderAttributes[T], id string) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.Slider, id, "", attributes, nil, true, 0, nil)
}
