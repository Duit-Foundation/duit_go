package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// RadioAttributes defines attributes for Radio widget.
// See: https://api.flutter.dev/flutter/material/Radio-class.html
type RadioAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	Value                             duit_utils.Tristate[any]         `json:"value,omitempty"`
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
func (r *RadioAttributes) Validate() error {
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
func NewRadioAttributes() *RadioAttributes {
	return &RadioAttributes{}
}

// SetValue sets the value for the radio widget.
func (r *RadioAttributes) SetValue(value any) *RadioAttributes {
	r.Value = duit_utils.TristateFrom[any](value)
	return r
}

// SetToggleable sets the toggleable for the radio widget.
func (r *RadioAttributes) SetToggleable(toggleable bool) *RadioAttributes {
	r.Toggleable = toggleable
	return r
}

// SetAutofocus sets the autofocus for the radio widget.
func (r *RadioAttributes) SetAutofocus(autofocus bool) *RadioAttributes {
	r.Autofocus = autofocus
	return r
}

// SetActiveColor sets the active color for the radio widget.
func (r *RadioAttributes) SetActiveColor(activeColor *duit_props.Color) *RadioAttributes {
	r.ActiveColor = activeColor
	return r
}

// SetFocusColor sets the focus color for the radio widget.
func (r *RadioAttributes) SetFocusColor(focusColor *duit_props.Color) *RadioAttributes {
	r.FocusColor = focusColor
	return r
}

// SetHoverColor sets the hover color for the radio widget.
func (r *RadioAttributes) SetHoverColor(hoverColor *duit_props.Color) *RadioAttributes {
	r.HoverColor = hoverColor
	return r
}

// SetFillColor sets the fill color for the radio widget.
func (r *RadioAttributes) SetFillColor(fillColor *duit_props.WidgetStateColor) *RadioAttributes {
	r.FillColor = fillColor
	return r
}

// SetOverlayColor sets the overlay color for the radio widget.
func (r *RadioAttributes) SetOverlayColor(overlayColor *duit_props.WidgetStateColor) *RadioAttributes {
	r.OverlayColor = overlayColor
	return r
}

// SetSplashRadius sets the splash radius for the radio widget.
func (r *RadioAttributes) SetSplashRadius(splashRadius float32) *RadioAttributes {
	r.SplashRadius = splashRadius
	return r
}

// SetVisualDensity sets the visual density for the radio widget.
func (r *RadioAttributes) SetVisualDensity(visualDensity *duit_props.VisualDensity) *RadioAttributes {
	r.VisualDensity = visualDensity
	return r
}

// SetMaterialTapTargetSize sets the material tap target size for the radio widget.
func (r *RadioAttributes) SetMaterialTapTargetSize(materialTapTargetSize duit_props.MaterialTapTargetSize) *RadioAttributes {
	r.MaterialTapTargetSize = materialTapTargetSize
	return r
}

// RadioGroupContextAttributes defines attributes for RadioGroupContext widget.
// See: https://api.flutter.dev/flutter/material/Radio-class.html
type RadioGroupContextAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
	GroupValue            duit_utils.Tristate[any] `json:"groupValue,omitempty"`
}

//bindgen:exclude
func (r *RadioGroupContextAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if r.GroupValue == nil {
		return errors.New("groupValue property is required")
	}

	return nil
}

// NewRadioGroupContextAttributes creates a new RadioGroupContextAttributes instance.
func NewRadioGroupContextAttributes() *RadioGroupContextAttributes {
	return &RadioGroupContextAttributes{}
}

// SetGroupValue sets the group value for the radio group context widget.
func (r *RadioGroupContextAttributes) SetGroupValue(groupValue any) *RadioGroupContextAttributes {
	r.GroupValue = duit_utils.TristateFrom[any](groupValue)
	return r
}
