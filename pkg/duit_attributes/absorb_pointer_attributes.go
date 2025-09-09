package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// AbsorbPointerAttributes represents attributes for AbsorbPointer widget.
// https://api.flutter.dev/flutter/widgets/AbsorbPointer-class.html
type AbsorbPointerAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
	Absorbing             duit_utils.Tristate[bool] `json:"absorbing,omitempty"`
}

// Create a new AbsorbPointerAttributes
func NewAbsorbPointerAttributes() *AbsorbPointerAttributes {
	return &AbsorbPointerAttributes{}
}

// Set the absorbing value
func (r *AbsorbPointerAttributes) SetAbsorbing(absorbing bool) *AbsorbPointerAttributes {
	r.Absorbing = duit_utils.BoolValue(absorbing)
	return r
}

//bindgen:exclude
func (r *AbsorbPointerAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	return nil
}
