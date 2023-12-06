package duit_attributes

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_alignment"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_clip"
	duit_decoration "github.com/lesleysin/duit_go/pkg/duit_attributes/duit_decorations"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_edge_insets"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_flex"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_utility_styles"
)

type ContainerAttributes[TInsets duit_edge_insets.EdgeInsets, TColor duit_utility_styles.Color] struct {
	Width                float32                               `json:"width,omitempty"`
	Height               float32                               `json:"height,omitempty"`
	Color                TColor                                `json:"color,omitempty"`
	ClipBehavior         duit_clip.Clip                        `json:"clipBehavior,omitempty"`
	Decoration           duit_decoration.BoxDecoration[TColor] `json:"decoration,omitempty"`
	ForegroundDecoration duit_decoration.BoxDecoration[TColor] `json:"foregroundDecoration,omitempty"`
	Padding              TInsets                               `json:"padding,omitempty"`
	Margin               TInsets                               `json:"margin,omitempty"`
	Alignment            duit_alignment.Alignment              `json:"alignment,omitempty"`
	TransformAlignment   duit_alignment.Alignment              `json:"transformAlignment,omitempty"`
	Constraints          duit_flex.BoxConstraits               `json:"constraints,omitempty"`
}
