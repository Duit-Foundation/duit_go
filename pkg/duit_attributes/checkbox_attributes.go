package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// CheckboxAttributes defines attributes for Checkbox widget.
// See: https://api.flutter.dev/flutter/material/Checkbox-class.html
type CheckboxAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
	*ThemeConsumer        `gen:"wrap"`
	Value                 duit_utils.Tristate[bool]    `json:"value"`
	Autofocus             duit_utils.Tristate[bool]    `json:"autofocus,omitempty"`
	Tristate              duit_utils.Tristate[bool]    `json:"tristate,omitempty"`
	IsError               duit_utils.Tristate[bool]    `json:"isError,omitempty"`
	SplashRadius          float32                      `json:"splashRadius,omitempty"`
	SemanticLabel         string                       `json:"semanticLabel,omitempty"`
	Side                  *duit_props.BorderSide       `json:"side,omitempty"`
	VisualDensity         *duit_props.VisualDensity    `json:"visualDensity,omitempty"`
	CheckColor            *duit_props.Color            `json:"checkColor,omitempty"`
	HoverColor            *duit_props.Color            `json:"hoverColor,omitempty"`
	ActiveColor           *duit_props.Color            `json:"activeColor,omitempty"`
	FocusColor            *duit_props.Color            `json:"focusColor,omitempty"`
	FillColor             *duit_props.WidgetStateColor `json:"fillColor,omitempty"`
	OverlayColor          *duit_props.WidgetStateColor `json:"overlayColor,omitempty"`
}

//bindgen:exclude
func (r *CheckboxAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if r.Value == nil {
		return errors.New("value property is required")
	}

	return nil
}

// NewCheckboxAttributes creates a new CheckboxAttributes instance.
func NewCheckboxAttributes() *CheckboxAttributes {
	return &CheckboxAttributes{}
}

// SetValue sets the value for the checkbox widget.
func (r *CheckboxAttributes) SetValue(value duit_utils.Tristate[bool]) *CheckboxAttributes {
	r.Value = value
	return r
}

// SetAutofocus sets the autofocus for the checkbox widget.
func (r *CheckboxAttributes) SetAutofocus(autofocus duit_utils.Tristate[bool]) *CheckboxAttributes {
	r.Autofocus = autofocus
	return r
}

// SetTristate sets the tristate for the checkbox widget.
func (r *CheckboxAttributes) SetTristate(tristate duit_utils.Tristate[bool]) *CheckboxAttributes {
	r.Tristate = tristate
	return r
}

// SetIsError sets the is error for the checkbox widget.
func (r *CheckboxAttributes) SetIsError(isError duit_utils.Tristate[bool]) *CheckboxAttributes {
	r.IsError = isError
	return r
}

// SetSplashRadius sets the splash radius for the checkbox widget.
func (r *CheckboxAttributes) SetSplashRadius(splashRadius float32) *CheckboxAttributes {
	r.SplashRadius = splashRadius
	return r
}

// SetSemanticLabel sets the semantic label for the checkbox widget.
func (r *CheckboxAttributes) SetSemanticLabel(semanticLabel string) *CheckboxAttributes {
	r.SemanticLabel = semanticLabel
	return r
}

// SetSide sets the side for the checkbox widget.
func (r *CheckboxAttributes) SetSide(side *duit_props.BorderSide) *CheckboxAttributes {
	r.Side = side
	return r
}

// SetVisualDensity sets the visual density for the checkbox widget.
func (r *CheckboxAttributes) SetVisualDensity(visualDensity *duit_props.VisualDensity) *CheckboxAttributes {
	r.VisualDensity = visualDensity
	return r
}

// SetCheckColor sets the check color for the checkbox widget.
func (r *CheckboxAttributes) SetCheckColor(checkColor *duit_props.Color) *CheckboxAttributes {
	r.CheckColor = checkColor
	return r
}

// SetHoverColor sets the hover color for the checkbox widget.
func (r *CheckboxAttributes) SetHoverColor(hoverColor *duit_props.Color) *CheckboxAttributes {
	r.HoverColor = hoverColor
	return r
}

// SetActiveColor sets the active color for the checkbox widget.
func (r *CheckboxAttributes) SetActiveColor(activeColor *duit_props.Color) *CheckboxAttributes {
	r.ActiveColor = activeColor
	return r
}

// SetFocusColor sets the focus color for the checkbox widget.
func (r *CheckboxAttributes) SetFocusColor(focusColor *duit_props.Color) *CheckboxAttributes {
	r.FocusColor = focusColor
	return r
}

// SetFillColor sets the fill color for the checkbox widget.
func (r *CheckboxAttributes) SetFillColor(fillColor *duit_props.WidgetStateColor) *CheckboxAttributes {
	r.FillColor = fillColor
	return r
}

// SetOverlayColor sets the overlay color for the checkbox widget.
func (r *CheckboxAttributes) SetOverlayColor(overlayColor *duit_props.WidgetStateColor) *CheckboxAttributes {
	r.OverlayColor = overlayColor
	return r
}
