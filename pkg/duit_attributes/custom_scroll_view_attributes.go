package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_clip"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_flex"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_gestures"
)

type CustomScrollViewAttributes struct {
	ValueReferenceHolder
	ThemeConsumer
	Physics                 duit_gestures.ScrollPhysics                     `json:"physics,omitempty"`
	Reverse                 bool                                            `json:"reverse,omitempty"`
	Primary                 bool                                            `json:"primary,omitempty"`
	ShrinkWrap              bool                                            `json:"shrinkWrap,omitempty"`
	ScrollDirection         duit_flex.Axis                                  `json:"scrollDirection,omitempty"`
	Anchor                  float32                                         `json:"anchors,omitempty"`
	SemantickChildCount     int                                             `json:"semanticChildCount,omitempty"`
	ClipBehavior            duit_clip.Clip                                  `json:"clipBehavior,omitempty"`
	RestorationId           string                                          `json:"restorationId,omitempty"`
	DragStarnBehavior       duit_gestures.DragStartBehavior                 `json:"dragStartBehavior,omitempty"`
	KeyboardDismissBehavior duit_gestures.ScrollViewKeyboardDismissBehavior `json:"keyboardDismissBehavior,omitempty"`
	HitTestBehavior         duit_gestures.HitTestBehavior                   `json:"hitTestBehavior,omitempty"`
	CacheExtent             float32                                         `json:"cacheExtent,omitempty"`
	Center                  string                                          `json:"center,omitempty"`
}
