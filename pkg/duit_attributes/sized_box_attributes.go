package duit_attributes

import "github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"

// SizedBoxAttributes represents attributes for SizedBox widget.
// https://api.flutter.dev/flutter/widgets/SizedBox-class.html
type SizedBoxAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	Width                             float32 `json:"width,omitempty"`
	Height                            float32 `json:"height,omitempty"`
}

// NewSizedBoxAttributes creates a new instance of SizedBoxAttributes.
func NewSizedBoxAttributes() *SizedBoxAttributes {
	return &SizedBoxAttributes{}
}

// SetWidth sets the width property and returns the SizedBoxAttributes instance for method chaining.
func (r *SizedBoxAttributes) SetWidth(width float32) *SizedBoxAttributes {
	r.Width = width
	return r
}

// SetHeight sets the height property and returns the SizedBoxAttributes instance for method chaining.
func (r *SizedBoxAttributes) SetHeight(height float32) *SizedBoxAttributes {
	r.Height = height
	return r
}

//bindgen:exclude
func (r *SizedBoxAttributes) Validate() error {
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
