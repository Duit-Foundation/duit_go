package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_gestures"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type GestureDetectorAttributes[TAction duit_action.Action] struct {
	*ValueReferenceHolder
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
	ExcludeFromSemantics  duit_utils.Tristate[bool]       `json:"excludeFromSemantics,omitempty"`
	DragStarnBehavior     duit_gestures.DragStartBehavior `json:"dragStartBehavior,omitempty"`
	Behavior              duit_gestures.HitTestBehavior   `json:"behavior,omitempty"`
}

func (r *GestureDetectorAttributes[TAction]) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	callbacks := []*TAction{
		r.OnTap,
		r.OnTapDown,
		r.OnTapUp,
		r.OnTapCancel,
		r.OnDoubleTap,
		r.OnDoubleTapDown,
		r.OnDoubleTapCancel,
		r.OnLongPressDown,
		r.OnLongPressCancel,
		r.OnLongPress,
		r.OnLongPressStart,
		r.OnLongPressMoveUpdate,
		r.OnLongPressUp,
		r.OnLongPressEnd,
		r.OnPanStart,
		r.OnPanDown,
		r.OnPanUpdate,
		r.OnPanEnd,
		r.OnPanCancel,
	}

	for _, callback := range callbacks {
		if err := duit_action.CheckActionType(callback); err != nil {
			return err
		}
	}

	return nil
}
