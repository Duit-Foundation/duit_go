package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// SwitchAttributes represents attributes for Switch widget.
// https://api.flutter.dev/flutter/material/Switch-class.html
type SwitchAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
	*ThemeConsumer        `gen:"wrap"`
	*FocusCapability      `gen:"wrap"`
	Value                 duit_utils.Tristate[bool]        `json:"value,omitempty"`
	ActiveColor           *duit_props.Color                `json:"activeColor,omitempty"`
	FocusColor            *duit_props.Color                `json:"focusColor,omitempty"`
	HoverColor            *duit_props.Color                `json:"hoverColor,omitempty"`
	ActiveTrackColor      *duit_props.Color                `json:"activeTrackColor,omitempty"`
	InactiveTrackColor    *duit_props.Color                `json:"inactiveTrackColor,omitempty"`
	OverlayColor          *duit_props.WidgetStateColor     `json:"overlayColor,omitempty"`
	TrackColor            *duit_props.WidgetStateColor     `json:"trackColor,omitempty"`
	ThumbColor            *duit_props.WidgetStateColor     `json:"thumbColor,omitempty"`
	TrackOutlineColor     *duit_props.WidgetStateColor     `json:"trackOutlineColor,omitempty"`
	TrackOutlineWidth     *duit_props.WidgetStateFloat32   `json:"trackOutlineWidth,omitempty"`
	SplashRadius          float32                          `json:"splashRadius,omitempty"`
	MaterialTapTargetSize duit_props.MaterialTapTargetSize `json:"materialTapTargetSize,omitempty"`
	Autofocus             duit_utils.Tristate[bool]        `json:"autofocus,omitempty"`
}

// NewSwitchAttributes creates a new instance of SwitchAttributes.
func NewSwitchAttributes() *SwitchAttributes {
	return &SwitchAttributes{}
}

// SetValue sets the value property and returns the SwitchAttributes instance for method chaining.
func (r *SwitchAttributes) SetValue(value bool) *SwitchAttributes {
	r.Value = duit_utils.BoolValue(value)
	return r
}

// SetActiveColor sets the activeColor property and returns the SwitchAttributes instance for method chaining.
func (r *SwitchAttributes) SetActiveColor(activeColor *duit_props.Color) *SwitchAttributes {
	r.ActiveColor = activeColor
	return r
}

// SetFocusColor sets the focusColor property and returns the SwitchAttributes instance for method chaining.
func (r *SwitchAttributes) SetFocusColor(focusColor *duit_props.Color) *SwitchAttributes {
	r.FocusColor = focusColor
	return r
}

// SetHoverColor sets the hoverColor property and returns the SwitchAttributes instance for method chaining.
func (r *SwitchAttributes) SetHoverColor(hoverColor *duit_props.Color) *SwitchAttributes {
	r.HoverColor = hoverColor
	return r
}

// SetActiveTrackColor sets the activeTrackColor property and returns the SwitchAttributes instance for method chaining.
func (r *SwitchAttributes) SetActiveTrackColor(activeTrackColor *duit_props.Color) *SwitchAttributes {
	r.ActiveTrackColor = activeTrackColor
	return r
}

// SetInactiveTrackColor sets the inactiveTrackColor property and returns the SwitchAttributes instance for method chaining.
func (r *SwitchAttributes) SetInactiveTrackColor(inactiveTrackColor *duit_props.Color) *SwitchAttributes {
	r.InactiveTrackColor = inactiveTrackColor
	return r
}

// SetOverlayColor sets the overlayColor property and returns the SwitchAttributes instance for method chaining.
func (r *SwitchAttributes) SetOverlayColor(overlayColor *duit_props.WidgetStateColor) *SwitchAttributes {
	r.OverlayColor = overlayColor
	return r
}

// SetTrackColor sets the trackColor property and returns the SwitchAttributes instance for method chaining.
func (r *SwitchAttributes) SetTrackColor(trackColor *duit_props.WidgetStateColor) *SwitchAttributes {
	r.TrackColor = trackColor
	return r
}

// SetThumbColor sets the thumbColor property and returns the SwitchAttributes instance for method chaining.
func (r *SwitchAttributes) SetThumbColor(thumbColor *duit_props.WidgetStateColor) *SwitchAttributes {
	r.ThumbColor = thumbColor
	return r
}

// SetTrackOutlineColor sets the trackOutlineColor property and returns the SwitchAttributes instance for method chaining.
func (r *SwitchAttributes) SetTrackOutlineColor(trackOutlineColor *duit_props.WidgetStateColor) *SwitchAttributes {
	r.TrackOutlineColor = trackOutlineColor
	return r
}

// SetTrackOutlineWidth sets the trackOutlineWidth property and returns the SwitchAttributes instance for method chaining.
func (r *SwitchAttributes) SetTrackOutlineWidth(trackOutlineWidth *duit_props.WidgetStateFloat32) *SwitchAttributes {
	r.TrackOutlineWidth = trackOutlineWidth
	return r
}

// SetSplashRadius sets the splashRadius property and returns the SwitchAttributes instance for method chaining.
func (r *SwitchAttributes) SetSplashRadius(splashRadius float32) *SwitchAttributes {
	r.SplashRadius = splashRadius
	return r
}

// SetMaterialTapTargetSize sets the materialTapTargetSize property and returns the SwitchAttributes instance for method chaining.
func (r *SwitchAttributes) SetMaterialTapTargetSize(materialTapTargetSize duit_props.MaterialTapTargetSize) *SwitchAttributes {
	r.MaterialTapTargetSize = materialTapTargetSize
	return r
}

// SetAutofocus sets the autofocus property and returns the SwitchAttributes instance for method chaining.
func (r *SwitchAttributes) SetAutofocus(autofocus bool) *SwitchAttributes {
	r.Autofocus = duit_utils.BoolValue(autofocus)
	return r
}

//bindgen:exclude
func (r *SwitchAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
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
