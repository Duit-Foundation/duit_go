package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_alignment"
	animations "github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_clip"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_flex"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_text_properties"
)

type StackAttributes struct {
	ValueReferenceHolder
	animations.AnimatedPropertyOwner
	ThemeConsumer
	TextDirection duit_text_properties.TextDirection `json:"textDirection,omitempty"`
	ClipBehavior  duit_clip.Clip                     `json:"clipBehavior,omitempty"`
	Alignment     duit_alignment.Alignment           `json:"alignment,omitempty"`
	Fit           duit_flex.StackFit                 `json:"fit,omitempty"`
}
