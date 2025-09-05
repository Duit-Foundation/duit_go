package duit_props

import (
	"errors"
)

func CheckTweenType(tween any) error {
	switch tween.(type) {
	case *tweenBase[float32],
		*tweenBase[*Color],
		*tweenBase[*TextStyle],
		*tweenBase[*BoxDecoration],
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
