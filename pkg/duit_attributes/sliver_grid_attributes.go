package duit_attributes

import (
	"encoding/json"
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_builder"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_clip"
)

type SliverGridConstructor uint

const (
	SliverGridCommon SliverGridConstructor = iota
	SliverGridCount
	SliverGridBuilder
	SliverGridExtent
)

type DefaultSliverGridAttributes struct {
	Constructor             SliverGridConstructor                           `json:"constructor"`
	AddAutomaticKeepAlives  *bool                                           `json:"addAutomaticKeepAlives,omitempty"`
	AddRepaintBoundaries    *bool                                           `json:"addRepaintBoundaries,omitempty"`
	AddSemanticIndexes      *bool                                           `json:"addSemanticIndexes,omitempty"`
	CacheExtent             float32                                         `json:"cacheExtent,omitempty"`
	SemanticChildCount      int                                             `json:"semanticChildCount,omitempty"`
	ClipBehavior            duit_clip.Clip                                  `json:"clipBehavior,omitempty"`
}

type SliverGridCommonAttributes struct {
	ValueReferenceHolder
	ThemeConsumer
	DefaultSliverGridAttributes
	SliverGridDelegateKey     string                 `json:"sliverGridDelegateKey"`
	SliverGridDelegateOptions map[string]interface{} `json:"sliverGridDelegateOptions,omitempty"`
}

func (attributes *SliverGridCommonAttributes) MarshalJSON() ([]byte, error) {
	attributes.Constructor = SliverGridCommon
	if len(attributes.SliverGridDelegateKey) == 0 {
		return nil, errors.New("SliverGridDelegateKey is required")
	}
	return json.Marshal(*attributes)
}

type SliverGridCountAttributes struct {
	ValueReferenceHolder
	ThemeConsumer
	DefaultSliverGridAttributes
	CrossAxisCount   int     `json:"crossAxisCount"`
	MainAxisSpacing  float32 `json:"mainAxisSpacing,omitempty"`
	CrossAxisSpacing float32 `json:"crossAxisSpacing,omitempty"`
	ChildAspectRatio float32 `json:"childAspectRatio,omitempty"`
}

func (attributes *SliverGridCountAttributes) MarshalJSON() ([]byte, error) {
	attributes.Constructor = SliverGridCount
	if attributes.CrossAxisCount == 0 {
		return nil, errors.New("CrossAxisCount is required and must be greater than 0")
	}
	return json.Marshal(*attributes)
}

type SliverGridBuilderAttributes struct {
	ValueReferenceHolder
	ThemeConsumer
	duit_builder.Builder
	DefaultSliverGridAttributes
	SliverGridDelegateKey     string                 `json:"sliverGridDelegateKey"`
	SliverGridDelegateOptions map[string]interface{} `json:"sliverGridDelegateOptions,omitempty"`
}

func (attributes *SliverGridBuilderAttributes) MarshalJSON() ([]byte, error) {
	attributes.Constructor = SliverGridBuilder
	if len(attributes.SliverGridDelegateKey) == 0 {
		return nil, errors.New("SliverGridDelegateKey is required")
	}
	return json.Marshal(*attributes)
}

type SliverGridExtentAttributes struct {
	ValueReferenceHolder
	ThemeConsumer
	DefaultSliverGridAttributes
	MaxCrossAxisExtent float32 `json:"maxCrossAxisExtent"`
	MainAxisSpacing    float32 `json:"mainAxisSpacing,omitempty"`
	CrossAxisSpacing   float32 `json:"crossAxisSpacing,omitempty"`
	ChildAspectRatio   float32 `json:"childAspectRatio,omitempty"`
}

func (attributes *SliverGridExtentAttributes) MarshalJSON() ([]byte, error) {
	attributes.Constructor = SliverGridExtent
	if attributes.MaxCrossAxisExtent == 0 {
		return nil, errors.New("MaxCrossAxisExtent is required and must be greater than 0")
	}
	return json.Marshal(*attributes)
}

type SliverGridAttributes interface {
	SliverGridCommonAttributes | SliverGridCountAttributes | SliverGridBuilderAttributes | SliverGridExtentAttributes
} 