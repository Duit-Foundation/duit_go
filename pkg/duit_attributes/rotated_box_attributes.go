package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// RotatedBoxAttributes defines attributes for RotatedBox widget.
// See: https://api.flutter.dev/flutter/widgets/RotatedBox-class.html
type RotatedBoxAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	QuarterTurns                      duit_utils.Tristate[int] `json:"quarterTurns,omitempty"`
}

//bindgen:exclude
func (r *RotatedBoxAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if r.QuarterTurns == nil {
		return errors.New("quarterTurns property is required")
	}

	return nil
}

// NewRotatedBoxAttributes creates a new RotatedBoxAttributes instance.
func NewRotatedBoxAttributes() *RotatedBoxAttributes {
	return &RotatedBoxAttributes{}
}

// SetQuarterTurns sets the quarter turns for the rotated box widget.
func (r *RotatedBoxAttributes) SetQuarterTurns(quarterTurns duit_utils.Tristate[int]) *RotatedBoxAttributes {
	r.QuarterTurns = quarterTurns
	return r
}
