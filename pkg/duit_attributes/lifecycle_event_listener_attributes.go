package duit_attributes

import "github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"

type LifecycleEventListenerAttributes[TAction duit_core.Action] struct {
	ValueReferenceHolder
	OnStateChanged *TAction `json:"onStateChanged,omitempty"`
	OnResumed      *TAction `json:"onResumed,omitempty"`
	OnInactive     *TAction `json:"onInactive,omitempty"`
	OnPaused       *TAction `json:"onPaused,omitempty"`
	OnDetached     *TAction `json:"onDetached,omitempty"`
	OnHidden       *TAction `json:"onHidden,omitempty"`
}
