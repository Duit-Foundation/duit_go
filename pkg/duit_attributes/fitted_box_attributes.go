package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// FittedBoxAttributes defines attributes for FittedBox widget.
// See: https://api.flutter.dev/flutter/widgets/FittedBox-class.html
type FittedBoxAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	Fit                               duit_props.BoxFit    `json:"fit,omitempty"`
	ClipBehavior                      duit_props.Clip      `json:"clipBehavior,omitempty"`
	Alignment                         duit_props.Alignment `json:"alignment,omitempty"`
}

//bindgen:exclude
func (r *FittedBoxAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	return nil
}

// NewFittedBoxAttributes creates a new FittedBoxAttributes instance.
func NewFittedBoxAttributes() *FittedBoxAttributes {
	return &FittedBoxAttributes{}
}

// SetFit sets the fit for the fitted box widget.
func (r *FittedBoxAttributes) SetFit(fit duit_props.BoxFit) *FittedBoxAttributes {
	r.Fit = fit
	return r
}

// SetClipBehavior sets the clip behavior for the fitted box widget.
func (r *FittedBoxAttributes) SetClipBehavior(clipBehavior duit_props.Clip) *FittedBoxAttributes {
	r.ClipBehavior = clipBehavior
	return r
}

// SetAlignment sets the alignment for the fitted box widget.
func (r *FittedBoxAttributes) SetAlignment(alignment duit_props.Alignment) *FittedBoxAttributes {
	r.Alignment = alignment
	return r
}
