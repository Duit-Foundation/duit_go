package duit_widget

import (
	"github.com/lesleysin/duit_go/pkg/duit_core"
)

func EmptyUiElement() *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.Empty, "", "", nil, nil, false, 0, nil)
}
