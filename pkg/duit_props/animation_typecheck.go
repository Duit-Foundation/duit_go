package duit_props

import (
	"errors"
)

func CheckTweenType(tween any) error {
	switch tween.(type) {
	case *tweenBase[float32],
		*tweenBase[*ColorRGBO],
		*tweenBase[ColorString],
		*tweenBase[TextStyle[*ColorRGBO]],
		*tweenBase[TextStyle[ColorString]],
		*tweenBase[BoxDecoration[*ColorRGBO]],
		*tweenBase[BoxDecoration[ColorString]],
		*tweenBase[EdgeInsetsSymmentric],
		*tweenBase[EdgeInsetsLTRB],
		*tweenBase[BoxConstraints],
		*tweenBase[Size],
		*tweenGroup, nil:
		return nil
	default:
		return errors.New("invalid tween type. Must be instance of tweenBase or tweenGroup or nil")
	}
}
