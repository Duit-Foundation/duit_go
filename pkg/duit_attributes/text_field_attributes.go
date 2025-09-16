package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// TextFieldAttributes represents attributes for TextField widget.
// https://api.flutter.dev/flutter/material/TextField-class.html
type TextFieldAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
	*ThemeConsumer        `gen:"wrap"`
	Style                 *duit_props.TextStyle       `json:"style,omitempty"`
	TextAlign             duit_props.TextAlign        `json:"textAlign,omitempty"`
	TextDirection         duit_props.TextDirection    `json:"textDirection,omitempty"`
	Autocorrect           duit_utils.Tristate[bool]   `json:"autocorrect,omitempty"`
	EnableSuggestions     duit_utils.Tristate[bool]   `json:"enableSuggestions,omitempty"`
	Expands               duit_utils.Tristate[bool]   `json:"expands,omitempty"`
	ReadOnly              duit_utils.Tristate[bool]   `json:"readOnly,omitempty"`
	ShowCursor            duit_utils.Tristate[bool]   `json:"showCursor,omitempty"`
	Enabled               duit_utils.Tristate[bool]   `json:"enabled,omitempty"`
	ObscureText           duit_utils.Tristate[bool]   `json:"obscureText,omitempty"`
	Autofocus             duit_utils.Tristate[bool]   `json:"autofocus,omitempty"`
	ObscuringCharacter    string                      `json:"obscuringCharacter,omitempty"`
	MaxLines              uint                        `json:"maxLines,omitempty"`
	MinLines              uint                        `json:"minLines,omitempty"`
	MaxLength             uint                        `json:"maxLength,omitempty"`
	Decoration            *duit_props.InputDecoration `json:"inputDecoration,omitempty"`
}

// NewTextFieldAttributes creates a new instance of TextFieldAttributes.
func NewTextFieldAttributes() *TextFieldAttributes {
	return &TextFieldAttributes{}
}

// SetStyle sets the style property and returns the TextFieldAttributes instance for method chaining.
func (r *TextFieldAttributes) SetStyle(style *duit_props.TextStyle) *TextFieldAttributes {
	r.Style = style
	return r
}

// SetTextAlign sets the textAlign property and returns the TextFieldAttributes instance for method chaining.
func (r *TextFieldAttributes) SetTextAlign(textAlign duit_props.TextAlign) *TextFieldAttributes {
	r.TextAlign = textAlign
	return r
}

// SetTextDirection sets the textDirection property and returns the TextFieldAttributes instance for method chaining.
func (r *TextFieldAttributes) SetTextDirection(textDirection duit_props.TextDirection) *TextFieldAttributes {
	r.TextDirection = textDirection
	return r
}

// SetAutocorrect sets the autocorrect property and returns the TextFieldAttributes instance for method chaining.
func (r *TextFieldAttributes) SetAutocorrect(autocorrect bool) *TextFieldAttributes {
	r.Autocorrect = duit_utils.BoolValue(autocorrect)
	return r
}

// SetEnableSuggestions sets the enableSuggestions property and returns the TextFieldAttributes instance for method chaining.
func (r *TextFieldAttributes) SetEnableSuggestions(enableSuggestions bool) *TextFieldAttributes {
	r.EnableSuggestions = duit_utils.BoolValue(enableSuggestions)
	return r
}

// SetExpands sets the expands property and returns the TextFieldAttributes instance for method chaining.
func (r *TextFieldAttributes) SetExpands(expands bool) *TextFieldAttributes {
	r.Expands = duit_utils.BoolValue(expands)
	return r
}

// SetReadOnly sets the readOnly property and returns the TextFieldAttributes instance for method chaining.
func (r *TextFieldAttributes) SetReadOnly(readOnly bool) *TextFieldAttributes {
	r.ReadOnly = duit_utils.BoolValue(readOnly)
	return r
}

// SetShowCursor sets the showCursor property and returns the TextFieldAttributes instance for method chaining.
func (r *TextFieldAttributes) SetShowCursor(showCursor bool) *TextFieldAttributes {
	r.ShowCursor = duit_utils.BoolValue(showCursor)
	return r
}

// SetEnabled sets the enabled property and returns the TextFieldAttributes instance for method chaining.
func (r *TextFieldAttributes) SetEnabled(enabled bool) *TextFieldAttributes {
	r.Enabled = duit_utils.BoolValue(enabled)
	return r
}

// SetObscureText sets the obscureText property and returns the TextFieldAttributes instance for method chaining.
func (r *TextFieldAttributes) SetObscureText(obscureText bool) *TextFieldAttributes {
	r.ObscureText = duit_utils.BoolValue(obscureText)
	return r
}

// SetAutofocus sets the autofocus property and returns the TextFieldAttributes instance for method chaining.
func (r *TextFieldAttributes) SetAutofocus(autofocus bool) *TextFieldAttributes {
	r.Autofocus = duit_utils.BoolValue(autofocus)
	return r
}

// SetObscuringCharacter sets the obscuringCharacter property and returns the TextFieldAttributes instance for method chaining.
func (r *TextFieldAttributes) SetObscuringCharacter(obscuringCharacter string) *TextFieldAttributes {
	r.ObscuringCharacter = obscuringCharacter
	return r
}

// SetMaxLines sets the maxLines property and returns the TextFieldAttributes instance for method chaining.
func (r *TextFieldAttributes) SetMaxLines(maxLines uint) *TextFieldAttributes {
	r.MaxLines = maxLines
	return r
}

// SetMinLines sets the minLines property and returns the TextFieldAttributes instance for method chaining.
func (r *TextFieldAttributes) SetMinLines(minLines uint) *TextFieldAttributes {
	r.MinLines = minLines
	return r
}

// SetMaxLength sets the maxLength property and returns the TextFieldAttributes instance for method chaining.
func (r *TextFieldAttributes) SetMaxLength(maxLength uint) *TextFieldAttributes {
	r.MaxLength = maxLength
	return r
}

// SetDecoration sets the decoration property and returns the TextFieldAttributes instance for method chaining.
func (r *TextFieldAttributes) SetDecoration(decoration *duit_props.InputDecoration) *TextFieldAttributes {
	r.Decoration = decoration
	return r
}

//bindgen:exclude
func (r *TextFieldAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	return nil
}
