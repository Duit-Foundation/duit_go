package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type LimitedBoxAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	MaxWidth                          duit_utils.Tristate[float32] `json:"maxWidth,omitempty"`
	MaxHeight                         duit_utils.Tristate[float32] `json:"maxHeight,omitempty"`
}

func NewLimitedBoxAttributes() *LimitedBoxAttributes {
	return &LimitedBoxAttributes{}
}

func (r *LimitedBoxAttributes) SetMaxWidth(maxWidth float32) *LimitedBoxAttributes {
	r.MaxWidth = duit_utils.Float32Value(maxWidth)
	return r
}

func (r *LimitedBoxAttributes) SetMaxHeight(maxHeight float32) *LimitedBoxAttributes {
	r.MaxHeight = duit_utils.Float32Value(maxHeight)
	return r
}

func (r *LimitedBoxAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	return nil
}