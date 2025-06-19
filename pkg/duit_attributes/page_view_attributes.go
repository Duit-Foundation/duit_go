package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_builder"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_clip"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_flex"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_gestures"
)

type ListKind uint8

const (
	Common ListKind = iota
	Builder
)

type PageView[TInsets duit_edge_insets.EdgeInsets] struct {
	Type                   ListKind                        `json:"type"`
	ScrollDirection        duit_flex.Axis                  `json:"scrollDirection,omitempty"`
	Reverse                *bool                           `json:"reverse,omitempty"`
	Physics                duit_gestures.ScrollPhysics     `json:"physics,omitempty"`
	PageSnapping           *bool                           `json:"pageSnapping,omitempty"`
	DragStartBehavior      duit_gestures.DragStartBehavior `json:"dragStartBehavior,omitempty"`
	AllowImplicitScrolling *bool                           `json:"allowImplicitScrolling,omitempty"`
	RestorationId          string                          `json:"restorationId,omitempty"`
	ClipBehavior           duit_clip.Clip                  `json:"clipBehavior,omitempty"`
	HitTestBehavior        duit_gestures.HitTestBehavior   `json:"hitTestBehavior,omitempty"`
	ScrollBehavior         duit_gestures.ScrollBehavior    `json:"scrollBehavior,omitempty"`
	PadEnds                *bool                           `json:"padEnds,omitempty"`
}

type ListViewBuilderAttributes[TInsets duit_edge_insets.EdgeInsets] struct {
	ListView[TInsets]
	duit_builder.Builder
}

type PageViewAttributes[TInsets duit_edge_insets.EdgeInsets] interface {
	PageView[TInsets] | ListViewBuilderAttributes[TInsets]
}
