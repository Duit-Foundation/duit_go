package duit_attributes

import (
	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_clip"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_color"
	duit_decoration "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_flex"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

type ContainerAttributes[TInsets duit_edge_insets.EdgeInsets, TColor duit_color.Color] struct {
	*ValueReferenceHolder
	*animations.AnimatedPropertyOwner
	*ThemeConsumer
	Width                float32                                `json:"width,omitempty"`
	Height               float32                                `json:"height,omitempty"`
	Color                TColor                                 `json:"color,omitempty"`
	ClipBehavior         duit_clip.Clip                         `json:"clipBehavior,omitempty"`
	Decoration           *duit_decoration.BoxDecoration[TColor] `json:"decoration,omitempty"`
	ForegroundDecoration *duit_decoration.BoxDecoration[TColor] `json:"foregroundDecoration,omitempty"`
	Padding              TInsets                                `json:"padding,omitempty"`
	Margin               TInsets                                `json:"margin,omitempty"`
	Alignment            duit_props.Alignment               `json:"alignment,omitempty"`
	TransformAlignment   duit_props.Alignment               `json:"transformAlignment,omitempty"`
	Constraints          *duit_flex.BoxConstraints              `json:"constraints,omitempty"`
}

func (r *ContainerAttributes[TInsets, TColor]) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	return nil
}
