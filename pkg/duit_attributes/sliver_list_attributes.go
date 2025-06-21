package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_builder"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
)

type SliverListKind uint8

const (
	SliverListCommon SliverListKind = iota
	SliverListBuilder
	SliverListSeparated
)

type SliverList struct {
	Type                   SliverListKind `json:"type"`
	AddAutomaticKeepAlives *bool           `json:"addAutomaticKeepAlives,omitempty"`
	AddRepaintBoundaries   *bool           `json:"addRepaintBoundaries,omitempty"`
	AddSemanticIndexes     *bool           `json:"addSemanticIndexes,omitempty"`
}

type SliverListBuilderAttributes struct {
	SliverList
	duit_builder.Builder
}

type SliverListSeparatedAttributes struct {
	SliverListBuilderAttributes
	Separator *duit_core.DuitElementModel `json:"separator"`
}

type SliverListAttributes interface {
	SliverList | SliverListBuilderAttributes | SliverListSeparatedAttributes
} 