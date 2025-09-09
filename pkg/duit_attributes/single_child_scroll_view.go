package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// SingleChildScrollViewAttributes defines attributes for SingleChildScrollView widget.
// See: https://api.flutter.dev/flutter/widgets/SingleChildScrollView-class.html
type SingleChildScrollViewAttributes struct {
	*ValueReferenceHolder   `gen:"wrap"`
	*ThemeConsumer          `gen:"wrap"`
	ScrollDirection         duit_props.Axis                              `json:"scrollDirection,omitempty"`
	Reverse                 duit_utils.Tristate[bool]                    `json:"reverse,omitempty"`
	Primary                 duit_utils.Tristate[bool]                    `json:"primary,omitempty"`
	Padding                 *duit_props.EdgeInsets                       `json:"padding,omitempty"`
	RestorationId           string                                       `json:"restorationId,omitempty"`
	ClipBehavior            duit_props.Clip                              `json:"clipBehavior,omitempty"`
	DragStartBehavior       duit_props.DragStartBehavior                 `json:"dragStartBehavior,omitempty"`
	KeyboardDismissBehavior duit_props.ScrollViewKeyboardDismissBehavior `json:"keyboardDismissBehavior,omitempty"`
	Physics                 duit_props.ScrollPhysics                     `json:"physics,omitempty"`
}

//bindgen:exclude
func (r *SingleChildScrollViewAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	return nil
}

// NewSingleChildScrollViewAttributes creates a new SingleChildScrollViewAttributes instance.
func NewSingleChildScrollViewAttributes() *SingleChildScrollViewAttributes {
	return &SingleChildScrollViewAttributes{}
}

// SetScrollDirection sets the scroll direction for the single child scroll view widget.
func (r *SingleChildScrollViewAttributes) SetScrollDirection(scrollDirection duit_props.Axis) *SingleChildScrollViewAttributes {
	r.ScrollDirection = scrollDirection
	return r
}

// SetReverse sets the reverse for the single child scroll view widget.
func (r *SingleChildScrollViewAttributes) SetReverse(reverse duit_utils.Tristate[bool]) *SingleChildScrollViewAttributes {
	r.Reverse = reverse
	return r
}

// SetPrimary sets the primary for the single child scroll view widget.
func (r *SingleChildScrollViewAttributes) SetPrimary(primary duit_utils.Tristate[bool]) *SingleChildScrollViewAttributes {
	r.Primary = primary
	return r
}

// SetPadding sets the padding for the single child scroll view widget.
func (r *SingleChildScrollViewAttributes) SetPadding(padding *duit_props.EdgeInsets) *SingleChildScrollViewAttributes {
	r.Padding = padding
	return r
}

// SetRestorationId sets the restoration id for the single child scroll view widget.
func (r *SingleChildScrollViewAttributes) SetRestorationId(restorationId string) *SingleChildScrollViewAttributes {
	r.RestorationId = restorationId
	return r
}

// SetClipBehavior sets the clip behavior for the single child scroll view widget.
func (r *SingleChildScrollViewAttributes) SetClipBehavior(clipBehavior duit_props.Clip) *SingleChildScrollViewAttributes {
	r.ClipBehavior = clipBehavior
	return r
}

// SetDragStartBehavior sets the drag start behavior for the single child scroll view widget.
func (r *SingleChildScrollViewAttributes) SetDragStartBehavior(dragStartBehavior duit_props.DragStartBehavior) *SingleChildScrollViewAttributes {
	r.DragStartBehavior = dragStartBehavior
	return r
}

// SetKeyboardDismissBehavior sets the keyboard dismiss behavior for the single child scroll view widget.
func (r *SingleChildScrollViewAttributes) SetKeyboardDismissBehavior(keyboardDismissBehavior duit_props.ScrollViewKeyboardDismissBehavior) *SingleChildScrollViewAttributes {
	r.KeyboardDismissBehavior = keyboardDismissBehavior
	return r
}

// SetPhysics sets the physics for the single child scroll view widget.
func (r *SingleChildScrollViewAttributes) SetPhysics(physics duit_props.ScrollPhysics) *SingleChildScrollViewAttributes {
	r.Physics = physics
	return r
}
