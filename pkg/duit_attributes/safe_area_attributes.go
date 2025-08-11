package duit_attributes

import (
	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_edge_insets"
)

type SafeAreaAttributes[TInsets duit_edge_insets.EdgeInsets] struct {
	ValueReferenceHolder
	animations.AnimatedPropertyOwner
	ThemeConsumer
	Left                      *bool    `json:"left,omitempty"`
	Top                       *bool    `json:"top,omitempty"`
	Right                     *bool    `json:"right,omitempty"`
	Bottom                    *bool    `json:"bottom,omitempty"`
	Minimum                   TInsets `json:"minimum,omitempty"`
	MaintainBottomViewPadding *bool    `json:"maintainBottomViewPadding,omitempty"`
}
