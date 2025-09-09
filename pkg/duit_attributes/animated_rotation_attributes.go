package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// AnimatedRotationAttributes represents attributes for AnimatedRotation widget.
// https://api.flutter.dev/flutter/widgets/AnimatedRotation-class.html
type AnimatedRotationAttributes struct {
	*ValueReferenceHolder          `gen:"wrap"`
	*ThemeConsumer                 `gen:"wrap"`
	*duit_props.ImplicitAnimatable `gen:"wrap"`
	Turns                          duit_utils.Tristate[float32] `json:"turns,omitempty"`
	Alignment                      duit_props.Alignment         `json:"alignment,omitempty"`
	FilterQuality                  duit_props.FilterQuality     `json:"filterQuality,omitempty"`
}

// NewAnimatedRotationAttributes creates a new instance of AnimatedRotationAttributes.
func NewAnimatedRotationAttributes() *AnimatedRotationAttributes {
	return &AnimatedRotationAttributes{}
}

// SetTurns sets the turns property and returns the AnimatedRotationAttributes instance for method chaining.
func (r *AnimatedRotationAttributes) SetTurns(turns float32) *AnimatedRotationAttributes {
	r.Turns = duit_utils.Float32Value(turns)
	return r
}

// SetAlignment sets the alignment property and returns the AnimatedRotationAttributes instance for method chaining.
func (r *AnimatedRotationAttributes) SetAlignment(alignment duit_props.Alignment) *AnimatedRotationAttributes {
	r.Alignment = alignment
	return r
}

// SetFilterQuality sets the filterQuality property and returns the AnimatedRotationAttributes instance for method chaining.
func (r *AnimatedRotationAttributes) SetFilterQuality(filterQuality duit_props.FilterQuality) *AnimatedRotationAttributes {
	r.FilterQuality = filterQuality
	return r
}

//bindgen:exclude
func (r *AnimatedRotationAttributes) Validate() error {
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

	if r.Turns == nil {
		return errors.New("turns property is required")
	}

	return nil
}
