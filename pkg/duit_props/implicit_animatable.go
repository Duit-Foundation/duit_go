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

func (r *ImplicitAnimatable) SetDuration(duration uint) *ImplicitAnimatable {
	r.Duration = duration
	return r
}

func (r *ImplicitAnimatable) SetCurve(curve Curves) *ImplicitAnimatable {
	r.Curve = curve
	return r
}

func (r *ImplicitAnimatable) SetOnEnd(onEnd any) *ImplicitAnimatable {
	r.OnEnd = onEnd
	return r
}

//bindgen:exclude
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
