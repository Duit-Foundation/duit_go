package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_clip"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_color"
	duit_decoration "github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_decorations"
)

type PhysicalModelAttributes[TColor duit_color.Color] struct {
	ValueReferenceHolder
	duit_animations.AnimatedPropertyOwner
	ThemeConsumer
	Elevation    float32                       `json:"elevation"`
	Color        TColor                        `json:"color,omitempty"`
	ShadowColor  TColor                        `json:"shadowColor,omitempty"`
	ClipBehavior duit_clip.Clip                `json:"clipBehavior,omitempty"`
	BorderRadius *duit_decoration.BorderRadius `json:"borderRadius,omitempty"`
	Shape        duit_decoration.BoxShape      `json:"shape,omitempty"`
}
