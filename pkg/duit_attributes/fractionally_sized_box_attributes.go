package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type FractionallySizedBoxAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	WidthFactor                       duit_utils.Tristate[float32] `json:"widthFactor,omitempty"`
	HeightFactor                      duit_utils.Tristate[float32] `json:"heightFactor,omitempty"`
	Alignment                         duit_props.Alignment         `json:"alignment,omitempty"`
}

func NewFractionallySizedBoxAttributes() *FractionallySizedBoxAttributes {
	return &FractionallySizedBoxAttributes{}
}

func (r *FractionallySizedBoxAttributes) SetWidthFactor(widthFactor float32) *FractionallySizedBoxAttributes {
	r.WidthFactor = duit_utils.Float32Value(widthFactor)
	return r
}

func (r *FractionallySizedBoxAttributes) SetHeightFactor(heightFactor float32) *FractionallySizedBoxAttributes {
	r.HeightFactor = duit_utils.Float32Value(heightFactor)
	return r
}

func (r *FractionallySizedBoxAttributes) SetAlignment(alignment duit_props.Alignment) *FractionallySizedBoxAttributes {
	r.Alignment = alignment
	return r
}

func (r *FractionallySizedBoxAttributes) Validate() error {
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
