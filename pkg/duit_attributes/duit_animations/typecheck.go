package duit_animations

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

func CheckTweenType(tween any) error {
	switch tween.(type) {
	case *tweenBase[float32],
		*tweenBase[*duit_props.ColorRGBO],
		*tweenBase[duit_props.ColorString],
		*tweenBase[duit_props.TextStyle[*duit_props.ColorRGBO]],
		*tweenBase[duit_props.TextStyle[duit_props.ColorString]],
		*tweenBase[duit_props.BoxDecoration[*duit_props.ColorRGBO]],
		*tweenBase[duit_props.BoxDecoration[duit_props.ColorString]],
		*tweenBase[duit_props.EdgeInsetsSymmentric],
		*tweenBase[duit_props.EdgeInsetsLTRB],
		*tweenBase[duit_props.BoxConstraints],
		*tweenBase[duit_props.Size],
		*tweenGroup, nil:
		return nil
	default:
		return errors.New("invalid tween type. Must be instance of tweenBase or tweenGroup or nil")
	}
}
