package duit_attributes

import "github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"

// PositionedAttributes represents attributes for Positioned widget.
// https://api.flutter.dev/flutter/widgets/Positioned-class.html
type PositionedAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	Left                              float32 `json:"left,omitempty"`
	Top                               float32 `json:"top,omitempty"`
	Right                             float32 `json:"right,omitempty"`
	Bottom                            float32 `json:"bottom,omitempty"`
}

// NewPositionedAttributes creates a new instance of PositionedAttributes.
func NewPositionedAttributes() *PositionedAttributes {
	return &PositionedAttributes{}
}

// SetLeft sets the left property and returns the PositionedAttributes instance for method chaining.
func (r *PositionedAttributes) SetLeft(left float32) *PositionedAttributes {
	r.Left = left
	return r
}

// SetTop sets the top property and returns the PositionedAttributes instance for method chaining.
func (r *PositionedAttributes) SetTop(top float32) *PositionedAttributes {
	r.Top = top
	return r
}

// SetRight sets the right property and returns the PositionedAttributes instance for method chaining.
func (r *PositionedAttributes) SetRight(right float32) *PositionedAttributes {
	r.Right = right
	return r
}

// SetBottom sets the bottom property and returns the PositionedAttributes instance for method chaining.
func (r *PositionedAttributes) SetBottom(bottom float32) *PositionedAttributes {
	r.Bottom = bottom
	return r
}

//bindgen:exclude
func (r *PositionedAttributes) Validate() error {
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
