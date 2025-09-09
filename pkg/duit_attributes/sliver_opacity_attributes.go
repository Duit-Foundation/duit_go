package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// SliverOpacityAttributes defines attributes for SliverOpacity widget.
// See: https://api.flutter.dev/flutter/widgets/SliverOpacity-class.html
type SliverOpacityAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*SliverProps                      `gen:"wrap"`
	Opacity                           duit_utils.Tristate[float32] `json:"opacity,omitempty"`
}

//bindgen:exclude
func (r *SliverOpacityAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if r.Opacity == nil {
		return errors.New("opacity property is required")
	} else {
		return duit_utils.RangeValidator(*r.Opacity, 0, 1)
	}
}

// NewSliverOpacityAttributes creates a new SliverOpacityAttributes instance.
func NewSliverOpacityAttributes() *SliverOpacityAttributes {
	return &SliverOpacityAttributes{}
}

// SetOpacity sets the opacity for the sliver opacity widget.
func (r *SliverOpacityAttributes) SetOpacity(opacity duit_utils.Tristate[float32]) *SliverOpacityAttributes {
	r.Opacity = opacity
	return r
}
