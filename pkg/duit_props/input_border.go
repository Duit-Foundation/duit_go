package duit_props

import "errors"

// InputBorderOptions represents options for input border
type InputBorderOptions struct {
	GapPadding   float32     `json:"gapPadding,omitempty"`
	BorderRadius float32     `json:"borderRadius,omitempty"`
	BorderSide   *BorderSide `json:"borderSide,omitempty"`
}

func (r *InputBorderOptions) Validate() error {
	if r.BorderSide != nil {
		if err := r.BorderSide.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// InputBorder represents the appearance of a border for InputDecorator.
// The border is drawn relative to the input decorator's "container" which
// is the filled area above the helper text, error text, and counter.
// Borders can be decorated with a line whose color and thickness are
// defined by the borderSide property. InputDecorator animates the appearance
// of the border in response to state changes, such as gaining or losing focus.
//
// This is a Go implementation of Flutter's InputBorder class.
// For more information, see: https://api.flutter.dev/flutter/material/InputBorder-class.html
type InputBorder struct {
	Type    Border              `json:"type"`
	Options *InputBorderOptions `json:"options,omitempty"`
}

// NewInputBorder creates a new input border
func NewInputBorder(borderType Border) *InputBorder {
	return &InputBorder{
		Type:    borderType,
		Options: &InputBorderOptions{},
	}
}

// SetGapPadding sets the gap padding of the input border
func (r *InputBorder) SetGapPadding(gapPadding float32) *InputBorder {
	r.Options.GapPadding = gapPadding
	return r
}

// SetBorderRadius sets the border radius of the input border
func (r *InputBorder) SetBorderRadius(borderRadius float32) *InputBorder {
	r.Options.BorderRadius = borderRadius
	return r
}

// SetBorderSide sets the border side of the input border
func (r *InputBorder) SetBorderSide(borderSide *BorderSide) *InputBorder {
	r.Options.BorderSide = borderSide
	return r
}

func (r *InputBorder) Validate() error {
	if r.Options == nil {
		return errors.New("options are required")
	}

	if err := r.Options.Validate(); err != nil {
		return err
	}

	return nil
}
