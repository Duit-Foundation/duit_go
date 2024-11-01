package duit_attributes

import "github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"

type LifecycleEventListenerAttributes struct {
	ValueReferenceHolder
	OnStateChanged *duit_core.Action `json:"onStateChanged,omitempty"`
	OnResumed      *duit_core.Action `json:"onResumed,omitempty"`
	OnInactive     *duit_core.Action `json:"onInactive,omitempty"`
	OnPaused       *duit_core.Action `json:"onPaused,omitempty"`
	OnDetached     *duit_core.Action `json:"onDetached,omitempty"`
	OnHidden       *duit_core.Action `json:"onHidden,omitempty"`
}
