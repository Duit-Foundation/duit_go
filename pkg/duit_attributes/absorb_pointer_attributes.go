package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type AbsorbPointerAttributes struct {
	*ValueReferenceHolder
	Absorbing duit_utils.Tristate[bool] `json:"absorbing,omitempty"`
}

func (r *AbsorbPointerAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	return nil
}
