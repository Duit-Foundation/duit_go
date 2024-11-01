package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
)

func Subtree(id string, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.Subtree, id, "", nil, nil, true, 1, nil, child)
}
