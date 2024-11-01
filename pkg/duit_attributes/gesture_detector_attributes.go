package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/pkg/duit_attributes/duit_gestures"
	"github.com/Duit-Foundation/duit_go/pkg/duit_core"
)

type GestureDetectorAttributes struct {
	ValueReferenceHolder
	OnTap                 *duit_core.Action               `json:"onTap,omitempty"`
	OnTapDown             *duit_core.Action               `json:"onTapDown,omitempty"`
	OnTapUp               *duit_core.Action               `json:"onTapUp,omitempty"`
	OnTapCancel           *duit_core.Action               `json:"onTapCancel,omitempty"`
	OnDoubleTap           *duit_core.Action               `json:"onDoubleTap,omitempty"`
	OnDoubleTapDown       *duit_core.Action               `json:"onDoubleTapDown,omitempty"`
	OnDoubleTapCancel     *duit_core.Action               `json:"onDoubleTapCancel,omitempty"`
	OnLongPressDown       *duit_core.Action               `json:"onLongPressDown,omitempty"`
	OnLongPressCancel     *duit_core.Action               `json:"onLongPressCancel,omitempty"`
	OnLongPress           *duit_core.Action               `json:"onLongPress,omitempty"`
	OnLongPressStart      *duit_core.Action               `json:"onLongPressStart,omitempty"`
	OnLongPressMoveUpdate *duit_core.Action               `json:"onLongPressMoveUpdate,omitempty"`
	OnLongPressUp         *duit_core.Action               `json:"onLongPressUp,omitempty"`
	OnLongPressEnd        *duit_core.Action               `json:"onLongPressEnd,omitempty"`
	OnPanStart            *duit_core.Action               `json:"onPanStart,omitempty"`
	OnPanDown             *duit_core.Action               `json:"onPanDown,omitempty"`
	OnPanUpdate           *duit_core.Action               `json:"onPanUpdate,omitempty"`
	OnPanEnd              *duit_core.Action               `json:"onPanEnd,omitempty"`
	OnPanCancel           *duit_core.Action               `json:"onPanCancel,omitempty"`
	ExcludeFromSemantics  bool                            `json:"excludeFromSemantics,omitempty"`
	DragStarnBehavior     duit_gestures.DragStartBehavior `json:"dragStartBehavior,omitempty"`
	Behavior              duit_gestures.HitTestBehavior   `json:"behavior,omitempty"`
}
