package duit_attributes

import animations "github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_animations"

type ConstrainedBoxAttributes struct {
	ValueReferenceHolder
	animations.AnimatedPropertyOwner
	MinWidth  float32 `json:"minWidth,omitempty"`
	MaxWidth  float32 `json:"maxWidth,omitempty"`
	MinHeight float32 `json:"minHeight,omitempty"`
	MaxHeight float32 `json:"maxHeight,omitempty"`
}
