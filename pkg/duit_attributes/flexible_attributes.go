package duit_attributes

import (
	"errors"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// FlexibleAttributes represents attributes for Flexible widget.
// https://api.flutter.dev/flutter/widgets/Flexible-class.html
type FlexibleAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
	*ThemeConsumer        `gen:"wrap"`
	Flex                  uint16            `json:"flex,omitempty"`
	Fit                   duit_props.FlexFix `json:"fit,omitempty"`
}

// NewFlexibleAttributes creates a new instance of FlexibleAttributes.
func NewFlexibleAttributes() *FlexibleAttributes {
	return &FlexibleAttributes{}
}

// SetFlex sets the flex property and returns the FlexibleAttributes instance for method chaining.
func (r *FlexibleAttributes) SetFlex(flex uint16) *FlexibleAttributes {
	r.Flex = flex
	return r
}

// SetFit sets the fit property and returns the FlexibleAttributes instance for method chaining.
func (r *FlexibleAttributes) SetFit(fit duit_props.FlexFix) *FlexibleAttributes {
	r.Fit = fit
	return r
}

//bindgen:exclude
func (r *FlexibleAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if r.Flex < 1 {
		return errors.New("flex must be greater than 0")
	}

	return nil
}
