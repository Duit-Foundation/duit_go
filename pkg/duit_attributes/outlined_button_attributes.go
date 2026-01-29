package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// OutlinedButtonAttributes defines attributes for OutlinedButton widget.
// See: https://api.flutter.dev/flutter/material/OutlinedButton-class.html
type OutlinedButtonAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
	*ThemeConsumer        `gen:"wrap"`
	*FocusCapability      `gen:"wrap"`
	Autofocus             duit_utils.Tristate[bool] `json:"autofocus,omitempty"`
	ClipBehavior          duit_props.Clip           `json:"clipBehavior,omitempty"`
	Style                 *duit_props.ButtonStyle   `json:"style,omitempty"`
	OnLongPress           any                       `json:"onLongPress,omitempty"`
	OnHover               any                       `json:"onHover,omitempty"`
	OnFocusChange         any                       `json:"onFocusChange,omitempty"`
}

//bindgen:exclude
func (r *OutlinedButtonAttributes) Validate() error {
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

	if r.OnHover != nil {
		if err := duit_action.CheckActionType(r.OnHover); err != nil {
			return err
		}
	}

	if r.OnFocusChange != nil {
		if err := duit_action.CheckActionType(r.OnFocusChange); err != nil {
			return err
		}
	}

	return nil
}

// NewOutlinedButtonAttributes creates a new OutlinedButtonAttributes instance.
func NewOutlinedButtonAttributes() *OutlinedButtonAttributes {
	return &OutlinedButtonAttributes{}
}

// SetAutofocus sets the autofocus for the outlined button widget.
func (r *OutlinedButtonAttributes) SetAutofocus(autofocus duit_utils.Tristate[bool]) *OutlinedButtonAttributes {
	r.Autofocus = autofocus
	return r
}

// SetClipBehavior sets the clip behavior for the outlined button widget.
func (r *OutlinedButtonAttributes) SetClipBehavior(clipBehavior duit_props.Clip) *OutlinedButtonAttributes {
	r.ClipBehavior = clipBehavior
	return r
}

// SetStyle sets the style for the outlined button widget.
func (r *OutlinedButtonAttributes) SetStyle(style *duit_props.ButtonStyle) *OutlinedButtonAttributes {
	r.Style = style
	return r
}

// SetOnLongPress sets the on long press for the outlined button widget.
func (r *OutlinedButtonAttributes) SetOnLongPress(onLongPress any) *OutlinedButtonAttributes {
	r.OnLongPress = onLongPress
	return r
}

// SetOnHover sets the on hover for the outlined button widget.
func (r *OutlinedButtonAttributes) SetOnHover(onHover any) *OutlinedButtonAttributes {
	r.OnHover = onHover
	return r
}

// SetOnFocusChange sets the on focus change for the outlined button widget.
func (r *OutlinedButtonAttributes) SetOnFocusChange(onFocusChange any) *OutlinedButtonAttributes {
	r.OnFocusChange = onFocusChange
	return r
}
