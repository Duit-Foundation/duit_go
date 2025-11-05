package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// GestureDetectorAttributes defines attributes for GestureDetector widget.
// See: https://api.flutter.dev/flutter/widgets/GestureDetector-class.html
type GestureDetectorAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
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

//bindgen:exclude
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

// NewGestureDetectorAttributes creates a new GestureDetectorAttributes instance.
func NewGestureDetectorAttributes() *GestureDetectorAttributes {
	return &GestureDetectorAttributes{}
}

// SetOnTap sets the on tap for the gesture detector widget.
func (r *GestureDetectorAttributes) SetOnTap(onTap any) *GestureDetectorAttributes {
	r.OnTap = onTap
	return r
}

// SetOnTapDown sets the on tap down for the gesture detector widget.
func (r *GestureDetectorAttributes) SetOnTapDown(onTapDown any) *GestureDetectorAttributes {
	r.OnTapDown = onTapDown
	return r
}

// SetOnTapUp sets the on tap up for the gesture detector widget.
func (r *GestureDetectorAttributes) SetOnTapUp(onTapUp any) *GestureDetectorAttributes {
	r.OnTapUp = onTapUp
	return r
}

// SetOnTapCancel sets the on tap cancel for the gesture detector widget.
func (r *GestureDetectorAttributes) SetOnTapCancel(onTapCancel any) *GestureDetectorAttributes {
	r.OnTapCancel = onTapCancel
	return r
}

// SetOnDoubleTap sets the on double tap for the gesture detector widget.
func (r *GestureDetectorAttributes) SetOnDoubleTap(onDoubleTap any) *GestureDetectorAttributes {
	r.OnDoubleTap = onDoubleTap
	return r
}

// SetOnDoubleTapDown sets the on double tap down for the gesture detector widget.
func (r *GestureDetectorAttributes) SetOnDoubleTapDown(onDoubleTapDown any) *GestureDetectorAttributes {
	r.OnDoubleTapDown = onDoubleTapDown
	return r
}

// SetOnDoubleTapCancel sets the on double tap cancel for the gesture detector widget.
func (r *GestureDetectorAttributes) SetOnDoubleTapCancel(onDoubleTapCancel any) *GestureDetectorAttributes {
	r.OnDoubleTapCancel = onDoubleTapCancel
	return r
}

// SetOnLongPressDown sets the on long press down for the gesture detector widget.
func (r *GestureDetectorAttributes) SetOnLongPressDown(onLongPressDown any) *GestureDetectorAttributes {
	r.OnLongPressDown = onLongPressDown
	return r
}

// SetOnLongPressCancel sets the on long press cancel for the gesture detector widget.
func (r *GestureDetectorAttributes) SetOnLongPressCancel(onLongPressCancel any) *GestureDetectorAttributes {
	r.OnLongPressCancel = onLongPressCancel
	return r
}

// SetOnLongPress sets the on long press for the gesture detector widget.
func (r *GestureDetectorAttributes) SetOnLongPress(onLongPress any) *GestureDetectorAttributes {
	r.OnLongPress = onLongPress
	return r
}

// SetOnLongPressStart sets the on long press start for the gesture detector widget.
func (r *GestureDetectorAttributes) SetOnLongPressStart(onLongPressStart any) *GestureDetectorAttributes {
	r.OnLongPressStart = onLongPressStart
	return r
}

// SetOnLongPressMoveUpdate sets the on long press move update for the gesture detector widget.
func (r *GestureDetectorAttributes) SetOnLongPressMoveUpdate(onLongPressMoveUpdate any) *GestureDetectorAttributes {
	r.OnLongPressMoveUpdate = onLongPressMoveUpdate
	return r
}

// SetOnLongPressUp sets the on long press up for the gesture detector widget.
func (r *GestureDetectorAttributes) SetOnLongPressUp(onLongPressUp any) *GestureDetectorAttributes {
	r.OnLongPressUp = onLongPressUp
	return r
}

// SetOnLongPressEnd sets the on long press end for the gesture detector widget.
func (r *GestureDetectorAttributes) SetOnLongPressEnd(onLongPressEnd any) *GestureDetectorAttributes {
	r.OnLongPressEnd = onLongPressEnd
	return r
}

// SetOnPanStart sets the on pan start for the gesture detector widget.
func (r *GestureDetectorAttributes) SetOnPanStart(onPanStart any) *GestureDetectorAttributes {
	r.OnPanStart = onPanStart
	return r
}

// SetOnPanDown sets the on pan down for the gesture detector widget.
func (r *GestureDetectorAttributes) SetOnPanDown(onPanDown any) *GestureDetectorAttributes {
	r.OnPanDown = onPanDown
	return r
}

// SetOnPanUpdate sets the on pan update for the gesture detector widget.
func (r *GestureDetectorAttributes) SetOnPanUpdate(onPanUpdate any) *GestureDetectorAttributes {
	r.OnPanUpdate = onPanUpdate
	return r
}

// SetOnPanEnd sets the on pan end for the gesture detector widget.
func (r *GestureDetectorAttributes) SetOnPanEnd(onPanEnd any) *GestureDetectorAttributes {
	r.OnPanEnd = onPanEnd
	return r
}

// SetOnPanCancel sets the on pan cancel for the gesture detector widget.
func (r *GestureDetectorAttributes) SetOnPanCancel(onPanCancel any) *GestureDetectorAttributes {
	r.OnPanCancel = onPanCancel
	return r
}

// SetExcludeFromSemantics sets the exclude from semantics for the gesture detector widget.
func (r *GestureDetectorAttributes) SetExcludeFromSemantics(excludeFromSemantics duit_utils.Tristate[bool]) *GestureDetectorAttributes {
	r.ExcludeFromSemantics = excludeFromSemantics
	return r
}

// SetDragStarnBehavior sets the drag start behavior for the gesture detector widget.
func (r *GestureDetectorAttributes) SetDragStarnBehavior(dragStarnBehavior duit_props.DragStartBehavior) *GestureDetectorAttributes {
	r.DragStarnBehavior = dragStarnBehavior
	return r
}

// SetBehavior sets the behavior for the gesture detector widget.
func (r *GestureDetectorAttributes) SetBehavior(behavior duit_props.HitTestBehavior) *GestureDetectorAttributes {
	r.Behavior = behavior
	return r
}
