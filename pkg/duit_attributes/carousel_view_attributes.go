package duit_attributes

import (
	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_color"
	decorations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_flex"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_material"
)

type CarouselViewConstructor uint

const (
	CarouselViewCommon CarouselViewConstructor = iota
	CarouselViewWeighted
)

type CarouselViewAttributes[TColor duit_color.Color, TInsets duit_edge_insets.EdgeInsets, TShape decorations.ShapeBorder[TColor]] struct {
	ValueReferenceHolder
	animations.AnimatedPropertyOwner
	ThemeConsumer
	Constructor      CarouselViewConstructor                     `json:"constructor"`
	Padding          TInsets                                     `json:"padding,omitempty"`
	BackgroundColor  TColor                                      `json:"backgroundColor,omitempty"`
	Shape            *TShape                                     `json:"shape,omitempty"`
	OverlayColor     duit_material.MaterialStateProperty[TColor] `json:"overlayColor,omitempty"`
	Elevation        float32                                     `json:"elevation,omitempty"`
	ShrinkExtent     float32                                     `json:"shrinkExtent,omitempty"`
	ItemExtent       float32                                     `json:"itemExtent,omitempty"`
	ScrollDirection  duit_flex.Axis                              `json:"scrollDirection,omitempty"`
	EnableSplash     *bool                                       `json:"enableSplash,omitempty"`
	Reverse          *bool                                       `json:"reverse,omitempty"`
	ItemSnapping     *bool                                       `json:"itemSnapping,omitempty"`
	ConsumeMaxWeight *bool                                       `json:"consumeMaxWeight,omitempty"`
	FlexWeights      []uint                                      `json:"flexWeights,omitempty"`
}
