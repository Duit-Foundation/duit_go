package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// CustomScrollViewAttributes defines attributes for CustomScrollView widget.
// See: https://api.flutter.dev/flutter/widgets/CustomScrollView-class.html
type CustomScrollViewAttributes struct {
	*ValueReferenceHolder   `gen:"wrap"`
	*ThemeConsumer          `gen:"wrap"`
	Physics                 duit_props.ScrollPhysics                     `json:"physics,omitempty"`
	Reverse                 duit_utils.Tristate[bool]                    `json:"reverse,omitempty"`
	Primary                 duit_utils.Tristate[bool]                    `json:"primary,omitempty"`
	ShrinkWrap              duit_utils.Tristate[bool]                    `json:"shrinkWrap,omitempty"`
	ScrollDirection         duit_props.Axis                              `json:"scrollDirection,omitempty"`
	Anchor                  float32                                      `json:"anchors,omitempty"`
	SemantickChildCount     int                                          `json:"semanticChildCount,omitempty"`
	ClipBehavior            duit_props.Clip                              `json:"clipBehavior,omitempty"`
	RestorationId           string                                       `json:"restorationId,omitempty"`
	DragStarnBehavior       duit_props.DragStartBehavior                 `json:"dragStartBehavior,omitempty"`
	KeyboardDismissBehavior duit_props.ScrollViewKeyboardDismissBehavior `json:"keyboardDismissBehavior,omitempty"`
	HitTestBehavior         duit_props.HitTestBehavior                   `json:"hitTestBehavior,omitempty"`
	CacheExtent             float32                                      `json:"cacheExtent,omitempty"`
	Center                  string                                       `json:"center,omitempty"`
}

//bindgen:exclude
func (r *CustomScrollViewAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	return nil
}

// NewCustomScrollViewAttributes creates a new CustomScrollViewAttributes instance.
func NewCustomScrollViewAttributes() *CustomScrollViewAttributes {
	return &CustomScrollViewAttributes{}
}

// SetPhysics sets the physics for the custom scroll view widget.
func (r *CustomScrollViewAttributes) SetPhysics(physics duit_props.ScrollPhysics) *CustomScrollViewAttributes {
	r.Physics = physics
	return r
}

// SetReverse sets the reverse for the custom scroll view widget.
func (r *CustomScrollViewAttributes) SetReverse(reverse duit_utils.Tristate[bool]) *CustomScrollViewAttributes {
	r.Reverse = reverse
	return r
}

// SetPrimary sets the primary for the custom scroll view widget.
func (r *CustomScrollViewAttributes) SetPrimary(primary duit_utils.Tristate[bool]) *CustomScrollViewAttributes {
	r.Primary = primary
	return r
}

// SetShrinkWrap sets the shrink wrap for the custom scroll view widget.
func (r *CustomScrollViewAttributes) SetShrinkWrap(shrinkWrap duit_utils.Tristate[bool]) *CustomScrollViewAttributes {
	r.ShrinkWrap = shrinkWrap
	return r
}

// SetScrollDirection sets the scroll direction for the custom scroll view widget.
func (r *CustomScrollViewAttributes) SetScrollDirection(scrollDirection duit_props.Axis) *CustomScrollViewAttributes {
	r.ScrollDirection = scrollDirection
	return r
}

// SetAnchor sets the anchor for the custom scroll view widget.
func (r *CustomScrollViewAttributes) SetAnchor(anchor float32) *CustomScrollViewAttributes {
	r.Anchor = anchor
	return r
}

// SetSemantickChildCount sets the semantic child count for the custom scroll view widget.
func (r *CustomScrollViewAttributes) SetSemantickChildCount(semantickChildCount int) *CustomScrollViewAttributes {
	r.SemantickChildCount = semantickChildCount
	return r
}

// SetClipBehavior sets the clip behavior for the custom scroll view widget.
func (r *CustomScrollViewAttributes) SetClipBehavior(clipBehavior duit_props.Clip) *CustomScrollViewAttributes {
	r.ClipBehavior = clipBehavior
	return r
}

// SetRestorationId sets the restoration id for the custom scroll view widget.
func (r *CustomScrollViewAttributes) SetRestorationId(restorationId string) *CustomScrollViewAttributes {
	r.RestorationId = restorationId
	return r
}

// SetDragStarnBehavior sets the drag start behavior for the custom scroll view widget.
func (r *CustomScrollViewAttributes) SetDragStarnBehavior(dragStarnBehavior duit_props.DragStartBehavior) *CustomScrollViewAttributes {
	r.DragStarnBehavior = dragStarnBehavior
	return r
}

// SetKeyboardDismissBehavior sets the keyboard dismiss behavior for the custom scroll view widget.
func (r *CustomScrollViewAttributes) SetKeyboardDismissBehavior(keyboardDismissBehavior duit_props.ScrollViewKeyboardDismissBehavior) *CustomScrollViewAttributes {
	r.KeyboardDismissBehavior = keyboardDismissBehavior
	return r
}

// SetHitTestBehavior sets the hit test behavior for the custom scroll view widget.
func (r *CustomScrollViewAttributes) SetHitTestBehavior(hitTestBehavior duit_props.HitTestBehavior) *CustomScrollViewAttributes {
	r.HitTestBehavior = hitTestBehavior
	return r
}

// SetCacheExtent sets the cache extent for the custom scroll view widget.
func (r *CustomScrollViewAttributes) SetCacheExtent(cacheExtent float32) *CustomScrollViewAttributes {
	r.CacheExtent = cacheExtent
	return r
}

// SetCenter sets the center for the custom scroll view widget.
func (r *CustomScrollViewAttributes) SetCenter(center string) *CustomScrollViewAttributes {
	r.Center = center
	return r
}
