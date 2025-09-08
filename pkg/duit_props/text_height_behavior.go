package duit_props

import "github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"

// TextHeightBehavior represents the height behavior of text.
// It is used to control the height of text in a widget.
//
// This is a Go implementation of Flutter's TextHeightBehavior class.
// For more information, see: https://api.flutter.dev/flutter/painting/TextHeightBehavior-class.html
type TextHeightBehavior struct {
	ApplyHeightToFirstAscent duit_utils.Tristate[bool]                    `json:"applyHeightToFirstAscent,omitempty"`
	ApplyHeightToLastDescent duit_utils.Tristate[bool]                    `json:"applyHeightToLastDescent,omitempty"`
	LeadingDistribution      TextLeadingDistribution `json:"leadingDistribution,omitempty"`
}

// NewTextHeightBehavior creates a new instance of TextHeightBehavior
func NewTextHeightBehavior() *TextHeightBehavior {
	return &TextHeightBehavior{}
}

// SetApplyHeightToFirstAscent sets the apply height to first ascent
func (r *TextHeightBehavior) SetApplyHeightToFirstAscent(value bool) *TextHeightBehavior {
	r.ApplyHeightToFirstAscent = duit_utils.BoolValue(value)
	return r
}

// SetApplyHeightToLastDescent sets the apply height to last descent
func (r *TextHeightBehavior) SetApplyHeightToLastDescent(value bool) *TextHeightBehavior {
	r.ApplyHeightToLastDescent = duit_utils.BoolValue(value)
	return r
}

// SetLeadingDistribution sets the leading distribution
func (r *TextHeightBehavior) SetLeadingDistribution(value TextLeadingDistribution) *TextHeightBehavior {
	r.LeadingDistribution = value
	return r
}