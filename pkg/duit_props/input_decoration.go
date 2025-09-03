package duit_props

type InputDecoration[TInsets EdgeInsets, TColor Color] struct {
	LabelText          string               `json:"labelText,omitempty"`
	LabelStyle         *TextStyle[TColor]   `json:"labelStyle,omitempty"`
	FloatingLabelStyle *TextStyle[TColor]   `json:"floatingLabelStyle,omitempty"`
	HelperText         string               `json:"helperText,omitempty"`
	HelperStyle        *TextStyle[TColor]   `json:"helperStyle,omitempty"`
	HintText           string               `json:"hintText,omitempty"`
	HintStyle          *TextStyle[TColor]   `json:"hintStyle,omitempty"`
	HintMaxLines       int                  `json:"hintMaxLines,omitempty"`
	ErrorText          string               `json:"errorText,omitempty"`
	ErrorStyle         *TextStyle[TColor]   `json:"errorStyle,omitempty"`
	ErrorMaxLines      int                  `json:"errorMaxLines,omitempty"`
	Border             *InputBorder[TColor] `json:"border,omitempty"`
	ErrorBorder        *InputBorder[TColor] `json:"errorBorder,omitempty"`
	Enabledborder      *InputBorder[TColor] `json:"enabledBorder,omitempty"`
	FocusedBorder      *InputBorder[TColor] `json:"focusedBorder,omitempty"`
	FocusedErrorBorder *InputBorder[TColor] `json:"focusedErrorBorder,omitempty"`
	Enabled            bool                 `json:"enabled,omitempty"`
	IsCollapsed        bool                 `json:"isCollapsed,omitempty"`
	IsDense            bool                 `json:"isDense,omitempty"`
	SuffixText         string               `json:"suffixText,omitempty"`
	SuffixStyle        *TextStyle[TColor]   `json:"suffixStyle,omitempty"`
	PrefixText         string               `json:"prefixText,omitempty"`
	PrefixStyle        *TextStyle[TColor]   `json:"prefixStyle,omitempty"`
	CounterText        string               `json:"counterText,omitempty"`
	CounterStyle       *TextStyle[TColor]   `json:"counterStyle,omitempty"`
	AlignLabelWithHint bool                 `json:"alignLabelWithHint,omitempty"`
	Filled             bool                 `json:"filled,omitempty"`
	FillColor          string               `json:"fillColor,omitempty"`
	ContentPadding     TInsets              `json:"contentPadding,omitempty"`
}
