package duit_attributes

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_clip"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_edge_insets"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_flex"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_gestures"
	"github.com/lesleysin/duit_go/pkg/duit_core"
)

type ListKind uint8
type MergeStrategy uint8

const (
	Common ListKind = iota
	Builder
	Separated
)

const (
	AddToEnd MergeStrategy = iota
	AddToStart
	Replace
)

type ListView[TInsets duit_edge_insets.EdgeInsets] struct {
	Type                    ListKind                                        `json:"type"`
	ScrollPhysics           duit_gestures.ScrollPhysics                     `json:"scrollPhysics,omitempty"`
	Reverse                 bool                                            `json:"reverse,omitempty"`
	Primary                 bool                                            `json:"primary,omitempty"`
	ShrinkWrap              bool                                            `json:"shrinkWrap,omitempty"`
	AddAutomaticKeepAlives  bool                                            `json:"addAutomaticKeepAlives,omitempty"`
	AddRepaintBoundaries    bool                                            `json:"addRepaintBoundaries,omitempty"`
	AddSemanticIndexes      bool                                            `json:"addSemanticIndexes,omitempty"`
	ScrollDirection         duit_flex.Axis                                  `json:"scrollDirection,omitempty"`
	CacheExtent             float32                                         `json:"cacheExtent,omitempty"`
	Anchor                  float32                                         `json:"anchors,omitempty"`
	SemantickChildCount     int                                             `json:"semanticChildCount,omitempty"`
	Padding                 *TInsets                                        `json:"padding,omitempty"`
	ItemExtent              float32                                         `json:"itemExtent,omitempty"`
	ClipBehavior            duit_clip.Clip                                  `json:"clipBehavior,omitempty"`
	RestorationId           string                                          `json:"restorationId,omitempty"`
	DragStarnBehavior       duit_gestures.DragStartBehavior                 `json:"dragStartBehavior,omitempty"`
	KeyboardDismissBehavior duit_gestures.ScrollViewKeyboardDismissBehavior `json:"keyboardDismissBehavior,omitempty"`
}

type ListViewBuilderAttributes[TInsets duit_edge_insets.EdgeInsets] struct {
	ListView[TInsets]
	ChildObjects              []*duit_core.DuitElementModel `json:"children,omitempty"`
	ScrollEndReachedThreshold float32                       `json:"scrollEndReachedThreshold,omitempty"`
	MergeStrategy             MergeStrategy                 `json:"mergeStrategy,omitempty"`
}

type ListViewSeparatedAttributes[TInsets duit_edge_insets.EdgeInsets] struct {
	ListViewBuilderAttributes[TInsets]
	Separator *duit_core.DuitElementModel `json:"separator"`
}

type ListViewAttributes[TInsets duit_edge_insets.EdgeInsets] interface {
	ListView[TInsets] | ListViewBuilderAttributes[TInsets] | ListViewSeparatedAttributes[TInsets]
}
