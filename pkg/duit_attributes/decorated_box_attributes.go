package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// DecoratedBoxAttributes defines attributes for DecoratedBox widget.
// See: https://api.flutter.dev/flutter/widgets/DecoratedBox-class.html
type DecoratedBoxAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	Decoration                        *duit_props.BoxDecoration `json:"decoration,omitempty"`
}

//bindgen:exclude
func (r *DecoratedBoxAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if r.Decoration == nil {
		return errors.New("decoration property is required ")
	} else {
		if err := r.Decoration.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// NewDecoratedBoxAttributes creates a new DecoratedBoxAttributes instance.
func NewDecoratedBoxAttributes() *DecoratedBoxAttributes {
	return &DecoratedBoxAttributes{}
}

// SetDecoration sets the decoration for the decorated box widget.
func (r *DecoratedBoxAttributes) SetDecoration(decoration *duit_props.BoxDecoration) *DecoratedBoxAttributes {
	r.Decoration = decoration
	return r
}
