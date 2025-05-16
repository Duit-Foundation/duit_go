package duit_attributes

import "github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"

type SliverFillViewportAttributes struct {
	ValueReferenceHolder
	ThemeConsumer
	ViewportFraction       float32                       `json:"viewportFraction,omitempty"`
	PadEnds                *bool                         `json:"padEnds,omitempty"`
	IsBuilderDelegate      *bool                         `json:"isBuilderDelegate,omitempty"`
	AddAutomaticKeepAlives *bool                         `json:"addAutomaticKeepAlives,omitempty"`
	AddRepaintBoundaries   *bool                         `json:"addRepaintBoundaries,omitempty"`
	AddSemanticIndexes     *bool                         `json:"addSemanticIndexes,omitempty"`
	ChildCount             int                           `json:"childCount,omitempty"`
	ChildObjects           []*duit_core.DuitElementModel `json:"childObjects,omitempty"`
}
