package duit_animations

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_alignment"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_color"
	duit_decoration "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_flex"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_text_properties"
)

func CheckTweenType(tween any) error {
	switch tween.(type) {
	case *tweenBase[float32]:
	case *tweenBase[*duit_color.ColorRGBO]:
    case *tweenBase[duit_color.ColorString]:
	case *tweenBase[duit_text_properties.TextStyle[*duit_color.ColorRGBO]]:
	case *tweenBase[duit_text_properties.TextStyle[duit_color.ColorString]]:
	case *tweenBase[duit_decoration.BoxDecoration[*duit_color.ColorRGBO]]:
	case *tweenBase[duit_decoration.BoxDecoration[duit_color.ColorString]]:
	case *tweenBase[duit_alignment.Alignment]:
	case *tweenBase[duit_edge_insets.EdgeInsetsAll]:
	case *tweenBase[duit_edge_insets.EdgeInsetsSymmentric]:
	case *tweenBase[duit_edge_insets.EdgeInsetsLTRB]:
	case *tweenBase[duit_flex.BoxConstraints]:
	case *tweenBase[duit_flex.Size]:
	case *tweenBase[duit_decoration.Border]:
	case *tweenGroup:
	case nil:
		return nil
	default:
		return errors.New("invalid tween type. Must be instance of tweenBase or tweenGroup or nil")
	}
    return nil
}
