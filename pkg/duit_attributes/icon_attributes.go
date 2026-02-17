package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// IconAttributes defines attributes for Icon widget.
// See: https://api.flutter.dev/flutter/widgets/Icon-class.html
type IconAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	Icon                              *duit_props.IconData       `json:"icon,omitempty"`
	Size                              duit_utils.Tristate[float32] `json:"size,omitempty"`
	IconSize                          duit_utils.Tristate[float32] `json:"iconSize,omitempty"`
	Fill                              duit_utils.Tristate[float32] `json:"fill,omitempty"`
	Weight                            duit_utils.Tristate[float32] `json:"weight,omitempty"`
	Grade                             duit_utils.Tristate[float32] `json:"grade,omitempty"`
	Color                             *duit_props.Color           `json:"color,omitempty"`
	TextDirection                     duit_props.TextDirection     `json:"textDirection,omitempty"`
	ApplyTextScaling                  duit_utils.Tristate[bool]   `json:"applyTextScaling,omitempty"`
	SemanticLabel                     string                     `json:"semanticLabel,omitempty"`
}

// NewIconAttributes creates a new IconAttributes instance.
func NewIconAttributes() *IconAttributes {
	return &IconAttributes{}
}

// SetIcon sets the icon property and returns IconAttributes for method chaining.
func (r *IconAttributes) SetIcon(icon *duit_props.IconData) *IconAttributes {
	r.Icon = icon
	return r
}

// SetSize sets the size property and returns IconAttributes for method chaining.
func (r *IconAttributes) SetSize(size float32) *IconAttributes {
	r.Size = duit_utils.Float32Value(size)
	return r
}

// SetIconSize sets the iconSize property and returns IconAttributes for method chaining.
func (r *IconAttributes) SetIconSize(size float32) *IconAttributes {
	r.IconSize = duit_utils.Float32Value(size)
	return r
}

// SetFill sets the fill property and returns IconAttributes for method chaining.
func (r *IconAttributes) SetFill(fill float32) *IconAttributes {
	r.Fill = duit_utils.Float32Value(fill)
	return r
}

// SetWeight sets the weight property and returns IconAttributes for method chaining.
func (r *IconAttributes) SetWeight(weight float32) *IconAttributes {
	r.Weight = duit_utils.Float32Value(weight)
	return r
}

// SetGrade sets the grade property and returns IconAttributes for method chaining.
func (r *IconAttributes) SetGrade(grade float32) *IconAttributes {
	r.Grade = duit_utils.Float32Value(grade)
	return r
}

// SetColor sets the color property and returns IconAttributes for method chaining.
func (r *IconAttributes) SetColor(color *duit_props.Color) *IconAttributes {
	r.Color = color
	return r
}

// SetTextDirection sets the textDirection property and returns IconAttributes for method chaining.
func (r *IconAttributes) SetTextDirection(textDirection duit_props.TextDirection) *IconAttributes {
	r.TextDirection = textDirection
	return r
}

// SetApplyTextScaling sets the applyTextScaling property and returns IconAttributes for method chaining.
func (r *IconAttributes) SetApplyTextScaling(applyTextScaling bool) *IconAttributes {
	r.ApplyTextScaling = duit_utils.BoolValue(applyTextScaling)
	return r
}

// SetSemanticLabel sets the semanticLabel property and returns IconAttributes for method chaining.
func (r *IconAttributes) SetSemanticLabel(semanticLabel string) *IconAttributes {
	r.SemanticLabel = semanticLabel
	return r
}

//bindgen:exclude
func (r *IconAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if r.Icon == nil {
		return errors.New("icon is required")
	}

	return nil
}
