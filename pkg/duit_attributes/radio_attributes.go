package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type PrimitiveValue interface {
	string | int | uint | uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64 | float32 | float64 | bool
}

// RadioAttributes defines attributes for Radio widget.
// See: https://api.flutter.dev/flutter/material/Radio-class.html
type RadioAttributes[TValue PrimitiveValue] struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	Value                             duit_utils.Tristate[TValue]      `json:"value,omitempty"`
	Toggleable                        bool                             `json:"toggleable,omitempty"`
	Autofocus                         bool                             `json:"autofocus,omitempty"`
	ActiveColor                       *duit_props.Color                `json:"activeColor,omitempty"`
	FocusColor                        *duit_props.Color                `json:"focusColor,omitempty"`
	HoverColor                        *duit_props.Color                `json:"hoverColor,omitempty"`
	FillColor                         *duit_props.WidgetStateColor     `json:"fillColor,omitempty"`
	OverlayColor                      *duit_props.WidgetStateColor     `json:"overlayColor,omitempty"`
	SplashRadius                      float32                          `json:"splashRadius,omitempty"`
	VisualDensity                     *duit_props.VisualDensity        `json:"visualDensity,omitempty"`
	MaterialTapTargetSize             duit_props.MaterialTapTargetSize `json:"materialTapTargetSize,omitempty"`
}

//bindgen:exclude
func (r *RadioAttributes[TValue]) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if r.Value == nil {
		return errors.New("value property is required")
	}

	return nil
}

// NewRadioAttributes creates a new RadioAttributes instance.
func NewRadioAttributes[TValue PrimitiveValue]() *RadioAttributes[TValue] {
	return &RadioAttributes[TValue]{}
}

// SetValue sets the value for the radio widget.
func (r *RadioAttributes[TValue]) SetValue(value duit_utils.Tristate[TValue]) *RadioAttributes[TValue] {
	r.Value = value
	return r
}

// SetToggleable sets the toggleable for the radio widget.
func (r *RadioAttributes[TValue]) SetToggleable(toggleable bool) *RadioAttributes[TValue] {
	r.Toggleable = toggleable
	return r
}

// SetAutofocus sets the autofocus for the radio widget.
func (r *RadioAttributes[TValue]) SetAutofocus(autofocus bool) *RadioAttributes[TValue] {
	r.Autofocus = autofocus
	return r
}

// SetActiveColor sets the active color for the radio widget.
func (r *RadioAttributes[TValue]) SetActiveColor(activeColor *duit_props.Color) *RadioAttributes[TValue] {
	r.ActiveColor = activeColor
	return r
}

// SetFocusColor sets the focus color for the radio widget.
func (r *RadioAttributes[TValue]) SetFocusColor(focusColor *duit_props.Color) *RadioAttributes[TValue] {
	r.FocusColor = focusColor
	return r
}

// SetHoverColor sets the hover color for the radio widget.
func (r *RadioAttributes[TValue]) SetHoverColor(hoverColor *duit_props.Color) *RadioAttributes[TValue] {
	r.HoverColor = hoverColor
	return r
}

// SetFillColor sets the fill color for the radio widget.
func (r *RadioAttributes[TValue]) SetFillColor(fillColor *duit_props.WidgetStateColor) *RadioAttributes[TValue] {
	r.FillColor = fillColor
	return r
}

// SetOverlayColor sets the overlay color for the radio widget.
func (r *RadioAttributes[TValue]) SetOverlayColor(overlayColor *duit_props.WidgetStateColor) *RadioAttributes[TValue] {
	r.OverlayColor = overlayColor
	return r
}

// SetSplashRadius sets the splash radius for the radio widget.
func (r *RadioAttributes[TValue]) SetSplashRadius(splashRadius float32) *RadioAttributes[TValue] {
	r.SplashRadius = splashRadius
	return r
}

// SetVisualDensity sets the visual density for the radio widget.
func (r *RadioAttributes[TValue]) SetVisualDensity(visualDensity *duit_props.VisualDensity) *RadioAttributes[TValue] {
	r.VisualDensity = visualDensity
	return r
}

// SetMaterialTapTargetSize sets the material tap target size for the radio widget.
func (r *RadioAttributes[TValue]) SetMaterialTapTargetSize(materialTapTargetSize duit_props.MaterialTapTargetSize) *RadioAttributes[TValue] {
	r.MaterialTapTargetSize = materialTapTargetSize
	return r
}

// RadioGroupContextAttributes defines attributes for RadioGroupContext widget.
// See: https://api.flutter.dev/flutter/material/Radio-class.html
type RadioGroupContextAttributes[TValue PrimitiveValue] struct {
	*ValueReferenceHolder `gen:"wrap"`
	GroupValue            duit_utils.Tristate[TValue] `json:"groupValue,omitempty"`
}

//bindgen:exclude
func (r *RadioGroupContextAttributes[TValue]) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if r.GroupValue == nil {
		return errors.New("groupValue property is required")
	}

	return nil
}

// NewRadioGroupContextAttributes creates a new RadioGroupContextAttributes instance.
func NewRadioGroupContextAttributes[TValue PrimitiveValue]() *RadioGroupContextAttributes[TValue] {
	return &RadioGroupContextAttributes[TValue]{}
}

// SetGroupValue sets the group value for the radio group context widget.
func (r *RadioGroupContextAttributes[TValue]) SetGroupValue(groupValue duit_utils.Tristate[TValue]) *RadioGroupContextAttributes[TValue] {
	r.GroupValue = groupValue
	return r
}
