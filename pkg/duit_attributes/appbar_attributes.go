package duit_attributes

import (
	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	decorations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_text_properties"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type AppBarAttributes[TColor duit_props.Color, TInsets duit_props.EdgeInsets, TShape decorations.ShapeBorder[TColor]] struct {
	*ValueReferenceHolder
	*animations.AnimatedPropertyOwner
	*ThemeConsumer
	TitleTextStyle            *duit_text_properties.TextStyle[TColor] `json:"titleTextStyle,omitempty"`
	ToolbarTextStyle          *duit_text_properties.TextStyle[TColor] `json:"toolbarTextStyle,omitempty"`
	ShadowColor               TColor                                  `json:"shadowColor,omitempty"`
	BackgroundColor           TColor                                  `json:"backgroundColor,omitempty"`
	ForegroundColor           TColor                                  `json:"foregroundColor,omitempty"`
	SurfaceTintColor          TColor                                  `json:"surfaceTintColor,omitempty"`
	ActionsPadding            TInsets                                 `json:"actionsPadding,omitempty"`
	Shape                     *TShape                                 `json:"shape,omitempty"`
	ClipBehavior              duit_props.Clip                          `json:"clipBehavior,omitempty"`
	TitleSpacing              float32                                 `json:"titleSpacing,omitempty"`
	Elevation                 float32                                 `json:"elevation,omitempty"`
	ToolbarHeight             float32                                 `json:"toolbarHeight,omitempty"`
	LeadingWidth              float32                                 `json:"leadingWidth,omitempty"`
	ScrolledUnderElevation    float32                                 `json:"scrolledUnderElevation,omitempty"`
	BottomOpacity             duit_utils.Tristate[float32]            `json:"bottomOpacity,omitempty"`
	ToolbarOpacity            duit_utils.Tristate[float32]            `json:"toolbarOpacity,omitempty"`
	ForceMaterialTransparency duit_utils.Tristate[bool]               `json:"forceMaterialTransparency,omitempty"`
	CenterTitle               duit_utils.Tristate[bool]               `json:"centerTitle,omitempty"`
	AutomaticallyImplyLeading duit_utils.Tristate[bool]               `json:"automaticallyImplyLeading,omitempty"`
	ExcludeHeaderSemantics    duit_utils.Tristate[bool]               `json:"excludeHeaderSemantics,omitempty"`
	Primary                   duit_utils.Tristate[bool]               `json:"primary,omitempty"`
}

func (r *AppBarAttributes[TColor, TInsets, TShape]) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if r.BottomOpacity != nil {
		if err := RangeValidator(*r.BottomOpacity, 0, 1); err != nil {
			return err
		}
	}

	if r.ToolbarOpacity != nil {
		if err := RangeValidator(*r.ToolbarOpacity, 0, 1); err != nil {
			return err
		}
	}

	return nil
}
