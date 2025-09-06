package duit_props

import (
	"errors"
	"fmt"
)

func CheckTweenType(tween any) error {
	switch tween.(type) {
	case *tweenBase[float32],
		*tweenBase[*Color],
		*tweenBase[*TextStyle],
		*tweenBase[*BoxDecoration],
		*tweenBase[*EdgeInsets],
		*tweenBase[*BoxConstraints],
		*tweenBase[Size],
		*tweenBase[Alignment],
		*tweenGroup, nil:
		return nil
	default:
		return errors.New("invalid tween type. Must be instance of tweenBase or tweenGroup or nil. Received: " + fmt.Sprintf("%T", tween))
	}
}
