package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_color"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
)

func AnimatedContainer[TInsets duit_edge_insets.EdgeInsets, TColor duit_color.Color](attributes *duit_attributes.AnimatedContainerAttributes[TInsets, TColor], id string, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.AnimatedContainer, id, "", attributes, nil, true, 1, nil, child)
}