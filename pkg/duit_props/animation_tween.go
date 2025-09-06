package duit_props

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
		return CurveLinear
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

func ColorTween(animatedPropKey string, begin, end *Color, duration int, trigger AnimationTrigger, method AnimationMethod, reverseOnRepeat bool, curve Curves) *tweenBase[*Color] {
	return &tweenBase[*Color]{
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

func TextStyleTween[TColor Color](animatedPropKey string, begin, end *TextStyle, duration int, trigger AnimationTrigger, method AnimationMethod, reverseOnRepeat bool, curve Curves) *tweenBase[*TextStyle] {
	return &tweenBase[*TextStyle]{
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

func DecorationTween(animatedPropKey string, begin, end *BoxDecoration, duration int, trigger AnimationTrigger, method AnimationMethod, reverseOnRepeat bool, curve Curves) *tweenBase[*BoxDecoration] {
	return &tweenBase[*BoxDecoration]{
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

func AlignmentTween(animatedPropKey string, begin, end Alignment, duration int, trigger AnimationTrigger, method AnimationMethod, reverseOnRepeat bool, curve Curves) *tweenBase[Alignment] {
	return &tweenBase[Alignment]{
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

func EdgeInsetsTween(animatedPropKey string, begin, end *EdgeInsets, duration int, trigger AnimationTrigger, method AnimationMethod, reverseOnRepeat bool, curve Curves) *tweenBase[*EdgeInsets] {
	return &tweenBase[*EdgeInsets]{
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

func BoxConstraintsTween(animatedPropKey string, begin, end BoxConstraints, duration int, trigger AnimationTrigger, method AnimationMethod, reverseOnRepeat bool, curve Curves) *tweenBase[BoxConstraints] {
	return &tweenBase[BoxConstraints]{
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

func SizeTween(animatedPropKey string, begin, end Size, duration int, trigger AnimationTrigger, method AnimationMethod, reverseOnRepeat bool, curve Curves) *tweenBase[Size] {
	return &tweenBase[Size]{
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

func BorderTween(animatedPropKey string, begin, end Border, duration int, trigger AnimationTrigger, method AnimationMethod, reverseOnRepeat bool, curve Curves) *tweenBase[Border] {
	return &tweenBase[Border]{
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
