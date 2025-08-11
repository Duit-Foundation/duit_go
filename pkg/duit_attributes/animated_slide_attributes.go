package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_flex"
)

type AnimatedSlideAttributes struct {
	ValueReferenceHolder
	duit_animations.ImplicitAnimatable
	ThemeConsumer
	Offset *duit_flex.Offset `json:"offset"`
} 