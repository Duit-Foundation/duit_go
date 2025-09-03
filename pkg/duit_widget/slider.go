package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

func Slider[TAction duit_action.Action, TColor duit_props.Color](attributes *duit_attributes.SliderAttributes[TAction, TColor], id string) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.Slider, id, "", attributes, nil, true, 0, nil)
}
