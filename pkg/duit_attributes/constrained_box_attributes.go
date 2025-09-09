package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// ConstrainedBoxAttributes defines attributes for ConstrainedBox widget.
// See: https://api.flutter.dev/flutter/widgets/ConstrainedBox-class.html
type ConstrainedBoxAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	Constraints                       *duit_props.BoxConstraints `json:"constraints,omitempty"`
}

//bindgen:exclude
func (r *ConstrainedBoxAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if r.Constraints != nil {
		if err := r.Constraints.Validate(); err != nil {
			return err
		}
	} else {
		return errors.New("constraints property is required")
	}

	return nil
}

// NewConstrainedBoxAttributes creates a new ConstrainedBoxAttributes instance.
func NewConstrainedBoxAttributes() *ConstrainedBoxAttributes {
	return &ConstrainedBoxAttributes{}
}

// SetConstraints sets the constraints for the constrained box widget.
func (r *ConstrainedBoxAttributes) SetConstraints(constraints *duit_props.BoxConstraints) *ConstrainedBoxAttributes {
	r.Constraints = constraints
	return r
}
