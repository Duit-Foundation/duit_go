package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type IgnorePointerAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
	*ThemeConsumer        `gen:"wrap"`
	Ignoring              duit_utils.Tristate[bool] `json:"ignoring,omitempty"`
}

func NewIgnorePointerAttributes() *IgnorePointerAttributes {
	return &IgnorePointerAttributes{}
}

func (r *IgnorePointerAttributes) Validate() error {
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
