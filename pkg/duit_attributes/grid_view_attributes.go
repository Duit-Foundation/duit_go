package duit_attributes

import (
	"encoding/json"
	"errors"

	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_builder"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_clip"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_flex"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_gestures"
)

type GridConstructor uint

const (
	GridCommon GridConstructor = iota
	GridCount
	GridBuilder
	GridExtent
)

type DefautlGridViewAttributes[TInsets duit_edge_insets.EdgeInsets] struct {
	Constructor             GridConstructor                                 `json:"constructor"`
	ScrollDirection         duit_flex.Axis                                  `json:"scrollDirection,omitempty"`
	Reverse                 *bool                                           `json:"reverse,omitempty"`
	Primary                 *bool                                           `json:"primary,omitempty"`
	ShrinkWrap              *bool                                           `json:"shrinkWrap,omitempty"`
	Padding                 TInsets                                         `json:"padding,omitempty"`
	AddAutomaticKeepAlives  *bool                                           `json:"addAutomaticKeepAlives,omitempty"`
	AddRepaintBoundaries    *bool                                           `json:"addRepaintBoundaries,omitempty"`
	AddSemanticIndexes      *bool                                           `json:"addSemanticIndexes,omitempty"`
	CacheExtent             float32                                         `json:"cacheExtent,omitempty"`
	SemanticChildCount      int                                             `json:"semanticChildCount,omitempty"`
	DragStartBehavior       duit_gestures.DragStartBehavior                 `json:"dragStartBehavior,omitempty"`
	ClipBehavior            duit_clip.Clip                                  `json:"clipBehavior,omitempty"`
	KeyboardDismissBehavior duit_gestures.ScrollViewKeyboardDismissBehavior `json:"keyboardDismissBehavior,omitempty"`
	Physics                 duit_gestures.ScrollPhysics                     `json:"physics,omitempty"`
	RestorationId           string                                          `json:"restorationId,omitempty"`
}

type GridViewCommonAttributes[TInsets duit_edge_insets.EdgeInsets] struct {
	ValueReferenceHolder
	ThemeConsumer
	DefautlGridViewAttributes[TInsets]
	SliverGridDelegateKey     string                 `json:"sliverGridDelegateKey"`
	SliverGridDelegateOptions map[string]interface{} `json:"sliverGridDelegateOptions,omitempty"`
}

func (attributes *GridViewCommonAttributes[TInsets]) MarshalJSON() ([]byte, error) {
	attributes.Constructor = GridCommon
	
	if len(attributes.SliverGridDelegateKey) == 0 {
		return nil, errors.New("SliverGridDelegateKey is required")
	}

	return json.Marshal(*attributes)
}

type GridViewCountAttributes[TInsets duit_edge_insets.EdgeInsets] struct {
	ValueReferenceHolder
	ThemeConsumer
	DefautlGridViewAttributes[TInsets]
	CrossAxisCount   int     `json:"crossAxisCount"`
	MainAxisSpacing  float32 `json:"mainAxisSpacing,omitempty"`
	CrossAxisSpacing float32 `json:"crossAxisSpacing,omitempty"`
	ChildAspectRatio float32 `json:"childAspectRatio,omitempty"`
}

func (attributes *GridViewCountAttributes[TInsets]) MarshalJSON() ([]byte, error) {
	attributes.Constructor = GridCount

	if attributes.CrossAxisCount == 0 {
		return nil, errors.New("CrossAxisCount is required and must be greater than 0")
	}

	return json.Marshal(*attributes)
}

type GridViewBuilderAttributes[TInsets duit_edge_insets.EdgeInsets] struct {
	ValueReferenceHolder
	ThemeConsumer
	duit_builder.Builder
	DefautlGridViewAttributes[TInsets]
	SliverGridDelegateKey     string                 `json:"sliverGridDelegateKey"`
	SliverGridDelegateOptions map[string]interface{} `json:"sliverGridDelegateOptions,omitempty"`
}

func (attributes *GridViewBuilderAttributes[TInsets]) MarshalJSON() ([]byte, error) {
	attributes.Constructor = GridBuilder

	if len(attributes.SliverGridDelegateKey) == 0 {
		return nil, errors.New("SliverGridDelegateKey is required")
	}

	return json.Marshal(*attributes)
}

type GridViewExtentAttributes[TInsets duit_edge_insets.EdgeInsets] struct {
	ValueReferenceHolder
	ThemeConsumer
	DefautlGridViewAttributes[TInsets]
	MaxCrossAxisExtent float32 `json:"maxCrossAxisExtent"`
	MainAxisSpacing    float32 `json:"mainAxisSpacing,omitempty"`
	CrossAxisSpacing   float32 `json:"crossAxisSpacing,omitempty"`
	ChildAspectRatio   float32 `json:"childAspectRatio,omitempty"`
}

func (attributes *GridViewExtentAttributes[TInsets]) MarshalJSON() ([]byte, error) {
	attributes.Constructor = GridExtent

	if attributes.MaxCrossAxisExtent == 0 {
		return nil, errors.New("MaxCrossAxisExtent is required and must be greater than 0")
	}

	return json.Marshal(*attributes)
}

type GridViewAttributes[TInsets duit_edge_insets.EdgeInsets] interface {
	GridViewCommonAttributes[TInsets] | GridViewCountAttributes[TInsets] | GridViewBuilderAttributes[TInsets] | GridViewExtentAttributes[TInsets]
}
