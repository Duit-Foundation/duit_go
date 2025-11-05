package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// AnimatedPositionedAttributes represents attributes for AnimatedPositioned widget.
// https://api.flutter.dev/flutter/widgets/AnimatedPositioned-class.html
type AnimatedPositionedAttributes struct {
	*ValueReferenceHolder          `gen:"wrap"`
	*duit_props.ImplicitAnimatable `gen:"wrap"`
	*ThemeConsumer                 `gen:"wrap"`
	Left                           duit_utils.Tristate[float32] `json:"left,omitempty"`
	Top                            duit_utils.Tristate[float32] `json:"top,omitempty"`
	Right                          duit_utils.Tristate[float32] `json:"right,omitempty"`
	Bottom                         duit_utils.Tristate[float32] `json:"bottom,omitempty"`
}

// NewAnimatedPositionedAttributes creates a new instance of AnimatedPositionedAttributes.
func NewAnimatedPositionedAttributes() *AnimatedPositionedAttributes {
	return &AnimatedPositionedAttributes{}
}

// SetLeft sets the left property and returns the AnimatedPositionedAttributes instance for method chaining.
func (r *AnimatedPositionedAttributes) SetLeft(left float32) *AnimatedPositionedAttributes {
	r.Left = duit_utils.Float32Value(left)
	return r
}

// SetTop sets the top property and returns the AnimatedPositionedAttributes instance for method chaining.
func (r *AnimatedPositionedAttributes) SetTop(top float32) *AnimatedPositionedAttributes {
	r.Top = duit_utils.Float32Value(top)
	return r
}

// SetRight sets the right property and returns the AnimatedPositionedAttributes instance for method chaining.
func (r *AnimatedPositionedAttributes) SetRight(right float32) *AnimatedPositionedAttributes {
	r.Right = duit_utils.Float32Value(right)
	return r
}

// SetBottom sets the bottom property and returns the AnimatedPositionedAttributes instance for method chaining.
func (r *AnimatedPositionedAttributes) SetBottom(bottom float32) *AnimatedPositionedAttributes {
	r.Bottom = duit_utils.Float32Value(bottom)
	return r
}

//bindgen:exclude
func (r *AnimatedPositionedAttributes) Validate() error {
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

	return nil
}
