package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// SliverOffstageAttributes defines attributes for SliverOffstage widget.
// See: https://api.flutter.dev/flutter/widgets/SliverOffstage-class.html
type SliverOffstageAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
	*ThemeConsumer        `gen:"wrap"`
	*SliverProps          `gen:"wrap"`
	Offstage              duit_utils.Tristate[bool] `json:"offstage,omitempty"`
}

//bindgen:exclude
func (r *SliverOffstageAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if r.Offstage == nil {
		return errors.New("offstage property is required")
	}

	return nil
}

// NewSliverOffstageAttributes creates a new SliverOffstageAttributes instance.
func NewSliverOffstageAttributes() *SliverOffstageAttributes {
	return &SliverOffstageAttributes{}
}

// SetOffstage sets the offstage for the sliver offstage widget.
func (r *SliverOffstageAttributes) SetOffstage(offstage duit_utils.Tristate[bool]) *SliverOffstageAttributes {
	r.Offstage = offstage
	return r
}
