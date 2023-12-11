package duit_widget

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_color"
	"github.com/lesleysin/duit_go/pkg/duit_core"
)

func DecoratedBoxUiElement[T duit_color.Color](attributes *duit_attributes.DecoratedBoxAttributes[T], id string, controlled bool, action *duit_core.Action) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.DecoratedBox, id, "", attributes, action, controlled, 1)
}
