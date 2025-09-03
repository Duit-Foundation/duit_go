package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

func Transform[T duit_props.Transform](attributes *duit_attributes.TransfromAttributes[T], id string, controlled bool, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.Transform, id, "", attributes, nil, controlled, 1, nil, child)
}
