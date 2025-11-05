package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// PhysicalModelAttributes defines attributes for PhysicalModel widget.
// See: https://api.flutter.dev/flutter/widgets/PhysicalModel-class.html
type PhysicalModelAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	Elevation                         duit_utils.Tristate[float32] `json:"elevation,omitempty"`
	Color                             *duit_props.Color            `json:"color,omitempty"`
	ShadowColor                       *duit_props.Color            `json:"shadowColor,omitempty"`
	ClipBehavior                      duit_props.Clip              `json:"clipBehavior,omitempty"`
	BorderRadius                      *duit_props.BorderRadius     `json:"borderRadius,omitempty"`
	Shape                             duit_props.BoxShape          `json:"boxShape,omitempty"`
}

//bindgen:exclude
func (r *PhysicalModelAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if r.Elevation == nil {
		return errors.New("elevation property is required")
	}

	return nil
}

// NewPhysicalModelAttributes creates a new PhysicalModelAttributes instance.
func NewPhysicalModelAttributes() *PhysicalModelAttributes {
	return &PhysicalModelAttributes{}
}

// SetElevation sets the elevation for the physical model widget.
func (r *PhysicalModelAttributes) SetElevation(elevation duit_utils.Tristate[float32]) *PhysicalModelAttributes {
	r.Elevation = elevation
	return r
}

// SetColor sets the color for the physical model widget.
func (r *PhysicalModelAttributes) SetColor(color *duit_props.Color) *PhysicalModelAttributes {
	r.Color = color
	return r
}

// SetShadowColor sets the shadow color for the physical model widget.
func (r *PhysicalModelAttributes) SetShadowColor(shadowColor *duit_props.Color) *PhysicalModelAttributes {
	r.ShadowColor = shadowColor
	return r
}

// SetClipBehavior sets the clip behavior for the physical model widget.
func (r *PhysicalModelAttributes) SetClipBehavior(clipBehavior duit_props.Clip) *PhysicalModelAttributes {
	r.ClipBehavior = clipBehavior
	return r
}

// SetBorderRadius sets the border radius for the physical model widget.
func (r *PhysicalModelAttributes) SetBorderRadius(borderRadius *duit_props.BorderRadius) *PhysicalModelAttributes {
	r.BorderRadius = borderRadius
	return r
}

// SetShape sets the shape for the physical model widget.
func (r *PhysicalModelAttributes) SetShape(shape duit_props.BoxShape) *PhysicalModelAttributes {
	r.Shape = shape
	return r
}
