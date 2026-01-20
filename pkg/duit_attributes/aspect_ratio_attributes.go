package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type AspectRatioAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	AspectRatio                       duit_utils.Tristate[float32] `json:"aspectRatio,omitempty"`
}

func NewAspectRatioAttributes() *AspectRatioAttributes {
	return &AspectRatioAttributes{}
}

func (r *AspectRatioAttributes) SetAspectRatio(aspectRatio float32) *AspectRatioAttributes {
	r.AspectRatio = duit_utils.Float32Value(aspectRatio)
	return r
}

func (r *AspectRatioAttributes) Validate() error {
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