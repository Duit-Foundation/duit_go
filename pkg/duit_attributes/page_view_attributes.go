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

type DefautlPageViewAttributes struct {
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
func (r *DefautlPageViewAttributes) SetConstructor(value PageViewConstructor) *DefautlPageViewAttributes {
	r.Constructor = duit_utils.TristateFrom[PageViewConstructor](value)
	return r
}

// SetScrollDirection defines the axis used for page scrolling.
func (r *DefautlPageViewAttributes) SetScrollDirection(value duit_props.Axis) *DefautlPageViewAttributes {
	r.ScrollDirection = value
	return r
}

// SetReverse toggles reverse scroll direction behavior.
func (r *DefautlPageViewAttributes) SetReverse(value bool) *DefautlPageViewAttributes {
	r.Reverse = duit_utils.BoolValue(value)
	return r
}

// SetPrimary marks the scroll controller as primary.
func (r *DefautlPageViewAttributes) SetPrimary(value bool) *DefautlPageViewAttributes {
	r.Primary = duit_utils.BoolValue(value)
	return r
}

// SetPageSnapping enables or disables page snapping.
func (r *DefautlPageViewAttributes) SetPageSnapping(value bool) *DefautlPageViewAttributes {
	r.PageSnapping = duit_utils.BoolValue(value)
	return r
}

// SetDragStartBehavior sets the drag start behavior.
func (r *DefautlPageViewAttributes) SetDragStartBehavior(value duit_props.DragStartBehavior) *DefautlPageViewAttributes {
	r.DragStartBehavior = value
	return r
}

// SetAllowImplicitScrolling toggles implicit scrolling support.
func (r *DefautlPageViewAttributes) SetAllowImplicitScrolling(value bool) *DefautlPageViewAttributes {
	r.AllowImplicitScrolling = duit_utils.BoolValue(value)
	return r
}

// SetClipBehavior configures the clipping behavior.
func (r *DefautlPageViewAttributes) SetClipBehavior(value duit_props.Clip) *DefautlPageViewAttributes {
	r.ClipBehavior = value
	return r
}

// SetKeyboardDismissBehavior defines when the keyboard dismisses.
func (r *DefautlPageViewAttributes) SetKeyboardDismissBehavior(value duit_props.ScrollViewKeyboardDismissBehavior) *DefautlPageViewAttributes {
	r.KeyboardDismissBehavior = value
	return r
}

// SetPhysics assigns the scroll physics.
func (r *DefautlPageViewAttributes) SetPhysics(value duit_props.ScrollPhysics) *DefautlPageViewAttributes {
	r.Physics = value
	return r
}

// SetRestorationId sets the state restoration identifier.
func (r *DefautlPageViewAttributes) SetRestorationId(value string) *DefautlPageViewAttributes {
	r.RestorationId = value
	return r
}

// SetHitTestBehavior configures the hit test interaction behavior.
func (r *DefautlPageViewAttributes) SetHitTestBehavior(value duit_props.HitTestBehavior) *DefautlPageViewAttributes {
	r.HitTestBehavior = value
	return r
}

// SetPadsEnd toggles end padding behavior.
func (r *DefautlPageViewAttributes) SetPadsEnd(value bool) *DefautlPageViewAttributes {
	r.PadsEnd = duit_utils.BoolValue(value)
	return r
}

//bindgen:exclude
func (t *DefautlPageViewAttributes) Validate() error {
	if t == nil {
		return nil
	}

	if t.Constructor == nil {
		return errors.New("constructor property is required")
	}

	return nil
}

// PhysicalModelAttributes defines attributes for PhysicalModel widget.
// See: https://api.flutter.dev/flutter/widgets/PageView/PageView.html
type PageViewCommonAttributes struct {
	*ValueReferenceHolder      `gen:"wrap"`
	*ThemeConsumer             `gen:"wrap"`
	*DefautlPageViewAttributes `gen:"wrap"`
}

func NewPageViewCommonAttributes() *PageViewCommonAttributes {
	attr := &PageViewCommonAttributes{}
	dAttrs := &DefautlPageViewAttributes{}
	dAttrs.SetConstructor(PageViewCommon)
	attr.DefautlPageViewAttributes = dAttrs
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

	if err := r.DefautlPageViewAttributes.Validate(); err != nil {
		return err
	}

	return nil
}

// PhysicalModelAttributes defines attributes for PhysicalModel widget.
// See: https://api.flutter.dev/flutter/widgets/PageView/PageView.builder.html
type PageViewBuilderAttributes struct {
	*ValueReferenceHolder      `gen:"wrap"`
	*ThemeConsumer             `gen:"wrap"`
	*DefautlPageViewAttributes `gen:"wrap"`
	*Builder                   `gen:"wrap"`
}

func NewPageViewBuilderAttributes() *PageViewBuilderAttributes {
	attr := &PageViewBuilderAttributes{}
	dAttrs := &DefautlPageViewAttributes{}
	dAttrs.SetConstructor(PageViewBuilder)
	attr.DefautlPageViewAttributes = dAttrs
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

	if err := r.DefautlPageViewAttributes.Validate(); err != nil {
		return err
	}

	return nil
}

// ListViewAttributes defines attributes for ListView widget.
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
