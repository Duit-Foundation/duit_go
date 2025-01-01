package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
)

func LifecycleEvent[TAction duit_core.Action](attributes *duit_attributes.LifecycleEventListenerAttributes[TAction], id string, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.LifecycleEventListener, id, "", attributes, nil, true, 1, nil, child)
}
