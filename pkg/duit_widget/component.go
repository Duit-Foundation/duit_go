package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

func componentUiElement(data interface{}, tag string, id string) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.Component, id, tag, nil, nil, true, 0, data)
}

func Component(data interface{}, tag string, id string) *duit_core.DuitElementModel {
	return componentUiElement(data, tag, id)
}
