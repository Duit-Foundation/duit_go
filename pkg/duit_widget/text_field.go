package duit_widget

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_color"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_edge_insets"
	"github.com/lesleysin/duit_go/pkg/duit_core"
)

func TextField[TInsets duit_edge_insets.EdgeInsets, TColor duit_color.Color](attributes *duit_attributes.TextFieldAttributes[TInsets, TColor], id string, action *duit_core.Action) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.TextField, id, "", attributes, action, true, 0, nil)
}
