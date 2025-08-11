package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_edge_insets"
)

type SliverSafeAreaAttributes[TInsets duit_edge_insets.EdgeInsets] struct {
	ValueReferenceHolder
	ThemeConsumer
	SliverProps
	Left                      *bool   `json:"left,omitempty"`
	Top                       *bool   `json:"top,omitempty"`
	Right                     *bool   `json:"right,omitempty"`
	Bottom                    *bool   `json:"bottom,omitempty"`
	Minimum                   TInsets `json:"minimum,omitempty"`
	MaintainBottomViewPadding *bool   `json:"maintainBottomViewPadding,omitempty"`
}
