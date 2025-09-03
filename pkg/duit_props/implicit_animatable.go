package duit_props

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
)

type ImplicitAnimatable struct {
	Duration uint   `json:"duration"`
	Curve    Curves `json:"curve,omitempty"`
	OnEnd    any    `json:"onEnd,omitempty"`
}

func (r *ImplicitAnimatable) Validate() error {
	if r == nil {
		return nil
	}

	if r.Duration == 0 {
		return errors.New("duration property is required and must be greater than 0")
	}

	if err := duit_action.CheckActionType(r.OnEnd); err != nil {
		return err
	}

	return nil
}
