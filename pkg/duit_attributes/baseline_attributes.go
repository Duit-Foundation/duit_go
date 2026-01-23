package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type BaselineAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	Baseline                          duit_utils.Tristate[float32]                 `json:"baseline,omitempty"`
	BaselineType                      duit_props.TextBaseline `json:"baselineType,omitempty"`
}

func NewBaselineAttributes() *BaselineAttributes {
	return &BaselineAttributes{}
}

func (r *BaselineAttributes) SetBaseline(baseline float32) *BaselineAttributes {
	r.Baseline = duit_utils.Float32Value(baseline)
	return r
}

func (r *BaselineAttributes) SetBaselineType(baselineType duit_props.TextBaseline) *BaselineAttributes {
	r.BaselineType = baselineType
	return r
}

func (r *BaselineAttributes) Validate() error {
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
