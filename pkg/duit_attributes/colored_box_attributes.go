package duit_attributes

import (
	animations "github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_color"
)

type ColoredBoxAttributes[T duit_color.Color] struct {
	ValueReferenceHolder
	animations.AnimatedPropertyOwner
	ThemeConsumer
	Color T `json:"color,omitempty"`
}
