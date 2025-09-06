package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// type CarouselViewConstructor uint

// const (
// 	CarouselViewCommon CarouselViewConstructor = iota
// 	CarouselViewWeighted
// )

type CarouselViewAttributes[TShape duit_props.ShapeBorder] struct {
	*ValueReferenceHolder
	*duit_props.AnimatedPropertyOwner
	*ThemeConsumer
	// Constructor      CarouselViewConstructor                     `json:"constructor"`
	Padding          *duit_props.EdgeInsets                               `json:"padding,omitempty"`
	BackgroundColor  *duit_props.Color                                    `json:"backgroundColor,omitempty"`
	Shape            *TShape                                              `json:"shape,omitempty"`
	OverlayColor     *duit_props.MaterialStateProperty[*duit_props.Color] `json:"overlayColor,omitempty"`
	Elevation        float32                                              `json:"elevation,omitempty"`
	ShrinkExtent     float32                                              `json:"shrinkExtent,omitempty"`
	ItemExtent       float32                                              `json:"itemExtent,omitempty"`
	ScrollDirection  duit_props.Axis                                      `json:"scrollDirection,omitempty"`
	EnableSplash     duit_utils.Tristate[bool]                            `json:"enableSplash,omitempty"`
	Reverse          duit_utils.Tristate[bool]                            `json:"reverse,omitempty"`
	ItemSnapping     duit_utils.Tristate[bool]                            `json:"itemSnapping,omitempty"`
	ConsumeMaxWeight duit_utils.Tristate[bool]                            `json:"consumeMaxWeight,omitempty"`
	FlexWeights      []uint                                               `json:"flexWeights,omitempty"`
}

func (r *CarouselViewAttributes[TShape]) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	return nil
}
