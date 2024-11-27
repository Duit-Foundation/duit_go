package duit_attributes

import animations "github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_animations"

type SizedBoxAttributes struct {
	ValueReferenceHolder
	animations.AnimatedPropertyOwner
	Width  float32 `json:"width,omitempty"`
	Height float32 `json:"height,omitempty"`
}
