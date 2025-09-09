package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// CardAttributes defines attributes for Card widget.
// See: https://api.flutter.dev/flutter/material/Card-class.html
type CardAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	Color                             *duit_props.Color         `json:"color,omitempty"`
	ShadowColor                       *duit_props.Color         `json:"shadowColor,omitempty"`
	Elevation                         float32                   `json:"elevation,omitempty"`
	BorderOnForeground                duit_utils.Tristate[bool] `json:"borderOnForeground,omitempty"`
	SemanticContainer                 duit_utils.Tristate[bool] `json:"semanticContainer,omitempty"`
	Margin                            *duit_props.EdgeInsets    `json:"margin,omitempty"`
	Shape                             *duit_props.ShapeBorder   `json:"shape,omitempty"`
}

//bindgen:exclude
func (r *CardAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if r.Shape != nil {
		if err := r.Shape.Validate(); err != nil {
			return err
		}
	}

	if r.Margin != nil {
		if err := r.Margin.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// NewCardAttributes creates a new CardAttributes instance.
func NewCardAttributes() *CardAttributes {
	return &CardAttributes{}
}

// SetColor sets the color for the card widget.
func (r *CardAttributes) SetColor(color *duit_props.Color) *CardAttributes {
	r.Color = color
	return r
}

// SetShadowColor sets the shadow color for the card widget.
func (r *CardAttributes) SetShadowColor(shadowColor *duit_props.Color) *CardAttributes {
	r.ShadowColor = shadowColor
	return r
}

// SetElevation sets the elevation for the card widget.
func (r *CardAttributes) SetElevation(elevation float32) *CardAttributes {
	r.Elevation = elevation
	return r
}

// SetBorderOnForeground sets the border on foreground for the card widget.
func (r *CardAttributes) SetBorderOnForeground(borderOnForeground duit_utils.Tristate[bool]) *CardAttributes {
	r.BorderOnForeground = borderOnForeground
	return r
}

// SetSemanticContainer sets the semantic container for the card widget.
func (r *CardAttributes) SetSemanticContainer(semanticContainer duit_utils.Tristate[bool]) *CardAttributes {
	r.SemanticContainer = semanticContainer
	return r
}

// SetMargin sets the margin for the card widget.
func (r *CardAttributes) SetMargin(margin *duit_props.EdgeInsets) *CardAttributes {
	r.Margin = margin
	return r
}

// SetShape sets the shape for the card widget.
func (r *CardAttributes) SetShape(shape *duit_props.ShapeBorder) *CardAttributes {
	r.Shape = shape
	return r
}
