package duit_attributes

import (
	animations "github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_edge_insets"
)

type PaddingAttributes[T duit_edge_insets.EdgeInsets] struct {
	ValueReferenceHolder
	animations.AnimatedPropertyOwner
	Padding T `json:"padding,omitempty"`
}
