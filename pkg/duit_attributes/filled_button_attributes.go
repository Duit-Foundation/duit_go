package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// FilledButtonAttributes defines attributes for FilledButton widget.
// See: https://api.flutter.dev/flutter/material/FilledButton-class.html
type FilledButtonAttributes struct {
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
func (r *FilledButtonAttributes) Validate() error {
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

// NewFilledButtonAttributes creates a new FilledButtonAttributes instance.
func NewFilledButtonAttributes() *FilledButtonAttributes {
	return &FilledButtonAttributes{}
}

// SetAutofocus sets the autofocus for the filled button widget.
func (r *FilledButtonAttributes) SetAutofocus(autofocus duit_utils.Tristate[bool]) *FilledButtonAttributes {
	r.Autofocus = autofocus
	return r
}

// SetClipBehavior sets the clip behavior for the filled button widget.
func (r *FilledButtonAttributes) SetClipBehavior(clipBehavior duit_props.Clip) *FilledButtonAttributes {
	r.ClipBehavior = clipBehavior
	return r
}

// SetStyle sets the style for the filled button widget.
func (r *FilledButtonAttributes) SetStyle(style *duit_props.ButtonStyle) *FilledButtonAttributes {
	r.Style = style
	return r
}

// SetOnLongPress sets the on long press for the filled button widget.
func (r *FilledButtonAttributes) SetOnLongPress(onLongPress any) *FilledButtonAttributes {
	r.OnLongPress = onLongPress
	return r
}

// SetOnHover sets the on hover for the filled button widget.
func (r *FilledButtonAttributes) SetOnHover(onHover any) *FilledButtonAttributes {
	r.OnHover = onHover
	return r
}

// SetOnFocusChange sets the on focus change for the filled button widget.
func (r *FilledButtonAttributes) SetOnFocusChange(onFocusChange any) *FilledButtonAttributes {
	r.OnFocusChange = onFocusChange
	return r
}
