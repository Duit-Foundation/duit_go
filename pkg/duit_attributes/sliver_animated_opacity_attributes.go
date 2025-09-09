package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// SliverAnimatedOpacityAttributes represents attributes for SliverAnimatedOpacity widget.
// https://api.flutter.dev/flutter/widgets/SliverAnimatedOpacity-class.html
type SliverAnimatedpacityAttributes struct {
	*ValueReferenceHolder          `gen:"wrap"`
	*ThemeConsumer                 `gen:"wrap"`
	*duit_props.ImplicitAnimatable `gen:"wrap"`
	*SliverProps                   `gen:"wrap"`
	Opacity                        duit_utils.Tristate[float32] `json:"opacity,omitempty"`
}

// NewSliverAnimatedOpacityAttributes creates a new instance of SliverAnimatedOpacityAttributes.
func NewSliverAnimatedOpacityAttributes() *SliverAnimatedpacityAttributes {
	return &SliverAnimatedpacityAttributes{}
}

// SetOpacity sets the opacity property and returns the SliverAnimatedOpacityAttributes instance for method chaining.
func (r *SliverAnimatedpacityAttributes) SetOpacity(opacity float32) *SliverAnimatedpacityAttributes {
	r.Opacity = duit_utils.Float32Value(opacity)
	return r
}

//bindgen:exclude
func (r *SliverAnimatedpacityAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ImplicitAnimatable.Validate(); err != nil {
		return err
	}

	if r.Opacity == nil {
		return errors.New("opacity property is required")
	} else {
		return duit_utils.RangeValidator(*r.Opacity, 0, 1)
	}
}
