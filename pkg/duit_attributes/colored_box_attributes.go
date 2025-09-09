package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// ColoredBoxAttributes defines attributes for ColoredBox widget.
// See: https://api.flutter.dev/flutter/widgets/ColoredBox-class.html
type ColoredBoxAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	Color                             *duit_props.Color `json:"color,omitempty"`
}

//bindgen:exclude
func (r *ColoredBoxAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if r.Color == nil {
		return errors.New("color property is required")
	} else {
		if err := r.Color.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// NewColoredBoxAttributes creates a new ColoredBoxAttributes instance.
func NewColoredBoxAttributes() *ColoredBoxAttributes {
	return &ColoredBoxAttributes{}
}

// SetColor sets the color for the colored box widget.
func (r *ColoredBoxAttributes) SetColor(color *duit_props.Color) *ColoredBoxAttributes {
	r.Color = color
	return r
}
