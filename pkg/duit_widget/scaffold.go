package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_color"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
)

func Scaffold[TColor duit_color.Color](attributes *duit_attributes.ScaffoldAttributes[TColor], id string) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.Scaffold, id, "", attributes, nil, true, 0, nil)
}
