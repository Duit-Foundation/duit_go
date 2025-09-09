package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// AnimatedPhysicalModelAttributes represents attributes for AnimatedPhysicalModel widget.
// https://api.flutter.dev/flutter/widgets/AnimatedPhysicalModel-class.html
type AnimatedPhysicalModelAttributes struct {
	*ValueReferenceHolder          `gen:"wrap"`
	*duit_props.ImplicitAnimatable `gen:"wrap"`
	*ThemeConsumer                 `gen:"wrap"`
	Elevation                      duit_utils.Tristate[float32] `json:"elevation"`
	Color                          *duit_props.Color            `json:"color,omitempty"`
	ShadowColor                    *duit_props.Color            `json:"shadowColor,omitempty"`
	ClipBehavior                   duit_props.Clip              `json:"clipBehavior,omitempty"`
	BorderRadius                   *duit_props.BorderRadius     `json:"borderRadius,omitempty"`
	Shape                          duit_props.BoxShape          `json:"shape,omitempty"`
	AnimateColor                   duit_utils.Tristate[bool]    `json:"animateColor,omitempty"`
	AnimateShadowColor             duit_utils.Tristate[bool]    `json:"animateShadowColor,omitempty"`
}

// NewAnimatedPhysicalModelAttributes creates a new instance of AnimatedPhysicalModelAttributes.
func NewAnimatedPhysicalModelAttributes() *AnimatedPhysicalModelAttributes {
	return &AnimatedPhysicalModelAttributes{}
}

// SetElevation sets the elevation property and returns the AnimatedPhysicalModelAttributes instance for method chaining.
func (r *AnimatedPhysicalModelAttributes) SetElevation(elevation float32) *AnimatedPhysicalModelAttributes {
	r.Elevation = duit_utils.Float32Value(elevation)
	return r
}

// SetColor sets the color property and returns the AnimatedPhysicalModelAttributes instance for method chaining.
func (r *AnimatedPhysicalModelAttributes) SetColor(color *duit_props.Color) *AnimatedPhysicalModelAttributes {
	r.Color = color
	return r
}

// SetShadowColor sets the shadowColor property and returns the AnimatedPhysicalModelAttributes instance for method chaining.
func (r *AnimatedPhysicalModelAttributes) SetShadowColor(shadowColor *duit_props.Color) *AnimatedPhysicalModelAttributes {
	r.ShadowColor = shadowColor
	return r
}

// SetClipBehavior sets the clipBehavior property and returns the AnimatedPhysicalModelAttributes instance for method chaining.
func (r *AnimatedPhysicalModelAttributes) SetClipBehavior(clipBehavior duit_props.Clip) *AnimatedPhysicalModelAttributes {
	r.ClipBehavior = clipBehavior
	return r
}

// SetBorderRadius sets the borderRadius property and returns the AnimatedPhysicalModelAttributes instance for method chaining.
func (r *AnimatedPhysicalModelAttributes) SetBorderRadius(borderRadius *duit_props.BorderRadius) *AnimatedPhysicalModelAttributes {
	r.BorderRadius = borderRadius
	return r
}

// SetShape sets the shape property and returns the AnimatedPhysicalModelAttributes instance for method chaining.
func (r *AnimatedPhysicalModelAttributes) SetShape(shape duit_props.BoxShape) *AnimatedPhysicalModelAttributes {
	r.Shape = shape
	return r
}

// SetAnimateColor sets the animateColor property and returns the AnimatedPhysicalModelAttributes instance for method chaining.
func (r *AnimatedPhysicalModelAttributes) SetAnimateColor(animateColor bool) *AnimatedPhysicalModelAttributes {
	r.AnimateColor = duit_utils.BoolValue(animateColor)
	return r
}

// SetAnimateShadowColor sets the animateShadowColor property and returns the AnimatedPhysicalModelAttributes instance for method chaining.
func (r *AnimatedPhysicalModelAttributes) SetAnimateShadowColor(animateShadowColor bool) *AnimatedPhysicalModelAttributes {
	r.AnimateShadowColor = duit_utils.BoolValue(animateShadowColor)
	return r
}

//bindgen:exclude
func (r *AnimatedPhysicalModelAttributes) Validate() error {
	if r.ImplicitAnimatable != nil {
		if err := r.ImplicitAnimatable.Validate(); err != nil {
			return err
		}
	} else {
		return errors.New("implicitAnimatable property is required on implicit animated widgets")
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if r.Elevation != nil {
		if *r.Elevation < 0 {
			return errors.New("elevation must be greater than 0")
		}
	} else {
		return errors.New("elevation property is required")
	}

	if r.Color != nil {
		if err := r.Color.Validate(); err != nil {
			return err
		}
	} else {
		return errors.New("color property is required")
	}

	if r.ShadowColor != nil {
		if err := r.ShadowColor.Validate(); err != nil {
			return err
		}
	} else {
		return errors.New("shadowColor property is required")
	}

	if r.BorderRadius != nil {
		if err := r.BorderRadius.Validate(); err != nil {
			return err
		}
	}

	return nil
}
