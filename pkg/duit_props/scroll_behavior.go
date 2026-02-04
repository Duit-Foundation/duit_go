package duit_props

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type ScrollBehavior struct {
	Overscroll             duit_utils.TBool       `json:"overscroll,omitempty"`
	Scrollbars             duit_utils.TBool       `json:"scrollbars,omitempty"`
	Physics                ScrollPhysics          `json:"physics,omitempty"`
	MultiTouchDragStrategy MultiTouchDragStrategy `json:"multiTouchDragStrategy,omitempty"`
	Platform               TargetPlatform         `json:"platform,omitempty"`
	DragDevices            []PointerDeviceKind    `json:"dragDevices,omitempty"`
	PointerAxisModifiers   []LogicalKeyboardKey   `json:"pointerAxisModifiers,omitempty"`
}

func NewScrollBehavior() *ScrollBehavior {
	return &ScrollBehavior{}
}

// SetOverscroll sets whether overscroll is enabled.
func (r *ScrollBehavior) SetOverscroll(value bool) *ScrollBehavior {
	r.Overscroll = duit_utils.TristateFrom[bool](value)
	return r
}

// SetScrollbars sets whether scrollbars are visible.
func (r *ScrollBehavior) SetScrollbars(value bool) *ScrollBehavior {
	r.Scrollbars = duit_utils.TristateFrom[bool](value)
	return r
}

// SetPhysics sets the scroll physics.
func (r *ScrollBehavior) SetPhysics(value ScrollPhysics) *ScrollBehavior {
	r.Physics = value
	return r
}

// SetMultiTouchDragStrategy sets the multi-touch drag strategy.
func (r *ScrollBehavior) SetMultiTouchDragStrategy(value MultiTouchDragStrategy) *ScrollBehavior {
	r.MultiTouchDragStrategy = value
	return r
}

// SetPlatform sets the target platform.
func (r *ScrollBehavior) SetPlatform(value TargetPlatform) *ScrollBehavior {
	r.Platform = value
	return r
}

// SetDragDevices sets the pointer device kinds used for dragging.
func (r *ScrollBehavior) SetDragDevices(value []PointerDeviceKind) *ScrollBehavior {
	r.DragDevices = value
	return r
}

// SetPointerAxisModifiers sets the logical keyboard keys used as pointer axis modifiers.
func (r *ScrollBehavior) SetPointerAxisModifiers(value []LogicalKeyboardKey) *ScrollBehavior {
	r.PointerAxisModifiers = value
	return r
}