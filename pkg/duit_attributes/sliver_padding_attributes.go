package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// SliverPaddingAttributes defines attributes for SliverPadding widget.
// See: https://api.flutter.dev/flutter/widgets/SliverPadding-class.html
type SliverPaddingAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*SliverProps                      `gen:"wrap"`
	Padding                           *duit_props.EdgeInsets `json:"padding,omitempty"`
}

//bindgen:exclude
func (r *SliverPaddingAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if r.Padding == nil {
		return errors.New("padding property is required")
	}

	return nil
}

// NewSliverPaddingAttributes creates a new SliverPaddingAttributes instance.
func NewSliverPaddingAttributes() *SliverPaddingAttributes {
	return &SliverPaddingAttributes{}
}

// SetPadding sets the padding for the sliver padding widget.
func (r *SliverPaddingAttributes) SetPadding(padding *duit_props.EdgeInsets) *SliverPaddingAttributes {
	r.Padding = padding
	return r
}
