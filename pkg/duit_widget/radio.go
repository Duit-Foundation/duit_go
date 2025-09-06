package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

func RadioGroupContext[TAction duit_action.Action, TValue duit_attributes.PrimitiveValue](attributes *duit_attributes.RadioGroupContextAttributes[TValue], id string, action *TAction, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.RadioGroupContext, id, "", attributes, action, true, 1, nil, child)
}

func Radio[TValue duit_attributes.PrimitiveValue](attributes *duit_attributes.RadioAttributes[TValue], id string, controlled bool) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.Radio, id, "", attributes, nil, controlled, 0, nil)
}
