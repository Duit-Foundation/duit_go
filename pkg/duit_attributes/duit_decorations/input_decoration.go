package duit_decoration

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_edge_insets"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_text_properties"
)

type InputDecoration[T duit_edge_insets.EdgeInsets] struct {
	LabelText          string                         `json:"labelText,omitempty"`
	LabelStyle         duit_text_properties.TextStyle `json:"labelStyle,omitempty"`
	FloatingLabelStyle duit_text_properties.TextStyle `json:"floatingLabelStyle,omitempty"`
	HelperText         string                         `json:"helperText,omitempty"`
	HelperStyle        duit_text_properties.TextStyle `json:"helperStyle,omitempty"`
	HintText           string                         `json:"hintText,omitempty"`
	HintStyle          duit_text_properties.TextStyle `json:"hintStyle,omitempty"`
	HintMaxLines       int                            `json:"hintMaxLines,omitempty"`
	ErrorText          string                         `json:"errorText,omitempty"`
	ErrorStyle         duit_text_properties.TextStyle `json:"errorStyle,omitempty"`
	ErrorMaxLines      int                            `json:"errorMaxLines,omitempty"`
	Border             Border                         `json:"border,omitempty"`
	ErrorBorder        Border                         `json:"errorBorder,omitempty"`
	Enabledborder      Border                         `json:"enabledBorder,omitempty"`
	FocusedBorder      Border                         `json:"focusedBorder,omitempty"`
	FocusedErrorBorder Border                         `json:"focusedErrorBorder,omitempty"`
	Enabled            bool                           `json:"enabled,omitempty"`
	IsCollapsed        bool                           `json:"isCollapsed,omitempty"`
	IsDense            bool                           `json:"isDense,omitempty"`
	SuffixText         string                         `json:"suffixText,omitempty"`
	SuffixStyle        duit_text_properties.TextStyle `json:"suffixStyle,omitempty"`
	PrefixText         string                         `json:"prefixText,omitempty"`
	PrefixStyle        duit_text_properties.TextStyle `json:"prefixStyle,omitempty"`
	CounterText        string                         `json:"counterText,omitempty"`
	CounterStyle       duit_text_properties.TextStyle `json:"counterStyle,omitempty"`
	AlignLabelWithHint bool                           `json:"alignLabelWithHint,omitempty"`
	Filled             bool                           `json:"filled,omitempty"`
	FillColor          string                         `json:"fillColor,omitempty"`
	ContentPadding     T                              `json:"contentPadding,omitempty"`
}
