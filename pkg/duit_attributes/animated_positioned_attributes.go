package duit_attributes

import animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"

type AnimatedPositionedAttributes struct {
	ValueReferenceHolder
	animations.ImplicitAnimatable
	ThemeConsumer
	Left   float32 `json:"left,omitempty"`
	Top    float32 `json:"top,omitempty"`
	Right  float32 `json:"right,omitempty"`
	Bottom float32 `json:"bottom,omitempty"`
}