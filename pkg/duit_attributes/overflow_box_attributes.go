package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// OverflowBoxAttributes defines attributes for OverflowBox widget.
// See: https://api.flutter.dev/flutter/widgets/OverflowBox-class.html
type OverflowBoxAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	MinWidth                          float32                   `json:"minWidth,omitempty"`
	MaxWidth                          float32                   `json:"maxWidth,omitempty"`
	MinHeight                         float32                   `json:"minHeight,omitempty"`
	MaxHeight                         float32                   `json:"maxHeight,omitempty"`
	Alignment                         duit_props.Alignment      `json:"alignment,omitempty"`
	Fit                               duit_props.OverflowBoxFit `json:"overflowBoxFit,omitempty"`
}

//bindgen:exclude
func (r *OverflowBoxAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	return nil
}

// NewOverflowBoxAttributes creates a new OverflowBoxAttributes instance.
func NewOverflowBoxAttributes() *OverflowBoxAttributes {
	return &OverflowBoxAttributes{}
}

// SetMinWidth sets the min width for the overflow box widget.
func (r *OverflowBoxAttributes) SetMinWidth(minWidth float32) *OverflowBoxAttributes {
	r.MinWidth = minWidth
	return r
}

// SetMaxWidth sets the max width for the overflow box widget.
func (r *OverflowBoxAttributes) SetMaxWidth(maxWidth float32) *OverflowBoxAttributes {
	r.MaxWidth = maxWidth
	return r
}

// SetMinHeight sets the min height for the overflow box widget.
func (r *OverflowBoxAttributes) SetMinHeight(minHeight float32) *OverflowBoxAttributes {
	r.MinHeight = minHeight
	return r
}

// SetMaxHeight sets the max height for the overflow box widget.
func (r *OverflowBoxAttributes) SetMaxHeight(maxHeight float32) *OverflowBoxAttributes {
	r.MaxHeight = maxHeight
	return r
}

// SetAlignment sets the alignment for the overflow box widget.
func (r *OverflowBoxAttributes) SetAlignment(alignment duit_props.Alignment) *OverflowBoxAttributes {
	r.Alignment = alignment
	return r
}

// SetFit sets the fit for the overflow box widget.
func (r *OverflowBoxAttributes) SetFit(fit duit_props.OverflowBoxFit) *OverflowBoxAttributes {
	r.Fit = fit
	return r
}
