package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// AnimatedOpacityAttributes represents attributes for AnimatedOpacity widget.
// https://api.flutter.dev/flutter/widgets/AnimatedOpacity-class.html
type AnimatedOpacityAttributes struct {
	*ValueReferenceHolder          `gen:"wrap"`
	*duit_props.ImplicitAnimatable `gen:"wrap"`
	*ThemeConsumer                 `gen:"wrap"`
	Opacity                        duit_utils.Tristate[float32] `json:"opacity,omitempty"`
}

// NewAnimatedOpacityAttributes creates a new instance of AnimatedOpacityAttributes.
func NewAnimatedOpacityAttributes() *AnimatedOpacityAttributes {
	return &AnimatedOpacityAttributes{}
}

// SetOpacity sets the opacity property and returns the AnimatedOpacityAttributes instance for method chaining.
func (r *AnimatedOpacityAttributes) SetOpacity(opacity float32) *AnimatedOpacityAttributes {
	r.Opacity = duit_utils.Float32Value(opacity)
	return r
}

//bindgen:exclude
func (r *AnimatedOpacityAttributes) Validate() error {
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

	if r.Opacity != nil {
		if err := duit_utils.RangeValidator(*r.Opacity, 0, 1); err != nil {
			return err
		}
	}

	return nil
}
