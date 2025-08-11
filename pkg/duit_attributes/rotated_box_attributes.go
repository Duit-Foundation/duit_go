package duit_attributes

import animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"

type RotatedBoxAttributes struct {
	ValueReferenceHolder
	ThemeConsumer
	animations.AnimatedPropertyOwner
	QuarterTurns int `json:"quarterTurns"`
}
