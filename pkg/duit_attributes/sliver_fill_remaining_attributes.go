package duit_attributes

import "github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"

// SliverFillRemainingAttributes defines attributes for SliverFillRemaining widget.
// See: https://api.flutter.dev/flutter/widgets/SliverFillRemaining-class.html
type SliverFillRemainingAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
	*ThemeConsumer        `gen:"wrap"`
	*SliverProps          `gen:"wrap"`
	HasScrollBody         duit_utils.Tristate[bool] `json:"hasScrollBody,omitempty"`
	FillOverscroll        duit_utils.Tristate[bool] `json:"fillOverscroll,omitempty"`
}

//bindgen:exclude
func (r *SliverFillRemainingAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	return nil
}

// NewSliverFillRemainingAttributes creates a new SliverFillRemainingAttributes instance.
func NewSliverFillRemainingAttributes() *SliverFillRemainingAttributes {
	return &SliverFillRemainingAttributes{}
}

// SetHasScrollBody sets the has scroll body for the sliver fill remaining widget.
func (r *SliverFillRemainingAttributes) SetHasScrollBody(hasScrollBody duit_utils.Tristate[bool]) *SliverFillRemainingAttributes {
	r.HasScrollBody = hasScrollBody
	return r
}

// SetFillOverscroll sets the fill overscroll for the sliver fill remaining widget.
func (r *SliverFillRemainingAttributes) SetFillOverscroll(fillOverscroll duit_utils.Tristate[bool]) *SliverFillRemainingAttributes {
	r.FillOverscroll = fillOverscroll
	return r
}
