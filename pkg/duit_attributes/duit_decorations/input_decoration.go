package duit_decoration

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_color"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_edge_insets"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_text_properties"
)

type InputDecoration[TInsets duit_edge_insets.EdgeInsets, TColor duit_color.Color] struct {
	LabelText          string                                 `json:"labelText,omitempty"`
	LabelStyle         duit_text_properties.TextStyle[TColor] `json:"labelStyle,omitempty"`
	FloatingLabelStyle duit_text_properties.TextStyle[TColor] `json:"floatingLabelStyle,omitempty"`
	HelperText         string                                 `json:"helperText,omitempty"`
	HelperStyle        duit_text_properties.TextStyle[TColor] `json:"helperStyle,omitempty"`
	HintText           string                                 `json:"hintText,omitempty"`
	HintStyle          duit_text_properties.TextStyle[TColor] `json:"hintStyle,omitempty"`
	HintMaxLines       int                                    `json:"hintMaxLines,omitempty"`
	ErrorText          string                                 `json:"errorText,omitempty"`
	ErrorStyle         duit_text_properties.TextStyle[TColor] `json:"errorStyle,omitempty"`
	ErrorMaxLines      int                                    `json:"errorMaxLines,omitempty"`
	Border             Border                                 `json:"border,omitempty"`
	ErrorBorder        Border                                 `json:"errorBorder,omitempty"`
	Enabledborder      Border                                 `json:"enabledBorder,omitempty"`
	FocusedBorder      Border                                 `json:"focusedBorder,omitempty"`
	FocusedErrorBorder Border                                 `json:"focusedErrorBorder,omitempty"`
	Enabled            bool                                   `json:"enabled,omitempty"`
	IsCollapsed        bool                                   `json:"isCollapsed,omitempty"`
	IsDense            bool                                   `json:"isDense,omitempty"`
	SuffixText         string                                 `json:"suffixText,omitempty"`
	SuffixStyle        duit_text_properties.TextStyle[TColor] `json:"suffixStyle,omitempty"`
	PrefixText         string                                 `json:"prefixText,omitempty"`
	PrefixStyle        duit_text_properties.TextStyle[TColor] `json:"prefixStyle,omitempty"`
	CounterText        string                                 `json:"counterText,omitempty"`
	CounterStyle       duit_text_properties.TextStyle[TColor] `json:"counterStyle,omitempty"`
	AlignLabelWithHint bool                                   `json:"alignLabelWithHint,omitempty"`
	Filled             bool                                   `json:"filled,omitempty"`
	FillColor          string                                 `json:"fillColor,omitempty"`
	ContentPadding     TInsets                                `json:"contentPadding,omitempty"`
}
