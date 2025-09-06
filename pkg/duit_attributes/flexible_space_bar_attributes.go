package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type FlexibleSpaceBarAttributes struct {
	*ValueReferenceHolder
	*ThemeConsumer
	TitlePadding       *duit_props.EdgeInsets    `json:"titlePadding,omitempty"`
	CollapseMode       duit_props.CollapseMode   `json:"collapseMode,omitempty"`
	StretchModes       []duit_props.StretchMode  `json:"stretchModes,omitempty"`
	CenterTitle        duit_utils.Tristate[bool] `json:"centerTitle,omitempty"`
	ExpandedTitleScale float32                   `json:"expandedTitleScale,omitempty"`
}

func (r *FlexibleSpaceBarAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	return nil
}
