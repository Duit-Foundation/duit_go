package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_flex"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_gestures"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type SingleChildScrollViewAttributes[TInsets duit_edge_insets.EdgeInsets] struct {
	*ValueReferenceHolder
	*ThemeConsumer
	ScrollDirection         duit_flex.Axis                                  `json:"scrollDirection,omitempty"`
	Reverse                 duit_utils.Tristate[bool]                       `json:"reverse,omitempty"`
	Primary                 duit_utils.Tristate[bool]                       `json:"primary,omitempty"`
	Padding                 TInsets                                         `json:"padding,omitempty"`
	RestorationId           string                                          `json:"restorationId,omitempty"`
	ClipBehavior            duit_props.Clip                                  `json:"clipBehavior,omitempty"`
	DragStartBehavior       duit_gestures.DragStartBehavior                 `json:"dragStartBehavior,omitempty"`
	KeyboardDismissBehavior duit_gestures.ScrollViewKeyboardDismissBehavior `json:"keyboardDismissBehavior,omitempty"`
	Physics                 duit_gestures.ScrollPhysics                     `json:"physics,omitempty"`
}

func (r *SingleChildScrollViewAttributes[TInsets]) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	return nil
}
