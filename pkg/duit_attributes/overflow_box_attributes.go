package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_alignment"
	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_flex"
)

type OverflowBoxAttributes struct {
	*ValueReferenceHolder
	*animations.AnimatedPropertyOwner
	*ThemeConsumer
	MinWidth  float32                  `json:"minWidth,omitempty"`
	MaxWidth  float32                  `json:"maxWidth,omitempty"`
	MinHeight float32                  `json:"minHeight,omitempty"`
	MaxHeight float32                  `json:"maxHeight,omitempty"`
	Alignment duit_alignment.Alignment `json:"alignment,omitempty"`
	Fit       duit_flex.OverflowBoxFit `json:"fit,omitempty"`
}

func (r *OverflowBoxAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	return nil
}