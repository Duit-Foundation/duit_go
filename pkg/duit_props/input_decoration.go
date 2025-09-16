package duit_props

import "github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"

// InputDecoration represents the decoration configuration for input fields.
// It provides visual styling options including labels, hints, borders, and error states.
//
// This is a Go implementation of Flutter's InputDecoration class.
// For more information, see: https://api.flutter.dev/flutter/material/InputDecoration-class.html
type InputDecoration struct {
	LabelText          string                    `json:"labelText,omitempty"`
	LabelStyle         *TextStyle                `json:"labelStyle,omitempty"`
	FloatingLabelStyle *TextStyle                `json:"floatingLabelStyle,omitempty"`
	HelperText         string                    `json:"helperText,omitempty"`
	HelperStyle        *TextStyle                `json:"helperStyle,omitempty"`
	HintText           string                    `json:"hintText,omitempty"`
	HintStyle          *TextStyle                `json:"hintStyle,omitempty"`
	HintMaxLines       int                       `json:"hintMaxLines,omitempty"`
	ErrorText          string                    `json:"errorText,omitempty"`
	ErrorStyle         *TextStyle                `json:"errorStyle,omitempty"`
	ErrorMaxLines      int                       `json:"errorMaxLines,omitempty"`
	Border             *InputBorder              `json:"inputBorder,omitempty"`
	ErrorBorder        *InputBorder              `json:"errorBorder,omitempty"`
	Enabledborder      *InputBorder              `json:"enabledBorder,omitempty"`
	FocusedBorder      *InputBorder              `json:"focusedBorder,omitempty"`
	FocusedErrorBorder *InputBorder              `json:"focusedErrorBorder,omitempty"`
	Enabled            duit_utils.Tristate[bool] `json:"enabled,omitempty"`
	IsCollapsed        duit_utils.Tristate[bool] `json:"isCollapsed,omitempty"`
	IsDense            duit_utils.Tristate[bool] `json:"isDense,omitempty"`
	SuffixText         string                    `json:"suffixText,omitempty"`
	SuffixStyle        *TextStyle                `json:"suffixStyle,omitempty"`
	PrefixText         string                    `json:"prefixText,omitempty"`
	PrefixStyle        *TextStyle                `json:"prefixStyle,omitempty"`
	CounterText        string                    `json:"counterText,omitempty"`
	CounterStyle       *TextStyle                `json:"counterStyle,omitempty"`
	AlignLabelWithHint duit_utils.Tristate[bool] `json:"alignLabelWithHint,omitempty"`
	Filled             duit_utils.Tristate[bool] `json:"filled,omitempty"`
	FillColor          string                    `json:"fillColor,omitempty"`
	ContentPadding     *EdgeInsets               `json:"contentPadding,omitempty"`
}

// NewInputDecoration creates a new instance of InputDecoration
func NewInputDecoration() *InputDecoration {
	return &InputDecoration{}
}

// SetLabelText sets the text that describes the input field
func (r *InputDecoration) SetLabelText(labelText string) *InputDecoration {
	r.LabelText = labelText
	return r
}

// SetLabelStyle sets the style of the label text
func (r *InputDecoration) SetLabelStyle(labelStyle *TextStyle) *InputDecoration {
	r.LabelStyle = labelStyle
	return r
}

// SetFloatingLabelStyle sets the style of the floating label
func (r *InputDecoration) SetFloatingLabelStyle(floatingLabelStyle *TextStyle) *InputDecoration {
	r.FloatingLabelStyle = floatingLabelStyle
	return r
}

// SetHelperText sets the helper text displayed below the input field
func (r *InputDecoration) SetHelperText(helperText string) *InputDecoration {
	r.HelperText = helperText
	return r
}

// SetHelperStyle sets the style of the helper text
func (r *InputDecoration) SetHelperStyle(helperStyle *TextStyle) *InputDecoration {
	r.HelperStyle = helperStyle
	return r
}

// SetHintText sets the hint text displayed when the input is empty
func (r *InputDecoration) SetHintText(hintText string) *InputDecoration {
	r.HintText = hintText
	return r
}

// SetHintStyle sets the style of the hint text
func (r *InputDecoration) SetHintStyle(hintStyle *TextStyle) *InputDecoration {
	r.HintStyle = hintStyle
	return r
}

// SetHintMaxLines sets the maximum number of lines for hint text
func (r *InputDecoration) SetHintMaxLines(hintMaxLines int) *InputDecoration {
	r.HintMaxLines = hintMaxLines
	return r
}

// SetErrorText sets the error text displayed below the input field
func (r *InputDecoration) SetErrorText(errorText string) *InputDecoration {
	r.ErrorText = errorText
	return r
}

// SetErrorStyle sets the style of the error text
func (r *InputDecoration) SetErrorStyle(errorStyle *TextStyle) *InputDecoration {
	r.ErrorStyle = errorStyle
	return r
}

// SetErrorMaxLines sets the maximum number of lines for error text
func (r *InputDecoration) SetErrorMaxLines(errorMaxLines int) *InputDecoration {
	r.ErrorMaxLines = errorMaxLines
	return r
}

// SetBorder sets the default border for the input field
func (r *InputDecoration) SetBorder(border *InputBorder) *InputDecoration {
	r.Border = border
	return r
}

// SetErrorBorder sets the border when the input field has an error
func (r *InputDecoration) SetErrorBorder(errorBorder *InputBorder) *InputDecoration {
	r.ErrorBorder = errorBorder
	return r
}

// SetEnabledBorder sets the border when the input field is enabled
func (r *InputDecoration) SetEnabledBorder(enabledBorder *InputBorder) *InputDecoration {
	r.Enabledborder = enabledBorder
	return r
}

// SetFocusedBorder sets the border when the input field is focused
func (r *InputDecoration) SetFocusedBorder(focusedBorder *InputBorder) *InputDecoration {
	r.FocusedBorder = focusedBorder
	return r
}

// SetFocusedErrorBorder sets the border when the input field is focused and has an error
func (r *InputDecoration) SetFocusedErrorBorder(focusedErrorBorder *InputBorder) *InputDecoration {
	r.FocusedErrorBorder = focusedErrorBorder
	return r
}

// SetEnabled sets whether the input field is enabled
func (r *InputDecoration) SetEnabled(enabled bool) *InputDecoration {
	r.Enabled = duit_utils.BoolValue(enabled)
	return r
}

// SetIsCollapsed sets whether the input field decoration is collapsed
func (r *InputDecoration) SetIsCollapsed(isCollapsed bool) *InputDecoration {
	r.IsCollapsed = duit_utils.BoolValue(isCollapsed)
	return r
}

// SetIsDense sets whether the input field has dense layout
func (r *InputDecoration) SetIsDense(isDense bool) *InputDecoration {
	r.IsDense = duit_utils.BoolValue(isDense)
	return r
}

// SetSuffixText sets the text displayed after the input
func (r *InputDecoration) SetSuffixText(suffixText string) *InputDecoration {
	r.SuffixText = suffixText
	return r
}

// SetSuffixStyle sets the style of the suffix text
func (r *InputDecoration) SetSuffixStyle(suffixStyle *TextStyle) *InputDecoration {
	r.SuffixStyle = suffixStyle
	return r
}

// SetPrefixText sets the text displayed before the input
func (r *InputDecoration) SetPrefixText(prefixText string) *InputDecoration {
	r.PrefixText = prefixText
	return r
}

// SetPrefixStyle sets the style of the prefix text
func (r *InputDecoration) SetPrefixStyle(prefixStyle *TextStyle) *InputDecoration {
	r.PrefixStyle = prefixStyle
	return r
}

// SetCounterText sets the counter text displayed below the input field
func (r *InputDecoration) SetCounterText(counterText string) *InputDecoration {
	r.CounterText = counterText
	return r
}

// SetCounterStyle sets the style of the counter text
func (r *InputDecoration) SetCounterStyle(counterStyle *TextStyle) *InputDecoration {
	r.CounterStyle = counterStyle
	return r
}

// SetAlignLabelWithHint sets whether the label should align with hint text
func (r *InputDecoration) SetAlignLabelWithHint(alignLabelWithHint bool) *InputDecoration {
	r.AlignLabelWithHint = duit_utils.BoolValue(alignLabelWithHint)
	return r
}

// SetFilled sets whether the input field should be filled
func (r *InputDecoration) SetFilled(filled bool) *InputDecoration {
	r.Filled = duit_utils.BoolValue(filled)
	return r
}

// SetFillColor sets the fill color of the input field
func (r *InputDecoration) SetFillColor(fillColor string) *InputDecoration {
	r.FillColor = fillColor
	return r
}

// SetContentPadding sets the padding inside the input field
func (r *InputDecoration) SetContentPadding(contentPadding *EdgeInsets) *InputDecoration {
	r.ContentPadding = contentPadding
	return r
}

func (r *InputDecoration) Validate() error {
	if r.LabelStyle != nil {
		if err := r.LabelStyle.Validate(); err != nil {
			return err
		}
	}

	if r.FloatingLabelStyle != nil {
		if err := r.FloatingLabelStyle.Validate(); err != nil {
			return err
		}
	}

	if r.HelperStyle != nil {
		if err := r.HelperStyle.Validate(); err != nil {
			return err
		}
	}

	if r.HintStyle != nil {
		if err := r.HintStyle.Validate(); err != nil {
			return err
		}
	}

	if r.ErrorStyle != nil {
		if err := r.ErrorStyle.Validate(); err != nil {
			return err
		}
	}

	if r.ErrorBorder != nil {
		if err := r.ErrorBorder.Validate(); err != nil {
			return err
		}
	}

	if r.Enabledborder != nil {
		if err := r.Enabledborder.Validate(); err != nil {
			return err
		}
	}

	if r.FocusedBorder != nil {
		if err := r.FocusedBorder.Validate(); err != nil {
			return err
		}
	}

	if r.FocusedErrorBorder != nil {
		if err := r.FocusedErrorBorder.Validate(); err != nil {
			return err
		}
	}

	if r.SuffixStyle != nil {
		if err := r.SuffixStyle.Validate(); err != nil {
			return err
		}
	}

	if r.PrefixStyle != nil {
		if err := r.PrefixStyle.Validate(); err != nil {
			return err
		}
	}

	if r.CounterStyle != nil {
		if err := r.CounterStyle.Validate(); err != nil {
			return err
		}
	}

	if r.ContentPadding != nil {
		if err := r.ContentPadding.Validate(); err != nil {
			return err
		}
	}

	if r.Border != nil {
		if err := r.Border.Validate(); err != nil {
			return err
		}
	}

	if r.ErrorBorder != nil {
		if err := r.ErrorBorder.Validate(); err != nil {
			return err
		}
	}

	if r.Enabledborder != nil {
		if err := r.Enabledborder.Validate(); err != nil {
			return err
		}
	}

	if r.FocusedBorder != nil {
		if err := r.FocusedBorder.Validate(); err != nil {
			return err
		}
	}

	if r.FocusedErrorBorder != nil {
		if err := r.FocusedErrorBorder.Validate(); err != nil {
			return err
		}
	}

	return nil
}
