package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// OffstageAttributes defines attributes for Offstage widget.
// See: https://api.flutter.dev/flutter/widgets/Offstage-class.html
type OffstageAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
	Offstage              duit_utils.Tristate[bool] `json:"offstage,omitempty"`
}

//bindgen:exclude
func (r *OffstageAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if r.Offstage == nil {
		return errors.New("offstage property is required")
	}

	return nil
}

// NewOffstageAttributes creates a new OffstageAttributes instance.
func NewOffstageAttributes() *OffstageAttributes {
	return &OffstageAttributes{}
}

// SetOffstage sets the offstage for the offstage widget.
func (r *OffstageAttributes) SetOffstage(offstage duit_utils.Tristate[bool]) *OffstageAttributes {
	r.Offstage = offstage
	return r
}
