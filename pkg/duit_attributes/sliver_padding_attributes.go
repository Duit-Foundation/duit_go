package duit_attributes

import (
	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_edge_insets"
)

type SliverPaddingAttributes[TInsets duit_edge_insets.EdgeInsets] struct {
	ValueReferenceHolder
	ThemeConsumer
	animations.AnimatedPropertyOwner
	SliverProps
	Padding TInsets `json:"padding,omitempty"`
}
