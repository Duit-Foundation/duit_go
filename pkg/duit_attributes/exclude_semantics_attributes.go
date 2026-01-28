package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// ExcludeSemanticsAttributes defines attributes for ExcludeSemantics widget.
// See: https://api.flutter.dev/flutter/widgets/ExcludeSemantics-class.html
type ExcludeSemanticsAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
	Excluding             duit_utils.Tristate[bool] `json:"excluding,omitempty"`
}

// NewExcludeSemanticsAttributes creates a new ExcludeSemanticsAttributes instance.
func NewExcludeSemanticsAttributes() *ExcludeSemanticsAttributes {
	return &ExcludeSemanticsAttributes{}
}

// SetExcluding sets the excluding property. When true, this widget (and its subtree) is excluded from the semantics tree.
func (r *ExcludeSemanticsAttributes) SetExcluding(excluding bool) *ExcludeSemanticsAttributes {
	r.Excluding = duit_utils.BoolValue(excluding)
	return r
}

//bindgen:exclude
func (r *ExcludeSemanticsAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if r.Excluding == nil {
		return errors.New("excluding property is required")
	}

	return nil
}
