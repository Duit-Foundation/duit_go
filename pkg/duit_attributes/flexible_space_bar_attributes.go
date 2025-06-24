package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_flex"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
)

type FlexibleSpaceBarAttributes[TInsets duit_edge_insets.EdgeInsets] struct {
	ValueReferenceHolder
	ThemeConsumer
	Title              *duit_core.DuitElementModel `json:"title,omitempty"`
	Background         *duit_core.DuitElementModel `json:"background,omitempty"`
	TitlePadding       TInsets                     `json:"titlePadding,omitempty"`
	CollapseMode       duit_flex.CollapseMode      `json:"collapseMode,omitempty"`
	StretchModes       []duit_flex.StretchMode     `json:"stretchModes,omitempty"`
	CenterTitle        *bool                       `json:"centerTitle,omitempty"`
	ExpandedTitleScale float32                     `json:"expandedTitleScale,omitempty"`
}
