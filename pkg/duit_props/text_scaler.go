package duit_props

import "errors"

// TextScaler represents the text scaler for a widget.
// It is used to scale the text in a widget.
//
// This is a Go implementation of Flutter's TextScaler class.
// For more information, see: https://api.flutter.dev/flutter/painting/TextScaler-class.html
type TextScaler struct {
	TextScaleFactor float32 `json:"textScaleFactor,omitempty"`
}

// NewTextScaler creates a new instance of TextScaler
func NewTextScaler(factor float32) *TextScaler {
	return &TextScaler{
		TextScaleFactor: factor,
	}
}

func (r *TextScaler) Validate() error {
	if r.TextScaleFactor <= 0 {
		return errors.New("textScaleFactor must be greater than 0 and positive")
	}
	return nil
}
