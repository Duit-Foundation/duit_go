package duit_attributes

import "github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"

type SliverVisibilityAttributes struct {
	ValueReferenceHolder
	ThemeConsumer
	SliverProps
	Visible               *bool                       `json:"visible,omitempty"`
	MaintainState         *bool                       `json:"maintainState,omitempty"`
	MaintainAnimation     *bool                       `json:"maintainAnimation,omitempty"`
	MaintainSize          *bool                       `json:"maintainSize,omitempty"`
	MaintainSemantics     *bool                       `json:"maintainSemantics,omitempty"`
	MaintainInteractivity *bool                       `json:"maintainInteractivity,omitempty"`
	ReplacementSliver     *duit_core.DuitElementModel `json:"replacementSliver,omitempty"`
}
