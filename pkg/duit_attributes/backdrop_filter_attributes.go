package duit_attributes

import (
	animations "github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_painting"
)

type BackdropFilterAttributes[T duit_painting.ImageFilter] struct {
	ValueReferenceHolder
	animations.AnimatedPropertyOwner
	ThemeConsumer
	Filter *T `json:"filter,omitempty"`
	BlendMode duit_painting.BlendMode `json:"blendMode,omitempty"`
}