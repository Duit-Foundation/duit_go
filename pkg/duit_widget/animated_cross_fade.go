package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

func AnimatedCrossFade(attributes *duit_attributes.AnimatedCrossFadeAttributes, id string, children []*duit_core.DuitElementModel) *duit_core.DuitElementModel {
	if len(children) != 2 {
		panic("AnimatedCrossFade must have 2 children")
	}
	return new(duit_core.DuitElementModel).CreateElement(duit_core.AnimatedCrossFade, id, "", attributes, nil, true, 2, nil, children...)
}
