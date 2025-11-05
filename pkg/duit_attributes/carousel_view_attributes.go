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

// CarouselViewAttributes defines attributes for CarouselView widget.
// See: https://api.flutter.dev/flutter/material/CarouselView-class.html
type CarouselViewAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	// Constructor      CarouselViewConstructor                     `json:"constructor"`
	Padding          *duit_props.EdgeInsets       `json:"padding,omitempty"`
	BackgroundColor  *duit_props.Color            `json:"backgroundColor,omitempty"`
	Shape            *duit_props.ShapeBorder      `json:"shape,omitempty"`
	OverlayColor     *duit_props.WidgetStateColor `json:"overlayColor,omitempty"`
	Elevation        float32                      `json:"elevation,omitempty"`
	ShrinkExtent     float32                      `json:"shrinkExtent,omitempty"`
	ItemExtent       float32                      `json:"itemExtent,omitempty"`
	ScrollDirection  duit_props.Axis              `json:"scrollDirection,omitempty"`
	EnableSplash     duit_utils.Tristate[bool]    `json:"enableSplash,omitempty"`
	Reverse          duit_utils.Tristate[bool]    `json:"reverse,omitempty"`
	ItemSnapping     duit_utils.Tristate[bool]    `json:"itemSnapping,omitempty"`
	ConsumeMaxWeight duit_utils.Tristate[bool]    `json:"consumeMaxWeight,omitempty"`
	FlexWeights      []uint                       `json:"flexWeights,omitempty"`
}

//bindgen:exclude
func (r *CarouselViewAttributes) Validate() error {
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

// NewCarouselViewAttributes creates a new CarouselViewAttributes instance.
func NewCarouselViewAttributes() *CarouselViewAttributes {
	return &CarouselViewAttributes{}
}

// SetPadding sets the padding for the carousel view widget.
func (r *CarouselViewAttributes) SetPadding(padding *duit_props.EdgeInsets) *CarouselViewAttributes {
	r.Padding = padding
	return r
}

// SetBackgroundColor sets the background color for the carousel view widget.
func (r *CarouselViewAttributes) SetBackgroundColor(backgroundColor *duit_props.Color) *CarouselViewAttributes {
	r.BackgroundColor = backgroundColor
	return r
}

// SetShape sets the shape for the carousel view widget.
func (r *CarouselViewAttributes) SetShape(shape *duit_props.ShapeBorder) *CarouselViewAttributes {
	r.Shape = shape
	return r
}

// SetOverlayColor sets the overlay color for the carousel view widget.
func (r *CarouselViewAttributes) SetOverlayColor(overlayColor *duit_props.WidgetStateColor) *CarouselViewAttributes {
	r.OverlayColor = overlayColor
	return r
}

// SetElevation sets the elevation for the carousel view widget.
func (r *CarouselViewAttributes) SetElevation(elevation float32) *CarouselViewAttributes {
	r.Elevation = elevation
	return r
}

// SetShrinkExtent sets the shrink extent for the carousel view widget.
func (r *CarouselViewAttributes) SetShrinkExtent(shrinkExtent float32) *CarouselViewAttributes {
	r.ShrinkExtent = shrinkExtent
	return r
}

// SetItemExtent sets the item extent for the carousel view widget.
func (r *CarouselViewAttributes) SetItemExtent(itemExtent float32) *CarouselViewAttributes {
	r.ItemExtent = itemExtent
	return r
}

// SetScrollDirection sets the scroll direction for the carousel view widget.
func (r *CarouselViewAttributes) SetScrollDirection(scrollDirection duit_props.Axis) *CarouselViewAttributes {
	r.ScrollDirection = scrollDirection
	return r
}

// SetEnableSplash sets the enable splash for the carousel view widget.
func (r *CarouselViewAttributes) SetEnableSplash(enableSplash duit_utils.Tristate[bool]) *CarouselViewAttributes {
	r.EnableSplash = enableSplash
	return r
}

// SetReverse sets the reverse for the carousel view widget.
func (r *CarouselViewAttributes) SetReverse(reverse duit_utils.Tristate[bool]) *CarouselViewAttributes {
	r.Reverse = reverse
	return r
}

// SetItemSnapping sets the item snapping for the carousel view widget.
func (r *CarouselViewAttributes) SetItemSnapping(itemSnapping duit_utils.Tristate[bool]) *CarouselViewAttributes {
	r.ItemSnapping = itemSnapping
	return r
}

// SetConsumeMaxWeight sets the consume max weight for the carousel view widget.
func (r *CarouselViewAttributes) SetConsumeMaxWeight(consumeMaxWeight duit_utils.Tristate[bool]) *CarouselViewAttributes {
	r.ConsumeMaxWeight = consumeMaxWeight
	return r
}

// SetFlexWeights sets the flex weights for the carousel view widget.
func (r *CarouselViewAttributes) SetFlexWeights(flexWeights []uint) *CarouselViewAttributes {
	r.FlexWeights = flexWeights
	return r
}

// AddFlexWeight adds a single flex weight to the carousel view widget.
func (r *CarouselViewAttributes) AddFlexWeight(flexWeight uint) *CarouselViewAttributes {
	r.FlexWeights = append(r.FlexWeights, flexWeight)
	return r
}
