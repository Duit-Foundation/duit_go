package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

func RadioGroupContext(attributes *duit_attributes.RadioGroupContextAttributes, id string, action any, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	checkAttributes(attributes)
	return new(duit_core.DuitElementModel).CreateElement(duit_core.RadioGroupContext, id, "", attributes, action, true, 1, nil, child)
}

func Radio(attributes *duit_attributes.RadioAttributes, id string, controlled bool) *duit_core.DuitElementModel {
	checkAttributes(attributes)
	return new(duit_core.DuitElementModel).CreateElement(duit_core.Radio, id, "", attributes, nil, controlled, 0, nil)
}
