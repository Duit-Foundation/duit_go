package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// AnimatedSizeAttributes defines attributes for AnimatedSize widget.
// See: https://api.flutter.dev/flutter/widgets/AnimatedSize-class.html
type AnimatedSizeAttributes struct {
	*ValueReferenceHolder          `gen:"wrap"`
	*duit_props.ImplicitAnimatable `gen:"wrap"`
	*ThemeConsumer                 `gen:"wrap"`
	ClipBehavior                   duit_props.Clip      `json:"clipBehavior,omitempty"`
	Alignment                      duit_props.Alignment `json:"alignment,omitempty"`
	ReverseDuration                uint                 `json:"reverseDuration,omitempty"`
}

//bindgen:exclude
func (r *AnimatedSizeAttributes) Validate() error {
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

// NewAnimatedSizeAttributes creates a new AnimatedSizeAttributes instance.
func NewAnimatedSizeAttributes() *AnimatedSizeAttributes {
	return &AnimatedSizeAttributes{}
}

// SetClipBehavior sets the clip behavior for the animated size widget.
func (r *AnimatedSizeAttributes) SetClipBehavior(clipBehavior duit_props.Clip) *AnimatedSizeAttributes {
	r.ClipBehavior = clipBehavior
	return r
}

// SetAlignment sets the alignment for the animated size widget.
func (r *AnimatedSizeAttributes) SetAlignment(alignment duit_props.Alignment) *AnimatedSizeAttributes {
	r.Alignment = alignment
	return r
}

// SetReverseDuration sets the reverse duration for the animated size widget.
func (r *AnimatedSizeAttributes) SetReverseDuration(reverseDuration uint) *AnimatedSizeAttributes {
	r.ReverseDuration = reverseDuration
	return r
}
