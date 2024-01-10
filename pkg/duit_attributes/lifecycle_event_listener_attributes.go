package duit_attributes

import "github.com/lesleysin/duit_go/pkg/duit_core"

type LifecycleEventListenerAttributes struct {
	OnStateChanged *duit_core.Action `json:"onStateChanged,omitempty"`
	OnResumed      *duit_core.Action `json:"onResumed,omitempty"`
	OnInactive     *duit_core.Action `json:"onInactive,omitempty"`
	OnPaused       *duit_core.Action `json:"onPaused,omitempty"`
	OnDetached     *duit_core.Action `json:"onDetached,omitempty"`
	OnHidden       *duit_core.Action `json:"onHidden,omitempty"`
}
