package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
)

type LifecycleEventListenerAttributes struct {
	*ValueReferenceHolder
	OnStateChanged any `json:"onStateChanged,omitempty"`
	OnResumed      any `json:"onResumed,omitempty"`
	OnInactive     any `json:"onInactive,omitempty"`
	OnPaused       any `json:"onPaused,omitempty"`
	OnDetached     any `json:"onDetached,omitempty"`
	OnHidden       any `json:"onHidden,omitempty"`
}

func (r *LifecycleEventListenerAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if r.OnStateChanged == nil && r.OnResumed == nil && r.OnInactive == nil && r.OnPaused == nil && r.OnDetached == nil && r.OnHidden == nil {
		return errors.New("at least one of the onStateChanged, onResumed, onInactive, onPaused, onDetached, onHidden properties must be non-nil")
	}

	if err := duit_action.CheckActionType(r.OnStateChanged); err != nil {
		return err
	}

	if err := duit_action.CheckActionType(r.OnResumed); err != nil {
		return err
	}

	if err := duit_action.CheckActionType(r.OnInactive); err != nil {
		return err
	}

	if err := duit_action.CheckActionType(r.OnPaused); err != nil {
		return err
	}

	if err := duit_action.CheckActionType(r.OnDetached); err != nil {
		return err
	}

	if err := duit_action.CheckActionType(r.OnHidden); err != nil {
		return err
	}

	return nil
}
