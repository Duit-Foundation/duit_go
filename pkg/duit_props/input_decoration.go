package duit_props

type InputDecoration struct {
	LabelText          string       `json:"labelText,omitempty"`
	LabelStyle         *TextStyle   `json:"labelStyle,omitempty"`
	FloatingLabelStyle *TextStyle   `json:"floatingLabelStyle,omitempty"`
	HelperText         string       `json:"helperText,omitempty"`
	HelperStyle        *TextStyle   `json:"helperStyle,omitempty"`
	HintText           string       `json:"hintText,omitempty"`
	HintStyle          *TextStyle   `json:"hintStyle,omitempty"`
	HintMaxLines       int          `json:"hintMaxLines,omitempty"`
	ErrorText          string       `json:"errorText,omitempty"`
	ErrorStyle         *TextStyle   `json:"errorStyle,omitempty"`
	ErrorMaxLines      int          `json:"errorMaxLines,omitempty"`
	Border             *InputBorder `json:"border,omitempty"`
	ErrorBorder        *InputBorder `json:"errorBorder,omitempty"`
	Enabledborder      *InputBorder `json:"enabledBorder,omitempty"`
	FocusedBorder      *InputBorder `json:"focusedBorder,omitempty"`
	FocusedErrorBorder *InputBorder `json:"focusedErrorBorder,omitempty"`
	Enabled            bool         `json:"enabled,omitempty"`
	IsCollapsed        bool         `json:"isCollapsed,omitempty"`
	IsDense            bool         `json:"isDense,omitempty"`
	SuffixText         string       `json:"suffixText,omitempty"`
	SuffixStyle        *TextStyle   `json:"suffixStyle,omitempty"`
	PrefixText         string       `json:"prefixText,omitempty"`
	PrefixStyle        *TextStyle   `json:"prefixStyle,omitempty"`
	CounterText        string       `json:"counterText,omitempty"`
	CounterStyle       *TextStyle   `json:"counterStyle,omitempty"`
	AlignLabelWithHint bool         `json:"alignLabelWithHint,omitempty"`
	Filled             bool         `json:"filled,omitempty"`
	FillColor          string       `json:"fillColor,omitempty"`
	ContentPadding     *EdgeInsets  `json:"contentPadding,omitempty"`
}
