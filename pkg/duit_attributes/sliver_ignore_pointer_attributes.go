package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// SliverIgnorePointerAttributes defines attributes for SliverIgnorePointer widget.
// See: https://api.flutter.dev/flutter/widgets/SliverIgnorePointer-class.html
type SliverIgnorePointerAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
	*ThemeConsumer        `gen:"wrap"`
	*SliverProps          `gen:"wrap"`
	Ignoring              duit_utils.Tristate[bool] `json:"ignoring"`
}

//bindgen:exclude
func (r *SliverIgnorePointerAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if r.Ignoring == nil {
		return errors.New("ignoring property is required")
	}

	return nil
}

// NewSliverIgnorePointerAttributes creates a new SliverIgnorePointerAttributes instance.
func NewSliverIgnorePointerAttributes() *SliverIgnorePointerAttributes {
	return &SliverIgnorePointerAttributes{}
}

// SetIgnoring sets the ignoring for the sliver ignore pointer widget.
func (r *SliverIgnorePointerAttributes) SetIgnoring(ignoring duit_utils.Tristate[bool]) *SliverIgnorePointerAttributes {
	r.Ignoring = ignoring
	return r
}
