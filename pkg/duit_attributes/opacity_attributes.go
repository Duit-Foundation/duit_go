package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// OpacityAttributes defines attributes for Opacity widget.
// See: https://api.flutter.dev/flutter/widgets/Opacity-class.html
type OpacityAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	Opacity                           duit_utils.Tristate[float32] `json:"opacity"`
}

//bindgen:exclude
func (r *OpacityAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if r.Opacity == nil {
		return errors.New("opacity property is required")
	} else {
		return duit_utils.RangeValidator(*r.Opacity, 0.0, 1.1)
	}
}

// NewOpacityAttributes creates a new OpacityAttributes instance.
func NewOpacityAttributes() *OpacityAttributes {
	return &OpacityAttributes{}
}

// SetOpacity sets the opacity for the opacity widget.
func (r *OpacityAttributes) SetOpacity(opacity duit_utils.Tristate[float32]) *OpacityAttributes {
	r.Opacity = opacity
	return r
}
