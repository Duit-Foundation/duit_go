package duit_widget

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_utility_styles"
	"github.com/lesleysin/duit_go/pkg/duit_core"
)

func ImageUiElement[T duit_utility_styles.Color](attributes *duit_attributes.ImageAttributes[T], id string, controlled bool, action *duit_core.Action) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.Image, id, "", attributes, action, controlled, 0)
}
