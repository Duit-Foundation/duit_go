package duit_widget

import (
	"github.com/lesleysin/duit_go/pkg/duit_core"
)

func SubtreeUiElement(id string) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.Subtree, id, "", nil, nil, true, 1, nil)
}
