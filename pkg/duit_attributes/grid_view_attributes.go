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

func (r *DefautlGridViewAttributes) SetConstructor(constructor GridConstructor) {
	r.Constructor = duit_utils.TristateFrom[GridConstructor](constructor)
}

func (r *DefautlGridViewAttributes) SetScrollDirection(scrollDirection duit_props.Axis) {
	r.ScrollDirection = scrollDirection
}

func (r *DefautlGridViewAttributes) SetReverse(reverse bool) {
	r.Reverse = duit_utils.TristateFrom[bool](reverse)
}

func (r *DefautlGridViewAttributes) SetPrimary(primary bool) {
	r.Primary = duit_utils.TristateFrom[bool](primary)
}

func (r *DefautlGridViewAttributes) SetShrinkWrap(shrinkWrap bool) {
	r.ShrinkWrap = duit_utils.TristateFrom[bool](shrinkWrap)
}

func (r *DefautlGridViewAttributes) SetAddAutomaticKeepAlives(addAutomaticKeepAlives bool) {
	r.AddAutomaticKeepAlives = duit_utils.TristateFrom[bool](addAutomaticKeepAlives)
}

func (r *DefautlGridViewAttributes) SetAddRepaintBoundaries(addRepaintBoundaries bool) {
	r.AddRepaintBoundaries = duit_utils.TristateFrom[bool](addRepaintBoundaries)
}

func (r *DefautlGridViewAttributes) SetAddSemanticIndexes(addSemanticIndexes bool) {
	r.AddSemanticIndexes = duit_utils.TristateFrom[bool](addSemanticIndexes)
}

func (r *DefautlGridViewAttributes) SetPadding(padding *duit_props.EdgeInsets) {
	r.Padding = padding
}

func (r *DefautlGridViewAttributes) SetCacheExtent(cacheExtent float32) {
	r.CacheExtent = cacheExtent
}

func (r *DefautlGridViewAttributes) SetSemanticChildCount(semanticChildCount int) {
	r.SemanticChildCount = semanticChildCount
}

func (r *DefautlGridViewAttributes) SetDragStartBehavior(dragStartBehavior duit_props.DragStartBehavior) {
	r.DragStartBehavior = dragStartBehavior
}

func (r *DefautlGridViewAttributes) SetClipBehavior(clipBehavior duit_props.Clip) {
	r.ClipBehavior = clipBehavior
}

func (r *DefautlGridViewAttributes) SetKeyboardDismissBehavior(keyboardDismissBehavior duit_props.ScrollViewKeyboardDismissBehavior) {
	r.KeyboardDismissBehavior = keyboardDismissBehavior
}

func (r *DefautlGridViewAttributes) SetRestorationId(restorationId string) {
	r.RestorationId = restorationId
}

func (r *DefautlGridViewAttributes) SetPhysics(physics duit_props.ScrollPhysics) {
	r.Physics = physics
}

type GridViewCommonAttributes struct {
	*ValueReferenceHolder
	*ThemeConsumer
	*DefautlGridViewAttributes
	SliverGridDelegateKey     string         `json:"sliverGridDelegateKey"`
	SliverGridDelegateOptions map[string]any `json:"sliverGridDelegateOptions,omitempty"`
}

func NewGridViewCommonAttributes(data *GridViewCommonAttributes) *GridViewAttributes {
	data.DefautlGridViewAttributes.Constructor = duit_utils.TristateFrom[GridConstructor](GridCommon)
	return &GridViewAttributes{data: data}
}

func (r *GridViewCommonAttributes) SetSliverGridDelegateKey(sliverGridDelegateKey string) *GridViewCommonAttributes {
	r.SliverGridDelegateKey = sliverGridDelegateKey
	return r
}

func (r *GridViewCommonAttributes) SetSliverGridDelegateOptions(sliverGridDelegateOptions map[string]any) *GridViewCommonAttributes {
	r.SliverGridDelegateOptions = sliverGridDelegateOptions
	return r
}

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

type GridViewCountAttributes struct {
	*ValueReferenceHolder
	*ThemeConsumer
	*DefautlGridViewAttributes
	CrossAxisCount   uint    `json:"crossAxisCount"`
	MainAxisSpacing  float32 `json:"mainAxisSpacing,omitempty"`
	CrossAxisSpacing float32 `json:"crossAxisSpacing,omitempty"`
	ChildAspectRatio float32 `json:"childAspectRatio,omitempty"`
}

func NewGridViewCountAttributes(data *GridViewCountAttributes) *GridViewAttributes {
	data.DefautlGridViewAttributes.Constructor = duit_utils.TristateFrom[GridConstructor](GridCount)
	return &GridViewAttributes{data: data}
}

func (r *GridViewCountAttributes) SetCrossAxisCount(crossAxisCount uint) *GridViewCountAttributes {
	r.CrossAxisCount = crossAxisCount
	return r
}

func (r *GridViewCountAttributes) SetMainAxisSpacing(mainAxisSpacing float32) *GridViewCountAttributes {
	r.MainAxisSpacing = mainAxisSpacing
	return r
}

func (r *GridViewCountAttributes) SetCrossAxisSpacing(crossAxisSpacing float32) *GridViewCountAttributes {
	r.CrossAxisSpacing = crossAxisSpacing
	return r
}

func (r *GridViewCountAttributes) SetChildAspectRatio(childAspectRatio float32) *GridViewCountAttributes {
	r.ChildAspectRatio = childAspectRatio
	return r
}

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

type GridViewBuilderAttributes struct {
	*ValueReferenceHolder
	*ThemeConsumer
	*Builder
	*DefautlGridViewAttributes
	SliverGridDelegateKey     string         `json:"sliverGridDelegateKey"`
	SliverGridDelegateOptions map[string]any `json:"sliverGridDelegateOptions,omitempty"`
}

func NewGridViewBuilderAttributes(data *GridViewBuilderAttributes) *GridViewAttributes {
	data.DefautlGridViewAttributes.Constructor = duit_utils.TristateFrom[GridConstructor](GridBuilder)
	return &GridViewAttributes{data: data}
}

func (r *GridViewBuilderAttributes) SetSliverGridDelegateKey(sliverGridDelegateKey string) *GridViewBuilderAttributes {
	r.SliverGridDelegateKey = sliverGridDelegateKey
	return r
}

func (r *GridViewBuilderAttributes) SetSliverGridDelegateOptions(sliverGridDelegateOptions map[string]any) *GridViewBuilderAttributes {
	r.SliverGridDelegateOptions = sliverGridDelegateOptions
	return r
}


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

type GridViewExtentAttributes struct {
	*ValueReferenceHolder
	*ThemeConsumer
	*DefautlGridViewAttributes
	MaxCrossAxisExtent float32 `json:"maxCrossAxisExtent"`
	MainAxisSpacing    float32 `json:"mainAxisSpacing,omitempty"`
	CrossAxisSpacing   float32 `json:"crossAxisSpacing,omitempty"`
	ChildAspectRatio   float32 `json:"childAspectRatio,omitempty"`
}

func NewGridViewExtentAttributes(data *GridViewExtentAttributes) *GridViewAttributes {
	data.DefautlGridViewAttributes.Constructor = duit_utils.TristateFrom[GridConstructor](GridExtent)
	return &GridViewAttributes{data: data}
}

func (r *GridViewExtentAttributes) SetMaxCrossAxisExtent(maxCrossAxisExtent float32) *GridViewExtentAttributes {
	r.MaxCrossAxisExtent = maxCrossAxisExtent
	return r
}

func (r *GridViewExtentAttributes) SetMainAxisSpacing(mainAxisSpacing float32) *GridViewExtentAttributes {
	r.MainAxisSpacing = mainAxisSpacing
	return r
}

func (r *GridViewExtentAttributes) SetCrossAxisSpacing(crossAxisSpacing float32) *GridViewExtentAttributes {
	r.CrossAxisSpacing = crossAxisSpacing
	return r
}

func (r *GridViewExtentAttributes) SetChildAspectRatio(childAspectRatio float32) *GridViewExtentAttributes {
	r.ChildAspectRatio = childAspectRatio
	return r
}

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

type GridViewAttributes struct {
	data any
}

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
