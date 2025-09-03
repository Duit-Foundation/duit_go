package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

func AnimatedContainer[TInsets duit_props.EdgeInsets, TColor duit_props.Color](attributes *duit_attributes.AnimatedContainerAttributes[TInsets, TColor], id string, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.AnimatedContainer, id, "", attributes, nil, true, 1, nil, child)
}