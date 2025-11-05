package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// SliverVisibilityAttributes defines attributes for SliverVisibility widget.
// See: https://api.flutter.dev/flutter/widgets/SliverVisibility-class.html
type SliverVisibilityAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
	*ThemeConsumer        `gen:"wrap"`
	*SliverProps          `gen:"wrap"`
	Visible               duit_utils.Tristate[bool] `json:"visible,omitempty"`
	MaintainState         duit_utils.Tristate[bool] `json:"maintainState,omitempty"`
	MaintainAnimation     duit_utils.Tristate[bool] `json:"maintainAnimation,omitempty"`
	MaintainSize          duit_utils.Tristate[bool] `json:"maintainSize,omitempty"`
	MaintainSemantics     duit_utils.Tristate[bool] `json:"maintainSemantics,omitempty"`
	MaintainInteractivity duit_utils.Tristate[bool] `json:"maintainInteractivity,omitempty"`
}

//bindgen:exclude
func (r *SliverVisibilityAttributes) Validate() error {
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

// NewSliverVisibilityAttributes creates a new SliverVisibilityAttributes instance.
func NewSliverVisibilityAttributes() *SliverVisibilityAttributes {
	return &SliverVisibilityAttributes{}
}

// SetVisible sets the visible for the sliver visibility widget.
func (r *SliverVisibilityAttributes) SetVisible(visible duit_utils.Tristate[bool]) *SliverVisibilityAttributes {
	r.Visible = visible
	return r
}

// SetMaintainState sets the maintain state for the sliver visibility widget.
func (r *SliverVisibilityAttributes) SetMaintainState(maintainState duit_utils.Tristate[bool]) *SliverVisibilityAttributes {
	r.MaintainState = maintainState
	return r
}

// SetMaintainAnimation sets the maintain animation for the sliver visibility widget.
func (r *SliverVisibilityAttributes) SetMaintainAnimation(maintainAnimation duit_utils.Tristate[bool]) *SliverVisibilityAttributes {
	r.MaintainAnimation = maintainAnimation
	return r
}

// SetMaintainSize sets the maintain size for the sliver visibility widget.
func (r *SliverVisibilityAttributes) SetMaintainSize(maintainSize duit_utils.Tristate[bool]) *SliverVisibilityAttributes {
	r.MaintainSize = maintainSize
	return r
}

// SetMaintainSemantics sets the maintain semantics for the sliver visibility widget.
func (r *SliverVisibilityAttributes) SetMaintainSemantics(maintainSemantics duit_utils.Tristate[bool]) *SliverVisibilityAttributes {
	r.MaintainSemantics = maintainSemantics
	return r
}

// SetMaintainInteractivity sets the maintain interactivity for the sliver visibility widget.
func (r *SliverVisibilityAttributes) SetMaintainInteractivity(maintainInteractivity duit_utils.Tristate[bool]) *SliverVisibilityAttributes {
	r.MaintainInteractivity = maintainInteractivity
	return r
}
