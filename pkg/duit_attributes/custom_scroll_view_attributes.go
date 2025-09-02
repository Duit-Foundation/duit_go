package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_flex"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_gestures"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type CustomScrollViewAttributes struct {
	*ValueReferenceHolder
	*ThemeConsumer
	Physics                 duit_gestures.ScrollPhysics                     `json:"physics,omitempty"`
	Reverse                 duit_utils.Tristate[bool]                       `json:"reverse,omitempty"`
	Primary                 duit_utils.Tristate[bool]                       `json:"primary,omitempty"`
	ShrinkWrap              duit_utils.Tristate[bool]                       `json:"shrinkWrap,omitempty"`
	ScrollDirection         duit_flex.Axis                                  `json:"scrollDirection,omitempty"`
	Anchor                  float32                                         `json:"anchors,omitempty"`
	SemantickChildCount     int                                             `json:"semanticChildCount,omitempty"`
	ClipBehavior            duit_props.Clip                                  `json:"clipBehavior,omitempty"`
	RestorationId           string                                          `json:"restorationId,omitempty"`
	DragStarnBehavior       duit_gestures.DragStartBehavior                 `json:"dragStartBehavior,omitempty"`
	KeyboardDismissBehavior duit_gestures.ScrollViewKeyboardDismissBehavior `json:"keyboardDismissBehavior,omitempty"`
	HitTestBehavior         duit_gestures.HitTestBehavior                   `json:"hitTestBehavior,omitempty"`
	CacheExtent             float32                                         `json:"cacheExtent,omitempty"`
	Center                  string                                          `json:"center,omitempty"`
}

func (r *CustomScrollViewAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	return nil
}
