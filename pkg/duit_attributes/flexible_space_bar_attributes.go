package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_flex"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)
type FlexibleSpaceBarAttributes[TInsets duit_edge_insets.EdgeInsets] struct {
	*ValueReferenceHolder
	*ThemeConsumer
	TitlePadding       TInsets                     `json:"titlePadding,omitempty"`
	CollapseMode       duit_flex.CollapseMode      `json:"collapseMode,omitempty"`
	StretchModes       []duit_flex.StretchMode     `json:"stretchModes,omitempty"`
	CenterTitle        duit_utils.Tristate[bool]   `json:"centerTitle,omitempty"`
	ExpandedTitleScale float32                     `json:"expandedTitleScale,omitempty"`
}

func (r *FlexibleSpaceBarAttributes[TInsets]) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	return nil
}
