package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
)

func RemoteSubtree(attributes *duit_attributes.RemoteSubtreeAttributes, id string) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.RotatedBox, id, "", attributes, nil, true, 1, nil, nil)
}
