package duit_animations

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_alignment"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_color"
	duit_decoration "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_flex"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_text_properties"
)

type tweenBase[T any] struct {
	TweenType       string           `json:"type"`
	AnimatedPropKey string           `json:"animatedPropKey"`
	Duration        int              `json:"duration"`
	Curve           Curves           `json:"curve,omitempty"`
	Begin           T                `json:"begin"`
	End             T                `json:"end"`
	ReverseOnRepeat bool             `json:"reverseOnRepeat,omitempty"`
	Trigger         AnimationTrigger `json:"trigger"`
	Method          AnimationMethod  `json:"method"`
}

type tweenGroup struct {
	TweenType       string           `json:"type"`
	GroupId         string           `json:"groupId"`
	Tweens          []any            `json:"tweens"`
	Duration        int              `json:"duration"`
	Trigger         AnimationTrigger `json:"trigger"`
	Method          AnimationMethod  `json:"method"`
	Curve           Curves           `json:"curve,omitempty"`
	ReverseOnRepeat bool             `json:"reverseOnRepeat,omitempty"`
}

func TweenGroup(tweens []any, groupId string, duration int, trigger AnimationTrigger, method AnimationMethod, reverseOnRepeat bool, curve Curves) *tweenGroup {
	return &tweenGroup{
		Tweens:          tweens,
		GroupId:         groupId,
		TweenType:       "group",
		Duration:        duration,
		Curve:           prepareCurveValue(curve),
		ReverseOnRepeat: reverseOnRepeat,
		Trigger:         trigger,
		Method:          method,
	}
}

func prepareCurveValue(curve Curves) Curves {
	if len(curve) > 0 {
		return curve
	} else {
		return linear
	}
}

func Tween(animatedPropKey string, begin, end float32, duration int, trigger AnimationTrigger, method AnimationMethod, reverseOnRepeat bool, curve Curves) *tweenBase[float32] {
	return &tweenBase[float32]{
		TweenType:       "tween",
		AnimatedPropKey: animatedPropKey,
		Duration:        duration,
		Curve:           prepareCurveValue(curve),
		Begin:           begin,
		End:             end,
		ReverseOnRepeat: reverseOnRepeat,
		Trigger:         trigger,
		Method:          method,
	}
}

func ColorTween[TColor duit_color.Color](animatedPropKey string, begin, end TColor, duration int, trigger AnimationTrigger, method AnimationMethod, reverseOnRepeat bool, curve Curves) *tweenBase[TColor] {
	return &tweenBase[TColor]{
		TweenType:       "colorTween",
		AnimatedPropKey: animatedPropKey,
		Duration:        duration,
		Curve:           prepareCurveValue(curve),
		Begin:           begin,
		End:             end,
		ReverseOnRepeat: reverseOnRepeat,
		Trigger:         trigger,
		Method:          method,
	}
}

func TextStyleTween[TColor duit_color.Color](animatedPropKey string, begin, end duit_text_properties.TextStyle[TColor], duration int, trigger AnimationTrigger, method AnimationMethod, reverseOnRepeat bool, curve Curves) *tweenBase[duit_text_properties.TextStyle[TColor]] {
	return &tweenBase[duit_text_properties.TextStyle[TColor]]{
		TweenType:       "textStyleTween",
		AnimatedPropKey: animatedPropKey,
		Duration:        duration,
		Curve:           prepareCurveValue(curve),
		Begin:           begin,
		End:             end,
		ReverseOnRepeat: reverseOnRepeat,
		Trigger:         trigger,
		Method:          method,
	}
}

func DecorationTween[TColor duit_color.Color](animatedPropKey string, begin, end duit_decoration.BoxDecoration[TColor], duration int, trigger AnimationTrigger, method AnimationMethod, reverseOnRepeat bool, curve Curves) *tweenBase[duit_decoration.BoxDecoration[TColor]] {
	return &tweenBase[duit_decoration.BoxDecoration[TColor]]{
		TweenType:       "decorationTween",
		AnimatedPropKey: animatedPropKey,
		Duration:        duration,
		Curve:           prepareCurveValue(curve),
		Begin:           begin,
		End:             end,
		ReverseOnRepeat: reverseOnRepeat,
		Trigger:         trigger,
		Method:          method,
	}
}

func AlignmentTween(animatedPropKey string, begin, end duit_alignment.Alignment, duration int, trigger AnimationTrigger, method AnimationMethod, reverseOnRepeat bool, curve Curves) *tweenBase[duit_alignment.Alignment] {
	return &tweenBase[duit_alignment.Alignment]{
		TweenType:       "alignmentTween",
		AnimatedPropKey: animatedPropKey,
		Duration:        duration,
		Curve:           prepareCurveValue(curve),
		Begin:           begin,
		End:             end,
		ReverseOnRepeat: reverseOnRepeat,
		Trigger:         trigger,
		Method:          method,
	}
}

func EdgeInsetsTween[TInsets duit_edge_insets.EdgeInsets](animatedPropKey string, begin, end TInsets, duration int, trigger AnimationTrigger, method AnimationMethod, reverseOnRepeat bool, curve Curves) *tweenBase[TInsets] {
	return &tweenBase[TInsets]{
		TweenType:       "edgeInsetsTween",
		AnimatedPropKey: animatedPropKey,
		Duration:        duration,
		Curve:           prepareCurveValue(curve),
		Begin:           begin,
		End:             end,
		ReverseOnRepeat: reverseOnRepeat,
		Trigger:         trigger,
		Method:          method,
	}
}

func BoxConstraintsTween(animatedPropKey string, begin, end duit_flex.BoxConstraints, duration int, trigger AnimationTrigger, method AnimationMethod, reverseOnRepeat bool, curve Curves) *tweenBase[duit_flex.BoxConstraints] {
	return &tweenBase[duit_flex.BoxConstraints]{
		TweenType:       "boxConstraintsTween",
		AnimatedPropKey: animatedPropKey,
		Duration:        duration,
		Curve:           prepareCurveValue(curve),
		Begin:           begin,
		End:             end,
		ReverseOnRepeat: reverseOnRepeat,
		Trigger:         trigger,
		Method:          method,
	}
}

func SizeTween(animatedPropKey string, begin, end duit_flex.Size, duration int, trigger AnimationTrigger, method AnimationMethod, reverseOnRepeat bool, curve Curves) *tweenBase[duit_flex.Size] {
	return &tweenBase[duit_flex.Size]{
		TweenType:       "sizeTween",
		AnimatedPropKey: animatedPropKey,
		Duration:        duration,
		Curve:           prepareCurveValue(curve),
		Begin:           begin,
		End:             end,
		ReverseOnRepeat: reverseOnRepeat,
		Trigger:         trigger,
		Method:          method,
	}
}

func BorderTween(animatedPropKey string, begin, end duit_decoration.Border, duration int, trigger AnimationTrigger, method AnimationMethod, reverseOnRepeat bool, curve Curves) *tweenBase[duit_decoration.Border] {
	return &tweenBase[duit_decoration.Border]{
		TweenType:       "borderTween",
		AnimatedPropKey: animatedPropKey,
		Duration:        duration,
		Curve:           prepareCurveValue(curve),
		Begin:           begin,
		End:             end,
		ReverseOnRepeat: reverseOnRepeat,
		Trigger:         trigger,
		Method:          method,
	}
}
