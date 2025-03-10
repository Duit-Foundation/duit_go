package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_clip"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_flex"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_gestures"
)

type SingleChildScrollViewAttributes[TInsets duit_edge_insets.EdgeInsets] struct {
	ValueReferenceHolder
	ThemeConsumer
	ScrollDirection         duit_flex.Axis                                  `json:"scrollDirection,omitempty"`
	Reverse                 bool                                            `json:"reverse,omitempty"`
	Primary                 bool                                            `json:"primary,omitempty"`
	Padding                 TInsets                                         `json:"padding,omitempty"`
	RestorationId           string                                          `json:"restorationId,omitempty"`
	ClipBehavior            duit_clip.Clip                                  `json:"clipBehavior,omitempty"`
	DragStartBehavior       duit_gestures.DragStartBehavior                 `json:"dragStartBehavior,omitempty"`
	KeyboardDismissBehavior duit_gestures.ScrollViewKeyboardDismissBehavior `json:"keyboardDismissBehavior,omitempty"`
	Physics                 duit_gestures.ScrollPhysics                     `json:"physics,omitempty"`
}
