package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
)

func AnimatedPadding[TInsets duit_edge_insets.EdgeInsets](attributes *duit_attributes.AnimatedPaddingAttributes[TInsets], id string, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.AnimatedPadding, id, "", attributes, nil, true, 1, nil, child)
}