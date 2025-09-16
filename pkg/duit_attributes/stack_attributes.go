package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// StackAttributes represents attributes for Stack widget.
// https://api.flutter.dev/flutter/widgets/Stack-class.html
type StackAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	TextDirection                     duit_props.TextDirection        `json:"textDirection,omitempty"`
	ClipBehavior                      duit_props.Clip                 `json:"clipBehavior,omitempty"`
	Alignment                         duit_props.AlignmentDirectional `json:"alignmentDirectional,omitempty"`
	Fit                               duit_props.StackFit             `json:"stackFit,omitempty"`
}

// NewStackAttributes creates a new instance of StackAttributes.
func NewStackAttributes() *StackAttributes {
	return &StackAttributes{}
}

// SetTextDirection sets the textDirection property and returns the StackAttributes instance for method chaining.
func (r *StackAttributes) SetTextDirection(textDirection duit_props.TextDirection) *StackAttributes {
	r.TextDirection = textDirection
	return r
}

// SetClipBehavior sets the clipBehavior property and returns the StackAttributes instance for method chaining.
func (r *StackAttributes) SetClipBehavior(clipBehavior duit_props.Clip) *StackAttributes {
	r.ClipBehavior = clipBehavior
	return r
}

// SetAlignment sets the alignment property and returns the StackAttributes instance for method chaining.
func (r *StackAttributes) SetAlignment(alignment duit_props.AlignmentDirectional) *StackAttributes {
	r.Alignment = alignment
	return r
}

// SetFit sets the fit property and returns the StackAttributes instance for method chaining.
func (r *StackAttributes) SetFit(fit duit_props.StackFit) *StackAttributes {
	r.Fit = fit
	return r
}

//bindgen:exclude
func (r *StackAttributes) Validate() error {
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
