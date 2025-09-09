package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// PaddingAttributes defines attributes for Padding widget.
// See: https://api.flutter.dev/flutter/widgets/Padding-class.html
type PaddingAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	Padding                           *duit_props.EdgeInsets `json:"padding,omitempty"`
}

//bindgen:exclude
func (r *PaddingAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if r.Padding == nil {
		return errors.New("padding property is required")
	} else {
		if err := r.Padding.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// NewPaddingAttributes creates a new PaddingAttributes instance.
func NewPaddingAttributes() *PaddingAttributes {
	return &PaddingAttributes{}
}

// SetPadding sets the padding for the padding widget.
func (r *PaddingAttributes) SetPadding(padding *duit_props.EdgeInsets) *PaddingAttributes {
	r.Padding = padding
	return r
}
