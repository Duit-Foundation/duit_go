package duit_widget

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_core"
)

func LifecycleEvent(attributes *duit_attributes.LifecycleEventListenerAttributes, id string) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.LifecycleEventListener, id, "", attributes, nil, true, 1, nil)
}
