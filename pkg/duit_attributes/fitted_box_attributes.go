package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_alignment"
	animations "github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_clip"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_flex"
)

type FittedBoxAttributes struct {
	ValueReferenceHolder
	animations.AnimatedPropertyOwner
	Fit          duit_flex.BoxFit         `json:"fit,omitempty"`
	ClipBehavior duit_clip.Clip           `json:"clipBehavior,omitempty"`
	Alignment    duit_alignment.Alignment `json:"alignment,omitempty"`
}
