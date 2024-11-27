package duit_attributes

import animations "github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_animations"

type PositionedAttributes struct {
	ValueReferenceHolder
	animations.AnimatedPropertyOwner
	Left   float32 `json:"left,omitempty"`
	Top    float32 `json:"top,omitempty"`
	Right  float32 `json:"right,omitempty"`
	Bottom float32 `json:"bottom,omitempty"`
}
