package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// BackdropFilterAttributes defines attributes for BackdropFilter widget.
// See: https://api.flutter.dev/flutter/widgets/BackdropFilter-class.html
type BackdropFilterAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	Filter                            *duit_props.ImageFilter `json:"filter,omitempty"`
	BlendMode                         duit_props.BlendMode    `json:"blendMode,omitempty"`
}

//bindgen:exclude
func (r *BackdropFilterAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if r.Filter == nil {
		return errors.New("filter property is required")
	}

	if r.BlendMode == "" {
		return errors.New("blendMode property is required")
	}

	return nil
}

// NewBackdropFilterAttributes creates a new BackdropFilterAttributes instance.
func NewBackdropFilterAttributes() *BackdropFilterAttributes {
	return &BackdropFilterAttributes{}
}

// SetFilter sets the image filter for the backdrop filter widget.
func (r *BackdropFilterAttributes) SetFilter(filter *duit_props.ImageFilter) *BackdropFilterAttributes {
	r.Filter = filter
	return r
}

// SetBlendMode sets the blend mode for the backdrop filter widget.
func (r *BackdropFilterAttributes) SetBlendMode(blendMode duit_props.BlendMode) *BackdropFilterAttributes {
	r.BlendMode = blendMode
	return r
}
