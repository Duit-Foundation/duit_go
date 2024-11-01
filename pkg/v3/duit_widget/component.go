package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/pkg/duit_core"
)

func componentUiElement(data interface{}, tag string, id string) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.Component, id, tag, nil, nil, true, 0, data)
}

type componentDescription struct {
	Tag        string                      `json:"tag"`
	LayoutRoot *duit_core.DuitElementModel `json:"layoutRoot"`
}

func Component(data interface{}, tag string, id string) *duit_core.DuitElementModel {
	return componentUiElement(data, tag, id)
}

func ComponentDescription(tag string, layoutRoot *duit_core.DuitElementModel) *componentDescription {
	return &componentDescription{
		Tag:        tag,
		LayoutRoot: layoutRoot,
	}
}
