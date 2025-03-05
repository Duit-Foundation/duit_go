package duit_attributes

import animations "github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_animations"

type IntrinsicWidthAttributes struct {
	ValueReferenceHolder
	animations.AnimatedPropertyOwner
	ThemeConsumer
	StepWidth  float32 `json:"stepWidth,omitempty"`
	StepHeight float32 `json:"stepHeight,omitempty"`
}
