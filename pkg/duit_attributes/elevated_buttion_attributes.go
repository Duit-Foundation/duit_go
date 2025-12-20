package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// ElevatedButtonAttributes defines attributes for ElevatedButton widget.
// See: https://api.flutter.dev/flutter/material/ElevatedButton-class.html
type ElevatedButtonAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
	*ThemeConsumer        `gen:"wrap"`
	*FocusCapability      `gen:"wrap"`
	Autofocus             duit_utils.Tristate[bool] `json:"autofocus,omitempty"`
	ClipBehavior          duit_props.Clip           `json:"clipBehavior,omitempty"`
	Style                 *duit_props.ButtonStyle   `json:"style,omitempty"`
	OnLongPress           any                       `json:"onLongPress,omitempty"`
}

//bindgen:exclude
func (r *ElevatedButtonAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if r.OnLongPress != nil {
		if err := duit_action.CheckActionType(r.OnLongPress); err != nil {
			return err
		}
	}

	return nil
}

// NewElevatedButtonAttributes creates a new ElevatedButtonAttributes instance.
func NewElevatedButtonAttributes() *ElevatedButtonAttributes {
	return &ElevatedButtonAttributes{}
}

// SetAutofocus sets the autofocus for the elevated button widget.
func (r *ElevatedButtonAttributes) SetAutofocus(autofocus duit_utils.Tristate[bool]) *ElevatedButtonAttributes {
	r.Autofocus = autofocus
	return r
}

// SetClipBehavior sets the clip behavior for the elevated button widget.
func (r *ElevatedButtonAttributes) SetClipBehavior(clipBehavior duit_props.Clip) *ElevatedButtonAttributes {
	r.ClipBehavior = clipBehavior
	return r
}

// SetStyle sets the style for the elevated button widget.
func (r *ElevatedButtonAttributes) SetStyle(style *duit_props.ButtonStyle) *ElevatedButtonAttributes {
	r.Style = style
	return r
}

// SetOnLongPress sets the on long press for the elevated button widget.
func (r *ElevatedButtonAttributes) SetOnLongPress(onLongPress any) *ElevatedButtonAttributes {
	r.OnLongPress = onLongPress
	return r
}
