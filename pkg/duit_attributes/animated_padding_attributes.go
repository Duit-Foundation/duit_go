package duit_attributes

import (
	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_edge_insets"
)

type AnimatedPaddingAttributes[T duit_edge_insets.EdgeInsets] struct {
	ValueReferenceHolder
	animations.ImplicitAnimatable
	ThemeConsumer
	Padding T `json:"padding,omitempty"`
}