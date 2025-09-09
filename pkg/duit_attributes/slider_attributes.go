package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// SliderAttributes represents attributes for Slider widget.
// https://api.flutter.dev/flutter/material/Slider-class.html
type SliderAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
	*ThemeConsumer        `gen:"wrap"`
	Value                 duit_utils.Tristate[float32] `json:"value,omitempty"`
	Min                   float32                      `json:"min,omitempty"`
	Max                   float32                      `json:"max,omitempty"`
	Divisions             uint32                       `json:"divisions,omitempty"`
	SecondaryTrackValue   float32                      `json:"secondaryTrackValue,omitempty"`
	OnChanged             any                          `json:"onChanged,omitempty"`
	OnChangeStart         any                          `json:"onChangeStart,omitempty"`
	OnChangeEnd           any                          `json:"onChangeEnd,omitempty"`
	ActiveColor           *duit_props.Color            `json:"activeColor,omitempty"`
	InactiveColor         *duit_props.Color            `json:"inactiveColor,omitempty"`
	ThumbColor            *duit_props.Color            `json:"thumbColor,omitempty"`
	SecondaryActiveColor  *duit_props.Color            `json:"secondaryActiveColor,omitempty"`
	OverlayColor          *duit_props.WidgetStateColor `json:"overlayColor,omitempty"`
	Autofocus             duit_utils.Tristate[bool]    `json:"autofocus,omitempty"`
	Label                 string                       `json:"label,omitempty"`
	AllowedInteraction    duit_props.SliderInteraction `json:"allowedInteraction,omitempty"`
}

// NewSliderAttributes creates a new instance of SliderAttributes.
func NewSliderAttributes() *SliderAttributes {
	return &SliderAttributes{}
}

// SetValue sets the value property and returns the SliderAttributes instance for method chaining.
func (r *SliderAttributes) SetValue(value float32) *SliderAttributes {
	r.Value = duit_utils.Float32Value(value)
	return r
}

// SetMin sets the min property and returns the SliderAttributes instance for method chaining.
func (r *SliderAttributes) SetMin(min float32) *SliderAttributes {
	r.Min = min
	return r
}

// SetMax sets the max property and returns the SliderAttributes instance for method chaining.
func (r *SliderAttributes) SetMax(max float32) *SliderAttributes {
	r.Max = max
	return r
}

// SetDivisions sets the divisions property and returns the SliderAttributes instance for method chaining.
func (r *SliderAttributes) SetDivisions(divisions uint32) *SliderAttributes {
	r.Divisions = divisions
	return r
}

// SetSecondaryTrackValue sets the secondaryTrackValue property and returns the SliderAttributes instance for method chaining.
func (r *SliderAttributes) SetSecondaryTrackValue(secondaryTrackValue float32) *SliderAttributes {
	r.SecondaryTrackValue = secondaryTrackValue
	return r
}

// SetOnChanged sets the onChanged property and returns the SliderAttributes instance for method chaining.
func (r *SliderAttributes) SetOnChanged(onChanged any) *SliderAttributes {
	r.OnChanged = onChanged
	return r
}

// SetOnChangeStart sets the onChangeStart property and returns the SliderAttributes instance for method chaining.
func (r *SliderAttributes) SetOnChangeStart(onChangeStart any) *SliderAttributes {
	r.OnChangeStart = onChangeStart
	return r
}

// SetOnChangeEnd sets the onChangeEnd property and returns the SliderAttributes instance for method chaining.
func (r *SliderAttributes) SetOnChangeEnd(onChangeEnd any) *SliderAttributes {
	r.OnChangeEnd = onChangeEnd
	return r
}

// SetActiveColor sets the activeColor property and returns the SliderAttributes instance for method chaining.
func (r *SliderAttributes) SetActiveColor(activeColor *duit_props.Color) *SliderAttributes {
	r.ActiveColor = activeColor
	return r
}

// SetInactiveColor sets the inactiveColor property and returns the SliderAttributes instance for method chaining.
func (r *SliderAttributes) SetInactiveColor(inactiveColor *duit_props.Color) *SliderAttributes {
	r.InactiveColor = inactiveColor
	return r
}

// SetThumbColor sets the thumbColor property and returns the SliderAttributes instance for method chaining.
func (r *SliderAttributes) SetThumbColor(thumbColor *duit_props.Color) *SliderAttributes {
	r.ThumbColor = thumbColor
	return r
}

// SetSecondaryActiveColor sets the secondaryActiveColor property and returns the SliderAttributes instance for method chaining.
func (r *SliderAttributes) SetSecondaryActiveColor(secondaryActiveColor *duit_props.Color) *SliderAttributes {
	r.SecondaryActiveColor = secondaryActiveColor
	return r
}

// SetOverlayColor sets the overlayColor property and returns the SliderAttributes instance for method chaining.
func (r *SliderAttributes) SetOverlayColor(overlayColor *duit_props.WidgetStateColor) *SliderAttributes {
	r.OverlayColor = overlayColor
	return r
}

// SetAutofocus sets the autofocus property and returns the SliderAttributes instance for method chaining.
func (r *SliderAttributes) SetAutofocus(autofocus bool) *SliderAttributes {
	r.Autofocus = duit_utils.BoolValue(autofocus)
	return r
}

// SetLabel sets the label property and returns the SliderAttributes instance for method chaining.
func (r *SliderAttributes) SetLabel(label string) *SliderAttributes {
	r.Label = label
	return r
}

// SetAllowedInteraction sets the allowedInteraction property and returns the SliderAttributes instance for method chaining.
func (r *SliderAttributes) SetAllowedInteraction(allowedInteraction duit_props.SliderInteraction) *SliderAttributes {
	r.AllowedInteraction = allowedInteraction
	return r
}

//bindgen:exclude
func (r *SliderAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if r.Value == nil {
		return errors.New("value property is required")
	}

	if r.Min >= r.Max {
		return errors.New("min must be less than max")
	}

	actions := []any{
		r.OnChanged,
		r.OnChangeStart,
		r.OnChangeEnd,
	}

	for _, action := range actions {
		if err := duit_action.CheckActionType(action); err != nil {
			return err
		}
	}

	return nil
}
