package duit_attributes

import "github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"

// IntrinsicWidthAttributes defines attributes for IntrinsicWidth widget.
// See: https://api.flutter.dev/flutter/widgets/IntrinsicWidth-class.html
type IntrinsicWidthAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	StepWidth                         float32 `json:"stepWidth,omitempty"`
	StepHeight                        float32 `json:"stepHeight,omitempty"`
}

//bindgen:exclude
func (r *IntrinsicWidthAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}
	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	return nil
}

// NewIntrinsicWidthAttributes creates a new IntrinsicWidthAttributes instance.
func NewIntrinsicWidthAttributes() *IntrinsicWidthAttributes {
	return &IntrinsicWidthAttributes{}
}

// SetStepWidth sets the step width for the intrinsic width widget.
func (r *IntrinsicWidthAttributes) SetStepWidth(stepWidth float32) *IntrinsicWidthAttributes {
	r.StepWidth = stepWidth
	return r
}

// SetStepHeight sets the step height for the intrinsic width widget.
func (r *IntrinsicWidthAttributes) SetStepHeight(stepHeight float32) *IntrinsicWidthAttributes {
	r.StepHeight = stepHeight
	return r
}
