package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/pkg/duit_attributes/duit_color"
	"github.com/Duit-Foundation/duit_go/pkg/duit_core"
)

func Slider[T duit_color.Color](attributes *duit_attributes.SliderAttributes[T], id string) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.Slider, id, "", attributes, nil, true, 0, nil)
}
