package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
)

// LifecycleEventListenerAttributes defines attributes for LifecycleEventListener widget.
// See: https://api.flutter.dev/flutter/widgets/LifecycleEventListener-class.html
type LifecycleEventListenerAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
	OnStateChanged        any `json:"onStateChanged,omitempty"`
	OnResumed             any `json:"onResumed,omitempty"`
	OnInactive            any `json:"onInactive,omitempty"`
	OnPaused              any `json:"onPaused,omitempty"`
	OnDetached            any `json:"onDetached,omitempty"`
	OnHidden              any `json:"onHidden,omitempty"`
}

//bindgen:exclude
func (r *LifecycleEventListenerAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if r.OnStateChanged == nil && r.OnResumed == nil && r.OnInactive == nil && r.OnPaused == nil && r.OnDetached == nil && r.OnHidden == nil {
		return errors.New("at least one of the onStateChanged, onResumed, onInactive, onPaused, onDetached, onHidden properties must be non-nil")
	}

	callbacks := []any{
		r.OnStateChanged,
		r.OnResumed,
		r.OnInactive,
		r.OnPaused,
		r.OnDetached,
		r.OnHidden,
	}

	for _, callback := range callbacks {
		if err := duit_action.CheckActionType(callback); err != nil {
			return err
		}
	}

	return nil
}

// NewLifecycleEventListenerAttributes creates a new LifecycleEventListenerAttributes instance.
func NewLifecycleEventListenerAttributes() *LifecycleEventListenerAttributes {
	return &LifecycleEventListenerAttributes{}
}

// SetOnStateChanged sets the on state changed for the lifecycle event listener widget.
func (r *LifecycleEventListenerAttributes) SetOnStateChanged(onStateChanged any) *LifecycleEventListenerAttributes {
	r.OnStateChanged = onStateChanged
	return r
}

// SetOnResumed sets the on resumed for the lifecycle event listener widget.
func (r *LifecycleEventListenerAttributes) SetOnResumed(onResumed any) *LifecycleEventListenerAttributes {
	r.OnResumed = onResumed
	return r
}

// SetOnInactive sets the on inactive for the lifecycle event listener widget.
func (r *LifecycleEventListenerAttributes) SetOnInactive(onInactive any) *LifecycleEventListenerAttributes {
	r.OnInactive = onInactive
	return r
}

// SetOnPaused sets the on paused for the lifecycle event listener widget.
func (r *LifecycleEventListenerAttributes) SetOnPaused(onPaused any) *LifecycleEventListenerAttributes {
	r.OnPaused = onPaused
	return r
}

// SetOnDetached sets the on detached for the lifecycle event listener widget.
func (r *LifecycleEventListenerAttributes) SetOnDetached(onDetached any) *LifecycleEventListenerAttributes {
	r.OnDetached = onDetached
	return r
}

// SetOnHidden sets the on hidden for the lifecycle event listener widget.
func (r *LifecycleEventListenerAttributes) SetOnHidden(onHidden any) *LifecycleEventListenerAttributes {
	r.OnHidden = onHidden
	return r
}
