package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// AnimatedSlideAttributes defines attributes for AnimatedSlide widget.
// See: https://api.flutter.dev/flutter/widgets/AnimatedSlide-class.html
type AnimatedSlideAttributes struct {
	*ValueReferenceHolder          `gen:"wrap"`
	*duit_props.ImplicitAnimatable `gen:"wrap"`
	*ThemeConsumer                 `gen:"wrap"`
	Offset                         *duit_props.Offset `json:"offset,omitempty"`
}

//bindgen:exclude
func (r *AnimatedSlideAttributes) Validate() error {
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

	if r.Offset == nil {
		return errors.New("offset property is required")
	}

	return nil
}

// NewAnimatedSlideAttributes creates a new AnimatedSlideAttributes instance.
func NewAnimatedSlideAttributes() *AnimatedSlideAttributes {
	return &AnimatedSlideAttributes{}
}

// SetOffset sets the offset for the animated slide widget.
func (r *AnimatedSlideAttributes) SetOffset(offset *duit_props.Offset) *AnimatedSlideAttributes {
	r.Offset = offset
	return r
}
