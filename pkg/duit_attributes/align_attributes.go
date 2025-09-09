package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// https://api.flutter.dev/flutter/widgets/Align-class.html
type AlignAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	Alignment                         duit_props.Alignment `json:"alignment,omitempty"`
	WidthFactor                       float32              `json:"widthFactor,omitempty"`
	HeightFactor                      float32              `json:"heightFactor,omitempty"`
}

// Create a new AlignAttributes
func NewAlignAttributes() *AlignAttributes {
	return &AlignAttributes{}
}

// Set the alignment value
func (r *AlignAttributes) SetAlignment(alignment duit_props.Alignment) *AlignAttributes {
	r.Alignment = alignment
	return r
}

// Set the width factor value
func (r *AlignAttributes) SetWidthFactor(widthFactor float32) *AlignAttributes {
	r.WidthFactor = widthFactor
	return r
}

// Set the height factor value
func (r *AlignAttributes) SetHeightFactor(heightFactor float32) *AlignAttributes {
	r.HeightFactor = heightFactor
	return r
}

// Validate the AlignAttributes
func (r *AlignAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if len(r.Alignment) == 0 {
		return errors.New("alignment property is required")
	}

	return nil
}
