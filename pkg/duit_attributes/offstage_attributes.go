package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type OffstageAttributes struct {
	*ValueReferenceHolder
	Offstage	duit_utils.Tristate[bool]	`json:"offstage,omitempty"`
}

func (r *OffstageAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if r.Offstage == nil {
		return errors.New("offstage property is required")
	}

	return nil
}