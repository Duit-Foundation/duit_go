package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_alignment"
	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type AnimatedCrossFadeAttributes struct {
	*ValueReferenceHolder
	*animations.ImplicitAnimatable
	*ThemeConsumer
	ReverseDuration    uint                      `json:"reverseDuration,omitempty"`
	FirstCurve         animations.Curves         `json:"firstCurve,omitempty"`
	SecondCurve        animations.Curves         `json:"secondCurve,omitempty"`
	SizeCurve          animations.Curves         `json:"sizeCurve,omitempty"`
	ExcludeBottomFocus duit_utils.Tristate[bool] `json:"excludeBottomFocus,omitempty"`
	Alignment          duit_alignment.Alignment  `json:"alignment,omitempty"`
	// Where 0 - CrossFadeState.showFirst and 1 - CrossFadeState.showSecond
	CrossFadeState uint8 `json:"crossFadeState"`
}

func (a *AnimatedCrossFadeAttributes) Validate() error {
	if err := a.ImplicitAnimatable.Validate(); err != nil {
		return err
	}

	if err := a.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := a.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if a.CrossFadeState != 0 && a.CrossFadeState != 1 {
		return errors.New("crossFadeState must be 0 or 1")
	}

	return nil
}
