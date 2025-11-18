package duit_attributes

import (
	"encoding/json"
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type PageViewConstructor uint

const (
	PageViewCommon PageViewConstructor = iota
	PageViewBuilder
)

type DefaultPageViewAttributes struct {
	Constructor             duit_utils.Tristate[PageViewConstructor]     `json:"constructor"`
	ScrollDirection         duit_props.Axis                              `json:"scrollDirection,omitempty"`
	Reverse                 duit_utils.TBool                             `json:"reverse,omitempty"`
	Primary                 duit_utils.TBool                             `json:"primary,omitempty"`
	PageSnapping            duit_utils.TBool                             `json:"pageSnapping,omitempty"`
	DragStartBehavior       duit_props.DragStartBehavior                 `json:"dragStartBehavior,omitempty"`
	AllowImplicitScrolling  duit_utils.TBool                             `json:"allowImplicitScrolling,omitempty"`
	ClipBehavior            duit_props.Clip                              `json:"clipBehavior,omitempty"`
	KeyboardDismissBehavior duit_props.ScrollViewKeyboardDismissBehavior `json:"keyboardDismissBehavior,omitempty"`
	Physics                 duit_props.ScrollPhysics                     `json:"physics,omitempty"`
	RestorationId           string                                       `json:"restorationId,omitempty"`
	HitTestBehavior         duit_props.HitTestBehavior                   `json:"hitTestBehavior,omitempty"`
	PadsEnd                 duit_utils.TBool                             `json:"padsEnd,omitempty"`
}

// SetConstructor assigns the PageView constructor variant.
//
//bindgen:exclude
func (r *DefaultPageViewAttributes) SetConstructor(value PageViewConstructor) *DefaultPageViewAttributes {
	r.Constructor = duit_utils.TristateFrom[PageViewConstructor](value)
	return r
}

// SetScrollDirection defines the axis used for page scrolling.
func (r *DefaultPageViewAttributes) SetScrollDirection(value duit_props.Axis) *DefaultPageViewAttributes {
	r.ScrollDirection = value
	return r
}

// SetReverse toggles reverse scroll direction behavior.
func (r *DefaultPageViewAttributes) SetReverse(value bool) *DefaultPageViewAttributes {
	r.Reverse = duit_utils.BoolValue(value)
	return r
}

// SetPrimary marks the scroll controller as primary.
func (r *DefaultPageViewAttributes) SetPrimary(value bool) *DefaultPageViewAttributes {
	r.Primary = duit_utils.BoolValue(value)
	return r
}

// SetPageSnapping enables or disables page snapping.
func (r *DefaultPageViewAttributes) SetPageSnapping(value bool) *DefaultPageViewAttributes {
	r.PageSnapping = duit_utils.BoolValue(value)
	return r
}

// SetDragStartBehavior sets the drag start behavior.
func (r *DefaultPageViewAttributes) SetDragStartBehavior(value duit_props.DragStartBehavior) *DefaultPageViewAttributes {
	r.DragStartBehavior = value
	return r
}

// SetAllowImplicitScrolling toggles implicit scrolling support.
func (r *DefaultPageViewAttributes) SetAllowImplicitScrolling(value bool) *DefaultPageViewAttributes {
	r.AllowImplicitScrolling = duit_utils.BoolValue(value)
	return r
}

// SetClipBehavior configures the clipping behavior.
func (r *DefaultPageViewAttributes) SetClipBehavior(value duit_props.Clip) *DefaultPageViewAttributes {
	r.ClipBehavior = value
	return r
}

// SetKeyboardDismissBehavior defines when the keyboard dismisses.
func (r *DefaultPageViewAttributes) SetKeyboardDismissBehavior(value duit_props.ScrollViewKeyboardDismissBehavior) *DefaultPageViewAttributes {
	r.KeyboardDismissBehavior = value
	return r
}

// SetPhysics assigns the scroll physics.
func (r *DefaultPageViewAttributes) SetPhysics(value duit_props.ScrollPhysics) *DefaultPageViewAttributes {
	r.Physics = value
	return r
}

// SetRestorationId sets the state restoration identifier.
func (r *DefaultPageViewAttributes) SetRestorationId(value string) *DefaultPageViewAttributes {
	r.RestorationId = value
	return r
}

// SetHitTestBehavior configures the hit test interaction behavior.
func (r *DefaultPageViewAttributes) SetHitTestBehavior(value duit_props.HitTestBehavior) *DefaultPageViewAttributes {
	r.HitTestBehavior = value
	return r
}

// SetPadsEnd toggles end padding behavior.
func (r *DefaultPageViewAttributes) SetPadsEnd(value bool) *DefaultPageViewAttributes {
	r.PadsEnd = duit_utils.BoolValue(value)
	return r
}

//bindgen:exclude
func (t *DefaultPageViewAttributes) Validate() error {
	if t == nil {
		return nil
	}

	if t.Constructor == nil {
		return errors.New("constructor property is required")
	}

	return nil
}

// PageViewCommonAttributes defines attributes for PageView widget.
// See: https://api.flutter.dev/flutter/widgets/PageView/PageView.html
type PageViewCommonAttributes struct {
	*ValueReferenceHolder      `gen:"wrap"`
	*ThemeConsumer             `gen:"wrap"`
	*DefaultPageViewAttributes `gen:"wrap"`
}

func NewPageViewCommonAttributes() *PageViewCommonAttributes {
	attr := &PageViewCommonAttributes{}
	dAttrs := &DefaultPageViewAttributes{}
	dAttrs.SetConstructor(PageViewCommon)
	attr.DefaultPageViewAttributes = dAttrs
	return attr
}

func (r *PageViewCommonAttributes) ToGenericAttributes() *PageViewAttributes {
	return &PageViewAttributes{data: r}
}

// Validate validates the GridViewCommonAttributes instance.
func (r *PageViewCommonAttributes) Validate() error {
	// Validate embedded components
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.DefaultPageViewAttributes.Validate(); err != nil {
		return err
	}

	return nil
}

// PageViewBuilderAttributes defines attributes for PageView.builder widget.
// See: https://api.flutter.dev/flutter/widgets/PageView/PageView.builder.html
type PageViewBuilderAttributes struct {
	*ValueReferenceHolder      `gen:"wrap"`
	*ThemeConsumer             `gen:"wrap"`
	*DefaultPageViewAttributes `gen:"wrap"`
	*Builder                   `gen:"wrap"`
}

func NewPageViewBuilderAttributes() *PageViewBuilderAttributes {
	attr := &PageViewBuilderAttributes{}
	dAttrs := &DefaultPageViewAttributes{}
	dAttrs.SetConstructor(PageViewBuilder)
	attr.DefaultPageViewAttributes = dAttrs
	return attr
}

func (r *PageViewBuilderAttributes) ToGenericAttributes() *PageViewAttributes {
	return &PageViewAttributes{data: r}
}

// Validate validates the GridViewCommonAttributes instance.
func (r *PageViewBuilderAttributes) Validate() error {
	// Validate embedded components
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.DefaultPageViewAttributes.Validate(); err != nil {
		return err
	}

	return nil
}

// PageViewAttributes defines attributes for Page widget.
// See: https://api.flutter.dev/flutter/widgets/PageView-class.html
type PageViewAttributes struct {
	data any
}

func (r *PageViewAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.data)
}

func (r *PageViewAttributes) Validate() error {
	switch data := r.data.(type) {
	case *PageViewCommonAttributes:
		return data.Validate()
	case *PageViewBuilderAttributes:
		return data.Validate()
	}
	return nil
}
