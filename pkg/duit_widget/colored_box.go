package duit_widget

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_color"
	"github.com/lesleysin/duit_go/pkg/duit_core"
)

func ColoredBoxUiElement[T duit_color.Color](attributes *duit_attributes.ColoredBoxAttributes[T], id string, controlled bool) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.ColoredBox, id, "", attributes, nil, controlled, 1)
}
