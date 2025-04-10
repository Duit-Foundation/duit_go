package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_alignment"
	animations "github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_clip"
)

type AnimatedSizeAttributes struct {
	ValueReferenceHolder
	animations.ImplicitAnimatable
	ThemeConsumer
	ClipBehavior    duit_clip.Clip           `json:"clipBehavior,omitempty"`
	Alignment       duit_alignment.Alignment `json:"alignment,omitempty"`
	ReverseDuration uint64                   `json:"reverseDuration,omitempty"`
}
