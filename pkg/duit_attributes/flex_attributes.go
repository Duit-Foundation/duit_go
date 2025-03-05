package duit_attributes

import (
	animations "github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_clip"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_cross_axis"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_flex"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_main_axis"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_text_properties"
)

type FlexAttributes struct {
	ValueReferenceHolder
	animations.AnimatedPropertyOwner
	ThemeConsumer
	MainAxisAlignment  duit_main_axis.MainAxisAlignment   `json:"mainAxisAlignment,omitempty"`
	MainAxisSize       duit_main_axis.MainAxisSize        `json:"mainAxisSize,omitempty"`
	CrossAxisAlignment duit_cross_axis.CrossAxisAlignment `json:"crossAxisAlignment,omitempty"`
	TextDirection      duit_text_properties.TextDirection `json:"textDirection,omitempty"`
	VerticalDirection  duit_flex.VerticalDirection        `json:"verticalDirection,omitempty"`
	ClipBehavior       duit_clip.Clip                     `json:"clipBehavior,omitempty"`
}
