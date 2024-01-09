package duit_widget

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_color"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_edge_insets"
	"github.com/lesleysin/duit_go/pkg/duit_core"
)

func ContainerUiElement[TInsets duit_edge_insets.EdgeInsets, TColor duit_color.Color](attributes *duit_attributes.ContainerAttributes[TInsets, TColor], id string, controlled bool) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.Container, id, "", attributes, nil, controlled, 1)
}
