package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type GestureDetectorAttributes struct {
	*ValueReferenceHolder
	OnTap                 any                          `json:"onTap,omitempty"`
	OnTapDown             any                          `json:"onTapDown,omitempty"`
	OnTapUp               any                          `json:"onTapUp,omitempty"`
	OnTapCancel           any                          `json:"onTapCancel,omitempty"`
	OnDoubleTap           any                          `json:"onDoubleTap,omitempty"`
	OnDoubleTapDown       any                          `json:"onDoubleTapDown,omitempty"`
	OnDoubleTapCancel     any                          `json:"onDoubleTapCancel,omitempty"`
	OnLongPressDown       any                          `json:"onLongPressDown,omitempty"`
	OnLongPressCancel     any                          `json:"onLongPressCancel,omitempty"`
	OnLongPress           any                          `json:"onLongPress,omitempty"`
	OnLongPressStart      any                          `json:"onLongPressStart,omitempty"`
	OnLongPressMoveUpdate any                          `json:"onLongPressMoveUpdate,omitempty"`
	OnLongPressUp         any                          `json:"onLongPressUp,omitempty"`
	OnLongPressEnd        any                          `json:"onLongPressEnd,omitempty"`
	OnPanStart            any                          `json:"onPanStart,omitempty"`
	OnPanDown             any                          `json:"onPanDown,omitempty"`
	OnPanUpdate           any                          `json:"onPanUpdate,omitempty"`
	OnPanEnd              any                          `json:"onPanEnd,omitempty"`
	OnPanCancel           any                          `json:"onPanCancel,omitempty"`
	ExcludeFromSemantics  duit_utils.Tristate[bool]    `json:"excludeFromSemantics,omitempty"`
	DragStarnBehavior     duit_props.DragStartBehavior `json:"dragStartBehavior,omitempty"`
	Behavior              duit_props.HitTestBehavior   `json:"behavior,omitempty"`
}

func (r *GestureDetectorAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	callbacks := []any{
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
