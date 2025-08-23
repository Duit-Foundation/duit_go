package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type SliverIgnorePointerAttributes struct {
	*ValueReferenceHolder
	*ThemeConsumer
	*SliverProps
	Ignoring duit_utils.Tristate[bool] `json:"ignoring"`
}

func (r *SliverIgnorePointerAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if r.Ignoring == nil {
		return errors.New("ignoring property is required")
	}

	return nil
}
