package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// VisibilityAttributes defines attributes for Visibility widget.
// See: https://api.flutter.dev/flutter/widgets/Visibility-class.html
type VisibilityAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
	*ThemeConsumer        `gen:"wrap"`
	Visible               duit_utils.Tristate[bool] `json:"visible,omitempty"`
	MaintainState         duit_utils.Tristate[bool] `json:"maintainState,omitempty"`
	MaintainAnimation     duit_utils.Tristate[bool] `json:"maintainAnimation,omitempty"`
	MaintainSize          duit_utils.Tristate[bool] `json:"maintainSize,omitempty"`
	MaintainSemantics     duit_utils.Tristate[bool] `json:"maintainSemantics,omitempty"`
	MaintainInteractivity duit_utils.Tristate[bool] `json:"maintainInteractivity,omitempty"`
	MaintainFocusability  duit_utils.Tristate[bool] `json:"maintainFocusability,omitempty"`
}

//bindgen:exclude
func (r *VisibilityAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if r.Visible == nil {
		return errors.New("visible property is required")
	}

	return nil
}

// NewVisibilityAttributes creates a new VisibilityAttributes instance.
func NewVisibilityAttributes() *VisibilityAttributes {
	return &VisibilityAttributes{}
}

// SetVisible sets the visible for the visibility widget.
func (r *VisibilityAttributes) SetVisible(visible bool) *VisibilityAttributes {
	r.Visible = duit_utils.BoolValue(visible)
	return r
}

// SetMaintainState sets the maintain state for the visibility widget.
func (r *VisibilityAttributes) SetMaintainState(maintainState bool) *VisibilityAttributes {
	r.MaintainState = duit_utils.BoolValue(maintainState)
	return r
}

// SetMaintainAnimation sets the maintain animation for the visibility widget.
func (r *VisibilityAttributes) SetMaintainAnimation(maintainAnimation bool) *VisibilityAttributes {
	r.MaintainAnimation = duit_utils.BoolValue(maintainAnimation)
	return r
}

// SetMaintainSize sets the maintain size for the visibility widget.
func (r *VisibilityAttributes) SetMaintainSize(maintainSize bool) *VisibilityAttributes {
	r.MaintainSize = duit_utils.BoolValue(maintainSize)
	return r
}

// SetMaintainSemantics sets the maintain semantics for the visibility widget.
func (r *VisibilityAttributes) SetMaintainSemantics(maintainSemantics bool) *VisibilityAttributes {
	r.MaintainSemantics = duit_utils.BoolValue(maintainSemantics)
	return r
}

// SetMaintainInteractivity sets the maintain interactivity for the visibility widget.
func (r *VisibilityAttributes) SetMaintainInteractivity(maintainInteractivity bool) *VisibilityAttributes {
	r.MaintainInteractivity = duit_utils.BoolValue(maintainInteractivity)
	return r
}

// SetMaintainFocusability sets the maintain focusability for the visibility widget.
func (r *VisibilityAttributes) SetMaintainFocusability(maintainFocusability bool) *VisibilityAttributes {
	r.MaintainFocusability = duit_utils.BoolValue(maintainFocusability)
	return r
}
