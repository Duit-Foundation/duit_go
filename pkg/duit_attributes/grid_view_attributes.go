package duit_attributes

import (
	"encoding/json"
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type GridConstructor uint

const (
	GridCommon GridConstructor = iota
	GridCount
	GridBuilder
	GridExtent
)

// DefautlGridViewAttributes defines attributes for DefaultGridView widget.
type DefautlGridViewAttributes struct {
	Constructor             duit_utils.Tristate[GridConstructor]         `json:"constructor"`
	ScrollDirection         duit_props.Axis                              `json:"scrollDirection,omitempty"`
	Reverse                 duit_utils.Tristate[bool]                    `json:"reverse,omitempty"`
	Primary                 duit_utils.Tristate[bool]                    `json:"primary,omitempty"`
	ShrinkWrap              duit_utils.Tristate[bool]                    `json:"shrinkWrap,omitempty"`
	AddAutomaticKeepAlives  duit_utils.Tristate[bool]                    `json:"addAutomaticKeepAlives,omitempty"`
	AddRepaintBoundaries    duit_utils.Tristate[bool]                    `json:"addRepaintBoundaries,omitempty"`
	AddSemanticIndexes      duit_utils.Tristate[bool]                    `json:"addSemanticIndexes,omitempty"`
	Padding                 *duit_props.EdgeInsets                       `json:"padding,omitempty"`
	CacheExtent             float32                                      `json:"cacheExtent,omitempty"`
	SemanticChildCount      int                                          `json:"semanticChildCount,omitempty"`
	DragStartBehavior       duit_props.DragStartBehavior                 `json:"dragStartBehavior,omitempty"`
	ClipBehavior            duit_props.Clip                              `json:"clipBehavior,omitempty"`
	KeyboardDismissBehavior duit_props.ScrollViewKeyboardDismissBehavior `json:"keyboardDismissBehavior,omitempty"`
	Physics                 duit_props.ScrollPhysics                     `json:"physics,omitempty"`
	RestorationId           string                                       `json:"restorationId,omitempty"`
}

// NewDefautlGridViewAttributes creates a new DefautlGridViewAttributes instance.
func (r *DefautlGridViewAttributes) SetConstructor(constructor GridConstructor) {
	r.Constructor = duit_utils.TristateFrom[GridConstructor](constructor)
}

// SetScrollDirection sets the scroll direction for the DefaultGridViewAttributes instance.
func (r *DefautlGridViewAttributes) SetScrollDirection(scrollDirection duit_props.Axis) {
	r.ScrollDirection = scrollDirection
}

// SetReverse sets the reverse for the DefaultGridViewAttributes instance.
func (r *DefautlGridViewAttributes) SetReverse(reverse bool) {
	r.Reverse = duit_utils.TristateFrom[bool](reverse)
}

// SetPrimary sets the primary for the DefaultGridViewAttributes instance.
func (r *DefautlGridViewAttributes) SetPrimary(primary bool) {
	r.Primary = duit_utils.TristateFrom[bool](primary)
}

// SetShrinkWrap sets the shrink wrap for the DefaultGridViewAttributes instance.
func (r *DefautlGridViewAttributes) SetShrinkWrap(shrinkWrap bool) {
	r.ShrinkWrap = duit_utils.TristateFrom[bool](shrinkWrap)
}

// SetAddAutomaticKeepAlives sets the add automatic keep alive for the DefaultGridViewAttributes instance.
func (r *DefautlGridViewAttributes) SetAddAutomaticKeepAlives(addAutomaticKeepAlives bool) {
	r.AddAutomaticKeepAlives = duit_utils.TristateFrom[bool](addAutomaticKeepAlives)
}

// SetAddRepaintBoundaries sets the add repaint boundaries for the DefaultGridViewAttributes instance.
func (r *DefautlGridViewAttributes) SetAddRepaintBoundaries(addRepaintBoundaries bool) {
	r.AddRepaintBoundaries = duit_utils.TristateFrom[bool](addRepaintBoundaries)
}

// SetAddSemanticIndexes sets the add semantic indexes for the DefaultGridViewAttributes instance.
func (r *DefautlGridViewAttributes) SetAddSemanticIndexes(addSemanticIndexes bool) {
	r.AddSemanticIndexes = duit_utils.TristateFrom[bool](addSemanticIndexes)
}

// SetPadding sets the padding for the DefaultGridViewAttributes instance.
func (r *DefautlGridViewAttributes) SetPadding(padding *duit_props.EdgeInsets) {
	r.Padding = padding
}

// SetCacheExtent sets the cache extent for the DefaultGridViewAttributes instance.
func (r *DefautlGridViewAttributes) SetCacheExtent(cacheExtent float32) {
	r.CacheExtent = cacheExtent
}

// SetSemanticChildCount sets the semantic child count for the DefaultGridViewAttributes instance.
func (r *DefautlGridViewAttributes) SetSemanticChildCount(semanticChildCount int) {
	r.SemanticChildCount = semanticChildCount
}

// SetDragStartBehavior sets the drag start behavior for the DefaultGridViewAttributes instance.
func (r *DefautlGridViewAttributes) SetDragStartBehavior(dragStartBehavior duit_props.DragStartBehavior) {
	r.DragStartBehavior = dragStartBehavior
}

// SetClipBehavior sets the clip behavior for the DefaultGridViewAttributes instance.
func (r *DefautlGridViewAttributes) SetClipBehavior(clipBehavior duit_props.Clip) {
	r.ClipBehavior = clipBehavior
}

// SetKeyboardDismissBehavior sets the keyboard dismiss behavior for the DefaultGridViewAttributes instance.
func (r *DefautlGridViewAttributes) SetKeyboardDismissBehavior(keyboardDismissBehavior duit_props.ScrollViewKeyboardDismissBehavior) {
	r.KeyboardDismissBehavior = keyboardDismissBehavior
}

// SetRestorationId sets the restoration id for the DefaultGridViewAttributes instance.
func (r *DefautlGridViewAttributes) SetRestorationId(restorationId string) {
	r.RestorationId = restorationId
}

// SetPhysics sets the physics for the DefaultGridViewAttributes instance.
func (r *DefautlGridViewAttributes) SetPhysics(physics duit_props.ScrollPhysics) {
	r.Physics = physics
}

// GridViewCommonAttributes defines attributes for GridViewCommon widget.
// See: https://api.flutter.dev/flutter/widgets/GridView/GridView.html
type GridViewCommonAttributes struct {
	*ValueReferenceHolder      `gen:"wrap"`
	*ThemeConsumer             `gen:"wrap"`
	*DefautlGridViewAttributes `gen:"wrap"`
	SliverGridDelegateKey      string         `json:"sliverGridDelegateKey"`
	SliverGridDelegateOptions  map[string]any `json:"sliverGridDelegateOptions,omitempty"`
}

// NewGridViewCommonAttributes creates a new GridViewCommonAttributes instance.
func NewGridViewCommonAttributes(data *GridViewCommonAttributes) *GridViewAttributes {
	data.DefautlGridViewAttributes.Constructor = duit_utils.TristateFrom[GridConstructor](GridCommon)
	return &GridViewAttributes{data: data}
}

// SetSliverGridDelegateKey sets the sliver grid delegate key for the GridViewCommonAttributes instance.
func (r *GridViewCommonAttributes) SetSliverGridDelegateKey(sliverGridDelegateKey string) *GridViewCommonAttributes {
	r.SliverGridDelegateKey = sliverGridDelegateKey
	return r
}

// SetSliverGridDelegateOptions sets the sliver grid delegate options for the GridViewCommonAttributes instance.
func (r *GridViewCommonAttributes) SetSliverGridDelegateOptions(sliverGridDelegateOptions map[string]any) *GridViewCommonAttributes {
	r.SliverGridDelegateOptions = sliverGridDelegateOptions
	return r
}

// Validate validates the GridViewCommonAttributes instance.
func (r *GridViewCommonAttributes) Validate() error {
	// Validate embedded components
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if r.DefautlGridViewAttributes != nil {
		if r.DefautlGridViewAttributes.Constructor == nil {
			return errors.New("DefautlGridViewAttributes.Constructor is required")
		}

		if *r.DefautlGridViewAttributes.Constructor != GridCommon {
			return errors.New("DefautlGridViewAttributes.Constructor must be GridCommon")
		}
	} else {
		return errors.New("DefautlGridViewAttributes is required")
	}

	// Validate grid delegate configuration
	if len(r.SliverGridDelegateKey) == 0 {
		return errors.New("SliverGridDelegateKey is required")
	}

	return nil
}

// GridViewCountAttributes defines attributes for GridViewCount widget.
// See: https://api.flutter.dev/flutter/widgets/GridView/GridView.count.html
type GridViewCountAttributes struct {
	*ValueReferenceHolder      `gen:"wrap"`
	*ThemeConsumer             `gen:"wrap"`
	*DefautlGridViewAttributes `gen:"wrap"`
	CrossAxisCount             uint    `json:"crossAxisCount"`
	MainAxisSpacing            float32 `json:"mainAxisSpacing,omitempty"`
	CrossAxisSpacing           float32 `json:"crossAxisSpacing,omitempty"`
	ChildAspectRatio           float32 `json:"childAspectRatio,omitempty"`
}

// NewGridViewCountAttributes creates a new GridViewCountAttributes instance.
func NewGridViewCountAttributes(data *GridViewCountAttributes) *GridViewAttributes {
	data.DefautlGridViewAttributes.Constructor = duit_utils.TristateFrom[GridConstructor](GridCount)
	return &GridViewAttributes{data: data}
}

// SetCrossAxisCount sets the cross axis count for the GridViewCountAttributes instance.
func (r *GridViewCountAttributes) SetCrossAxisCount(crossAxisCount uint) *GridViewCountAttributes {
	r.CrossAxisCount = crossAxisCount
	return r
}

// SetMainAxisSpacing sets the main axis spacing for the GridViewCountAttributes instance.
func (r *GridViewCountAttributes) SetMainAxisSpacing(mainAxisSpacing float32) *GridViewCountAttributes {
	r.MainAxisSpacing = mainAxisSpacing
	return r
}

// SetCrossAxisSpacing sets the cross axis spacing for the GridViewCountAttributes instance.
func (r *GridViewCountAttributes) SetCrossAxisSpacing(crossAxisSpacing float32) *GridViewCountAttributes {
	r.CrossAxisSpacing = crossAxisSpacing
	return r
}

// SetChildAspectRatio sets the child aspect ratio for the GridViewCountAttributes instance.
func (r *GridViewCountAttributes) SetChildAspectRatio(childAspectRatio float32) *GridViewCountAttributes {
	r.ChildAspectRatio = childAspectRatio
	return r
}

// Validate validates the GridViewCountAttributes instance.
func (r *GridViewCountAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if r.DefautlGridViewAttributes != nil {
		if r.DefautlGridViewAttributes.Constructor == nil {
			return errors.New("DefautlGridViewAttributes.Constructor is required")
		}

		if *r.DefautlGridViewAttributes.Constructor != GridCount {
			return errors.New("DefautlGridViewAttributes.Constructor must be GridCount")
		}
	} else {
		return errors.New("DefautlGridViewAttributes is required")
	}

	if r.CrossAxisCount == 0 {
		return errors.New("crossAxisCount is required and must be greater than 0")
	}

	return nil
}

// GridViewBuilderAttributes defines attributes for GridViewBuilder widget.
// See: https://api.flutter.dev/flutter/widgets/GridView/GridView.builder.html
type GridViewBuilderAttributes struct {
	*ValueReferenceHolder      `gen:"wrap"`
	*ThemeConsumer             `gen:"wrap"`
	*Builder                   `gen:"wrap"`
	*DefautlGridViewAttributes `gen:"wrap"`
	SliverGridDelegateKey      string         `json:"sliverGridDelegateKey"`
	SliverGridDelegateOptions  map[string]any `json:"sliverGridDelegateOptions,omitempty"`
}

// NewGridViewBuilderAttributes creates a new GridViewBuilderAttributes instance.
func NewGridViewBuilderAttributes(data *GridViewBuilderAttributes) *GridViewAttributes {
	data.DefautlGridViewAttributes.Constructor = duit_utils.TristateFrom[GridConstructor](GridBuilder)
	return &GridViewAttributes{data: data}
}

// SetSliverGridDelegateKey sets the sliver grid delegate key for the GridViewBuilderAttributes instance.
func (r *GridViewBuilderAttributes) SetSliverGridDelegateKey(sliverGridDelegateKey string) *GridViewBuilderAttributes {
	r.SliverGridDelegateKey = sliverGridDelegateKey
	return r
}

// SetSliverGridDelegateOptions sets the sliver grid delegate options for the GridViewBuilderAttributes instance.
func (r *GridViewBuilderAttributes) SetSliverGridDelegateOptions(sliverGridDelegateOptions map[string]any) *GridViewBuilderAttributes {
	r.SliverGridDelegateOptions = sliverGridDelegateOptions
	return r
}

// Validate validates the GridViewBuilderAttributes instance.
func (r *GridViewBuilderAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if r.DefautlGridViewAttributes != nil {
		if r.DefautlGridViewAttributes.Constructor == nil {
			return errors.New("DefautlGridViewAttributes.Constructor is required")
		}

		if *r.DefautlGridViewAttributes.Constructor != GridBuilder {
			return errors.New("DefautlGridViewAttributes.Constructor must be GridBuilder")
		}
	} else {
		return errors.New("DefautlGridViewAttributes is required")
	}

	if r.Builder == nil {
		return errors.New("builder is required")
	}

	if len(r.SliverGridDelegateKey) == 0 {
		return errors.New("sliverGridDelegateKey is required")
	}

	return nil
}

// GridViewExtentAttributes defines attributes for GridViewExtent widget.
// See: https://api.flutter.dev/flutter/widgets/GridView/GridView.extent.html
type GridViewExtentAttributes struct {
	*ValueReferenceHolder      `gen:"wrap"`
	*ThemeConsumer             `gen:"wrap"`
	*DefautlGridViewAttributes `gen:"wrap"`
	MaxCrossAxisExtent         float32 `json:"maxCrossAxisExtent"`
	MainAxisSpacing            float32 `json:"mainAxisSpacing,omitempty"`
	CrossAxisSpacing           float32 `json:"crossAxisSpacing,omitempty"`
	ChildAspectRatio           float32 `json:"childAspectRatio,omitempty"`
}

// NewGridViewExtentAttributes creates a new GridViewExtentAttributes instance.
func NewGridViewExtentAttributes(data *GridViewExtentAttributes) *GridViewAttributes {
	data.DefautlGridViewAttributes.Constructor = duit_utils.TristateFrom[GridConstructor](GridExtent)
	return &GridViewAttributes{data: data}
}

// SetMaxCrossAxisExtent sets the max cross axis extent for the GridViewExtentAttributes instance.
func (r *GridViewExtentAttributes) SetMaxCrossAxisExtent(maxCrossAxisExtent float32) *GridViewExtentAttributes {
	r.MaxCrossAxisExtent = maxCrossAxisExtent
	return r
}

// SetMainAxisSpacing sets the main axis spacing for the GridViewExtentAttributes instance.
func (r *GridViewExtentAttributes) SetMainAxisSpacing(mainAxisSpacing float32) *GridViewExtentAttributes {
	r.MainAxisSpacing = mainAxisSpacing
	return r
}

// SetCrossAxisSpacing sets the cross axis spacing for the GridViewExtentAttributes instance.
func (r *GridViewExtentAttributes) SetCrossAxisSpacing(crossAxisSpacing float32) *GridViewExtentAttributes {
	r.CrossAxisSpacing = crossAxisSpacing
	return r
}

// SetChildAspectRatio sets the child aspect ratio for the GridViewExtentAttributes instance.
func (r *GridViewExtentAttributes) SetChildAspectRatio(childAspectRatio float32) *GridViewExtentAttributes {
	r.ChildAspectRatio = childAspectRatio
	return r
}

// Validate validates the GridViewExtentAttributes instance.
func (r *GridViewExtentAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if r.DefautlGridViewAttributes != nil {
		if r.DefautlGridViewAttributes.Constructor == nil {
			return errors.New("DefautlGridViewAttributes.Constructor is required")
		}

		if *r.DefautlGridViewAttributes.Constructor != GridExtent {
			return errors.New("DefautlGridViewAttributes.Constructor must be GridExtent")
		}
	} else {
		return errors.New("DefautlGridViewAttributes is required")
	}

	if r.MaxCrossAxisExtent == 0 {
		return errors.New("MaxCrossAxisExtent is required and must be greater than 0")
	}

	return nil
}

// GridViewAttributes defines attributes for GridView widget.
// See: https://api.flutter.dev/flutter/widgets/GridView.html
type GridViewAttributes struct {
	data any
}

// NewGridViewAttributes creates a new GridViewAttributes instance.
func NewGridViewAttributes(data any) *GridViewAttributes {
	switch data := data.(type) {
	case *GridViewCommonAttributes:
		return NewGridViewCommonAttributes(data)
	case *GridViewCountAttributes:
		return NewGridViewCountAttributes(data)
	case *GridViewBuilderAttributes:
		return NewGridViewBuilderAttributes(data)
	case *GridViewExtentAttributes:
		return NewGridViewExtentAttributes(data)
	}

	return nil
}

//bindgen:exclude
func (r *GridViewAttributes) Validate() error {
	switch data := r.data.(type) {
	case *GridViewCommonAttributes:
		return data.Validate()
	case *GridViewCountAttributes:
		return data.Validate()
	case *GridViewBuilderAttributes:
		return data.Validate()
	case *GridViewExtentAttributes:
		return data.Validate()
	}
	return nil
}

func (r *GridViewAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.data)
}
