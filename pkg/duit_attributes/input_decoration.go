package duit_attributes

import "github.com/lesleysin/duit_go/pkg/duit_attributes/duit_border"

type InputDecoration[T EdgeInsets] struct {
	LabelText          string             `json:"labelText,omitempty"`
	LabelStyle         TextStyle          `json:"labelStyle,omitempty"`
	FloatingLabelStyle TextStyle          `json:"floatingLabelStyle,omitempty"`
	HelperText         string             `json:"helperText,omitempty"`
	HelperStyle        TextStyle          `json:"helperStyle,omitempty"`
	HintText           string             `json:"hintText,omitempty"`
	HintStyle          TextStyle          `json:"hintStyle,omitempty"`
	HintMaxLines       int                `json:"hintMaxLines,omitempty"`
	ErrorText          string             `json:"errorText,omitempty"`
	ErrorStyle         TextStyle          `json:"errorStyle,omitempty"`
	ErrorMaxLines      int                `json:"errorMaxLines,omitempty"`
	Border             duit_border.Border `json:"border,omitempty"`
	ErrorBorder        duit_border.Border `json:"errorBorder,omitempty"`
	Enabledborder      duit_border.Border `json:"enabledBorder,omitempty"`
	FocusedBorder      duit_border.Border `json:"focusedBorder,omitempty"`
	FocusedErrorBorder duit_border.Border `json:"focusedErrorBorder,omitempty"`
	Enabled            bool               `json:"enabled,omitempty"`
	IsCollapsed        bool               `json:"isCollapsed,omitempty"`
	IsDense            bool               `json:"isDense,omitempty"`
	SuffixText         string             `json:"suffixText,omitempty"`
	SuffixStyle        TextStyle          `json:"suffixStyle,omitempty"`
	PrefixText         string             `json:"prefixText,omitempty"`
	PrefixStyle        TextStyle          `json:"prefixStyle,omitempty"`
	CounterText        string             `json:"counterText,omitempty"`
	CounterStyle       TextStyle          `json:"counterStyle,omitempty"`
	AlignLabelWithHint bool               `json:"alignLabelWithHint,omitempty"`
	Filled             bool               `json:"filled,omitempty"`
	FillColor          string             `json:"fillColor,omitempty"`
	ContentPadding     T                  `json:"contentPadding,omitempty"`
}
