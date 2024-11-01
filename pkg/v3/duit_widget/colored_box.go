package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/pkg/duit_attributes/duit_color"
	"github.com/Duit-Foundation/duit_go/pkg/duit_core"
)

func ColoredBox[T duit_color.Color](attributes *duit_attributes.ColoredBoxAttributes[T], id string, controlled bool, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.ColoredBox, id, "", attributes, nil, controlled, 1, nil, child)
}
