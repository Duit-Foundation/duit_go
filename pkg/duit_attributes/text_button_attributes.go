package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// TextButtonAttributes defines attributes for TextButton widget.
// See: https://api.flutter.dev/flutter/material/TextButton-class.html
type TextButtonAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
	*ThemeConsumer        `gen:"wrap"`
	*FocusCapability      `gen:"wrap"`
	Autofocus      duit_utils.Tristate[bool] `json:"autofocus,omitempty"`
	ClipBehavior   duit_props.Clip           `json:"clipBehavior,omitempty"`
	Style          *duit_props.ButtonStyle   `json:"style,omitempty"`
	OnLongPress    any                       `json:"onLongPress,omitempty"`
	OnHover        any                       `json:"onHover,omitempty"`
	OnFocusChange  any                       `json:"onFocusChange,omitempty"`
}

//bindgen:exclude
func (r *TextButtonAttributes) Validate() error {
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

// NewTextButtonAttributes creates a new TextButtonAttributes instance.
func NewTextButtonAttributes() *TextButtonAttributes {
	return &TextButtonAttributes{}
}

// SetAutofocus sets the autofocus for the text button widget.
func (r *TextButtonAttributes) SetAutofocus(autofocus bool) *TextButtonAttributes {
	r.Autofocus = duit_utils.BoolValue(autofocus)
	return r
}

// SetClipBehavior sets the clip behavior for the text button widget.
func (r *TextButtonAttributes) SetClipBehavior(clipBehavior duit_props.Clip) *TextButtonAttributes {
	r.ClipBehavior = clipBehavior
	return r
}

// SetStyle sets the style for the text button widget.
func (r *TextButtonAttributes) SetStyle(style *duit_props.ButtonStyle) *TextButtonAttributes {
	r.Style = style
	return r
}

// SetOnLongPress sets the on long press for the text button widget.
func (r *TextButtonAttributes) SetOnLongPress(onLongPress any) *TextButtonAttributes {
	r.OnLongPress = onLongPress
	return r
}

// SetOnHover sets the on hover for the text button widget.
func (r *TextButtonAttributes) SetOnHover(onHover any) *TextButtonAttributes {
	r.OnHover = onHover
	return r
}

// SetOnFocusChange sets the on focus change for the text button widget.
func (r *TextButtonAttributes) SetOnFocusChange(onFocusChange any) *TextButtonAttributes {
	r.OnFocusChange = onFocusChange
	return r
}
