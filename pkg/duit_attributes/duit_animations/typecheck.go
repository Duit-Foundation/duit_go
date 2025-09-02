package duit_animations

import (
	"errors"

	duit_decoration "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_flex"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

func CheckTweenType(tween any) error {
	switch tween.(type) {
	case *tweenBase[float32],
		*tweenBase[*duit_props.ColorRGBO],
		*tweenBase[duit_props.ColorString],
		*tweenBase[duit_props.TextStyle[*duit_props.ColorRGBO]],
		*tweenBase[duit_props.TextStyle[duit_props.ColorString]],
		*tweenBase[duit_decoration.BoxDecoration[*duit_props.ColorRGBO]],
		*tweenBase[duit_decoration.BoxDecoration[duit_props.ColorString]],
		*tweenBase[duit_props.EdgeInsetsSymmentric],
		*tweenBase[duit_props.EdgeInsetsLTRB],
		*tweenBase[duit_flex.BoxConstraints],
		*tweenBase[duit_flex.Size],
		*tweenBase[duit_decoration.Border],
		*tweenGroup, nil:
		return nil
	default:
		return errors.New("invalid tween type. Must be instance of tweenBase or tweenGroup or nil")
	}
}
