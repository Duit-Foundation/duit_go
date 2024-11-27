package duit_attributes

import animations "github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_animations"

type RotatedBoxAttributes struct {
	ValueReferenceHolder
	animations.AnimatedPropertyOwner
	QuarterTurns int `json:"quarterTurns"`
}
