package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// AnimatedScaleAttributes represents attributes for AnimatedScale widget.
// https://api.flutter.dev/flutter/widgets/AnimatedScale-class.html
type AnimatedScaleAttributes struct {
	*ValueReferenceHolder          `gen:"wrap"`
	*ThemeConsumer                 `gen:"wrap"`
	*duit_props.ImplicitAnimatable `gen:"wrap"`
	Scale                          duit_utils.Tristate[float32] `json:"scale,omitempty"`
	Alignment                      duit_props.Alignment         `json:"alignment,omitempty"`
	FilterQuality                  duit_props.FilterQuality     `json:"filterQuality,omitempty"`
}

// NewAnimatedScaleAttributes creates a new instance of AnimatedScaleAttributes.
func NewAnimatedScaleAttributes() *AnimatedScaleAttributes {
	return &AnimatedScaleAttributes{}
}

// SetScale sets the scale property and returns the AnimatedScaleAttributes instance for method chaining.
func (r *AnimatedScaleAttributes) SetScale(scale float32) *AnimatedScaleAttributes {
	r.Scale = duit_utils.Float32Value(scale)
	return r
}

// SetAlignment sets the alignment property and returns the AnimatedScaleAttributes instance for method chaining.
func (r *AnimatedScaleAttributes) SetAlignment(alignment duit_props.Alignment) *AnimatedScaleAttributes {
	r.Alignment = alignment
	return r
}

// SetFilterQuality sets the filterQuality property and returns the AnimatedScaleAttributes instance for method chaining.
func (r *AnimatedScaleAttributes) SetFilterQuality(filterQuality duit_props.FilterQuality) *AnimatedScaleAttributes {
	r.FilterQuality = filterQuality
	return r
}

//bindgen:exclude
func (r *AnimatedScaleAttributes) Validate() error {
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

	var scaleV = r.Scale

	if scaleV == nil {
		return errors.New("scale property is required")
	} else {
		if *scaleV <= 0 {
			return errors.New("scale must be greater than 0 and positive")
		}
	}

	return nil
}
