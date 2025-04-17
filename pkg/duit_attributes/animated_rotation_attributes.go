package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_alignment"
	animations "github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_painting"
)

type AnimatedRotationAttributes struct {
	ValueReferenceHolder
	ThemeConsumer
	animations.ImplicitAnimatable
	Turns float32 `json:"turns"`
	Alignment duit_alignment.Alignment `json:"alignment,omitempty"`
	FilterQuality duit_painting.FilterQuality `json:"filterQuality,omitempty"`
}