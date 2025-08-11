package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_alignment"
	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
)

type AnimatedCrossFadeAttributes struct {
	ValueReferenceHolder
	animations.ImplicitAnimatable
	ThemeConsumer
	ReverseDuration    uint                     `json:"reverseDuration,omitempty"`
	FirstCurve         animations.Curves        `json:"firstCurve,omitempty"`
	SecondCurve        animations.Curves        `json:"secondCurve,omitempty"`
	SizeCurve          animations.Curves        `json:"sizeCurve,omitempty"`
	ExcludeBottomFocus *bool                    `json:"excludeBottomFocus,omitempty"`
	Alignment          duit_alignment.Alignment `json:"alignment,omitempty"`
	/// Where 0 - CrossFadeState.showFirst and 1 - CrossFadeState.showSecond
	CrossFadeState     uint8                    `json:"crossFadeState"`
}
