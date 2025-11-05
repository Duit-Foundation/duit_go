package duit_attributes

import "errors"

// ExpandedAttributes represents attributes for Expanded widget.
// https://api.flutter.dev/flutter/widgets/Expanded-class.html
type ExpandedAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
	*ThemeConsumer        `gen:"wrap"`
	Flex                  uint16 `json:"flex,omitempty"`
}

// NewExpandedAttributes creates a new instance of ExpandedAttributes.
func NewExpandedAttributes() *ExpandedAttributes {
	return &ExpandedAttributes{}
}

// SetFlex sets the flex property and returns the ExpandedAttributes instance for method chaining.
func (r *ExpandedAttributes) SetFlex(flex uint16) *ExpandedAttributes {
	r.Flex = flex
	return r
}

//bindgen:exclude
func (r *ExpandedAttributes) Validate() error {
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
