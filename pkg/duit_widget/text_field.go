package duit_widget

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_edge_insets"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_utility_styles"
	"github.com/lesleysin/duit_go/pkg/duit_core"
)

func TextFieldUiElement[TInsets duit_edge_insets.EdgeInsets, TColor duit_utility_styles.Color](attributes *duit_attributes.TextFieldAttributes[TInsets, TColor], id string, controlled bool, action *duit_core.Action) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.TextField, id, "", attributes, action, controlled, 0)
}
