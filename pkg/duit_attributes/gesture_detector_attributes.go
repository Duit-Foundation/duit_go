package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_gestures"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
)

type GestureDetectorAttributes[TAction duit_core.Action] struct {
	ValueReferenceHolder
	OnTap                 *TAction                        `json:"onTap,omitempty"`
	OnTapDown             *TAction                        `json:"onTapDown,omitempty"`
	OnTapUp               *TAction                        `json:"onTapUp,omitempty"`
	OnTapCancel           *TAction                        `json:"onTapCancel,omitempty"`
	OnDoubleTap           *TAction                        `json:"onDoubleTap,omitempty"`
	OnDoubleTapDown       *TAction                        `json:"onDoubleTapDown,omitempty"`
	OnDoubleTapCancel     *TAction                        `json:"onDoubleTapCancel,omitempty"`
	OnLongPressDown       *TAction                        `json:"onLongPressDown,omitempty"`
	OnLongPressCancel     *TAction                        `json:"onLongPressCancel,omitempty"`
	OnLongPress           *TAction                        `json:"onLongPress,omitempty"`
	OnLongPressStart      *TAction                        `json:"onLongPressStart,omitempty"`
	OnLongPressMoveUpdate *TAction                        `json:"onLongPressMoveUpdate,omitempty"`
	OnLongPressUp         *TAction                        `json:"onLongPressUp,omitempty"`
	OnLongPressEnd        *TAction                        `json:"onLongPressEnd,omitempty"`
	OnPanStart            *TAction                        `json:"onPanStart,omitempty"`
	OnPanDown             *TAction                        `json:"onPanDown,omitempty"`
	OnPanUpdate           *TAction                        `json:"onPanUpdate,omitempty"`
	OnPanEnd              *TAction                        `json:"onPanEnd,omitempty"`
	OnPanCancel           *TAction                        `json:"onPanCancel,omitempty"`
	ExcludeFromSemantics  bool                            `json:"excludeFromSemantics,omitempty"`
	DragStarnBehavior     duit_gestures.DragStartBehavior `json:"dragStartBehavior,omitempty"`
	Behavior              duit_gestures.HitTestBehavior   `json:"behavior,omitempty"`
}
