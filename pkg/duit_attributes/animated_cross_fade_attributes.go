package duit_attributes

import (
	"errors"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type AnimatedCrossFadeAttributes struct {
	*ValueReferenceHolder
	*duit_props.ImplicitAnimatable
	*ThemeConsumer
	ReverseDuration    uint                      `json:"reverseDuration,omitempty"`
	FirstCurve         duit_props.Curves         `json:"firstCurve,omitempty"`
	SecondCurve        duit_props.Curves         `json:"secondCurve,omitempty"`
	SizeCurve          duit_props.Curves         `json:"sizeCurve,omitempty"`
	ExcludeBottomFocus duit_utils.Tristate[bool] `json:"excludeBottomFocus,omitempty"`
	Alignment          duit_props.Alignment  `json:"alignment,omitempty"`
	// Where 0 - CrossFadeState.showFirst and 1 - CrossFadeState.showSecond
	CrossFadeState uint8 `json:"crossFadeState"`
}

func (r *AnimatedCrossFadeAttributes) Validate() error {
	if r.ImplicitAnimatable != nil {
		if err := r.ImplicitAnimatable.Validate(); err != nil {
			return err
		}
	} else {
		return errors.New("implicitAnimatable property is required on implicit animated widgets")
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if r.CrossFadeState != 0 && r.CrossFadeState != 1 {
		return errors.New("crossFadeState must be 0 or 1")
	}

	return nil
}
