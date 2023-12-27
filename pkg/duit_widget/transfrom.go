package duit_widget

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_transform"
	"github.com/lesleysin/duit_go/pkg/duit_core"
)

func TransformUiElement[T duit_transform.Transform](attributes *duit_attributes.TransfromAttributes[T], id string, controlled bool, action *duit_core.Action) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.Transform, id, "", attributes, action, controlled, 1)
}
