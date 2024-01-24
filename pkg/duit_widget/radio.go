package duit_widget

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_color"
	"github.com/lesleysin/duit_go/pkg/duit_core"
)

func RadioGroupContextUiElement[TValue duit_attributes.PrimitiveValue](attributes *duit_attributes.RadioGroupContextAttributes[TValue], id string, action *duit_core.Action) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.RadioGroupContext, id, "", attributes, action, true, 1, nil)
}

func RadioUiElement[TValue duit_attributes.PrimitiveValue, TColor duit_color.Color](attributes *duit_attributes.RadioAttributes[TValue, TColor], id string, controlled bool) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.Radio, id, "", attributes, nil, controlled, 0, nil)
}
