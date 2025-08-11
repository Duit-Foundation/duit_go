package duit_attributes

import (
	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_color"
	duit_decoration "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_edge_insets"
)

type CardAttributes[TColor duit_color.Color, TInsets duit_edge_insets.EdgeInsets, TShape duit_decoration.ShapeBorder[TColor]] struct {
	ValueReferenceHolder
	animations.AnimatedPropertyOwner
	ThemeConsumer
	Color TColor `json:"color,omitempty"`
	ShadowColor TColor `json:"shadowColor,omitempty"`
	Elevation float32 `json:"elevation,omitempty"`
	BorderOnForeground *bool `json:"borderOnForeground,omitempty"`
	SemanticContainer *bool `json:"semanticContainer,omitempty"`
	Margin TInsets `json:"margin,omitempty"`
	Shape *TShape `json:"shape,omitempty"`
}