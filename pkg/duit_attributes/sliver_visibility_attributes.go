package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type SliverVisibilityAttributes struct {
	*ValueReferenceHolder
	*ThemeConsumer
	*SliverProps
	Visible               duit_utils.Tristate[bool]   `json:"visible,omitempty"`
	MaintainState         duit_utils.Tristate[bool]   `json:"maintainState,omitempty"`
	MaintainAnimation     duit_utils.Tristate[bool]   `json:"maintainAnimation,omitempty"`
	MaintainSize          duit_utils.Tristate[bool]   `json:"maintainSize,omitempty"`
	MaintainSemantics     duit_utils.Tristate[bool]   `json:"maintainSemantics,omitempty"`
	MaintainInteractivity duit_utils.Tristate[bool]   `json:"maintainInteractivity,omitempty"`
	ReplacementSliver     *duit_core.DuitElementModel `json:"replacementSliver,omitempty"`
}

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
