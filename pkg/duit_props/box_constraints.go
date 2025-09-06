package duit_props

import "errors"

// BoxConstraints defines size constraints for a widget in Flutter's layout system.
// Contains minimum and maximum values for width and height that a widget can have.
//
// Used in the layout building process to pass information about available space
// from parent widget to child widget.
//
// Flutter documentation: https://api.flutter.dev/flutter/rendering/BoxConstraints-class.html
type BoxConstraints struct {
	// MinWidth defines the minimum width of the widget in logical pixels
	MinWidth float32 `json:"minWidth,omitempty"`
	// MaxWidth defines the maximum width of the widget in logical pixels
	MaxWidth float32 `json:"maxWidth,omitempty"`
	// MinHeight defines the minimum height of the widget in logical pixels
	MinHeight float32 `json:"minHeight,omitempty"`
	// MaxHeight defines the maximum height of the widget in logical pixels
	MaxHeight float32 `json:"maxHeight,omitempty"`
}

// NewBoxConstraints creates a new empty BoxConstraints structure with default values
func NewBoxConstraints() *BoxConstraints {
	return &BoxConstraints{}
}

// SetMinWidth sets the minimum width of the widget
func (r *BoxConstraints) SetMinWidth(width float32) *BoxConstraints {
	r.MinWidth = width
	return r
}

// SetMaxWidth sets the maximum width of the widget
func (r *BoxConstraints) SetMaxWidth(width float32) *BoxConstraints {
	r.MaxWidth = width
	return r
}

// SetMinHeight sets the minimum height of the widget
func (r *BoxConstraints) SetMinHeight(height float32) *BoxConstraints {
	r.MinHeight = height
	return r
}

// SetMaxHeight sets the maximum height of the widget
func (r *BoxConstraints) SetMaxHeight(height float32) *BoxConstraints {
	r.MaxHeight = height
	return r
}

// Validate checks if the BoxConstraints are valid
//
// - Check width constraints: 0.0 <= minWidth <= maxWidth
//
// - Check height constraints: 0.0 <= minHeight <= maxHeight
func (r *BoxConstraints) Validate() error {
	// Check width constraints: 0.0 <= minWidth <= maxWidth
	if r.MinWidth < 0.0 {
		return errors.New("minWidth must be >= 0.0")
	}
	if r.MaxWidth < 0.0 {
		return errors.New("maxWidth must be >= 0.0")
	}
	if r.MinWidth > r.MaxWidth {
		return errors.New("minWidth must be <= maxWidth")
	}

	// Check height constraints: 0.0 <= minHeight <= maxHeight
	if r.MinHeight < 0.0 {
		return errors.New("minHeight must be >= 0.0")
	}
	if r.MaxHeight < 0.0 {
		return errors.New("maxHeight must be >= 0.0")
	}
	if r.MinHeight > r.MaxHeight {
		return errors.New("minHeight must be <= maxHeight")
	}

	return nil
}
