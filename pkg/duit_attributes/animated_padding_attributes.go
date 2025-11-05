package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// AnimatedPaddingAttributes represents attributes for AnimatedPadding widget.
// https://api.flutter.dev/flutter/widgets/AnimatedPadding-class.html
type AnimatedPaddingAttributes struct {
	*ValueReferenceHolder          `gen:"wrap"`
	*duit_props.ImplicitAnimatable `gen:"wrap"`
	*ThemeConsumer                 `gen:"wrap"`
	Padding                        *duit_props.EdgeInsets `json:"padding,omitempty"`
}

// NewAnimatedPaddingAttributes creates a new instance of AnimatedPaddingAttributes.
func NewAnimatedPaddingAttributes() *AnimatedPaddingAttributes {
	return &AnimatedPaddingAttributes{}
}

// SetPadding sets the padding property and returns the AnimatedPaddingAttributes instance for method chaining.
func (r *AnimatedPaddingAttributes) SetPadding(padding *duit_props.EdgeInsets) *AnimatedPaddingAttributes {
	r.Padding = padding
	return r
}

//bindgen:exclude
func (r *AnimatedPaddingAttributes) Validate() error {
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

	if r.Padding != nil {
		if err := r.Padding.Validate(); err != nil {
			return err
		}
	} else {
		return errors.New("padding property is required")
	}

	return nil
}
