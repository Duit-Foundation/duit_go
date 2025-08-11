package duit_attributes

import (
	"encoding/json"
	"fmt"

	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_clip"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_color"
	decorations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_text_properties"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type AppBarAttributes[TColor duit_color.Color, TInsets duit_edge_insets.EdgeInsets, TShape decorations.ShapeBorder[TColor]] struct {
	ValueReferenceHolder
	animations.AnimatedPropertyOwner
	ThemeConsumer
	Title                     *duit_core.DuitElementModel             `json:"title,omitempty"`
	Leading                   *duit_core.DuitElementModel             `json:"leading,omitempty"`
	Actions                   []*duit_core.DuitElementModel           `json:"actions,omitempty"`
	FlexibleSpace             *duit_core.DuitElementModel             `json:"flexibleSpace,omitempty"`
	Bottom                    *duit_core.DuitElementModel             `json:"bottom,omitempty"`
	TitleTextStyle            *duit_text_properties.TextStyle[TColor] `json:"titleTextStyle,omitempty"`
	ToolbarTextStyle          *duit_text_properties.TextStyle[TColor] `json:"toolbarTextStyle,omitempty"`
	ShadowColor               TColor                                  `json:"shadowColor,omitempty"`
	BackgroundColor           TColor                                  `json:"backgroundColor,omitempty"`
	ForegroundColor           TColor                                  `json:"foregroundColor,omitempty"`
	SurfaceTintColor          TColor                                  `json:"surfaceTintColor,omitempty"`
	ActionsPadding            TInsets                                 `json:"actionsPadding,omitempty"`
	Shape                     *TShape                                 `json:"shape,omitempty"`
	ClipBehavior              duit_clip.Clip                          `json:"clipBehavior,omitempty"`
	TitleSpacing              float32                                 `json:"titleSpacing,omitempty"`
	Elevation                 float32                                 `json:"elevation,omitempty"`
	ToolbarHeight             float32                                 `json:"toolbarHeight,omitempty"`
	LeadingWidth              float32                                 `json:"leadingWidth,omitempty"`
	ScrolledUnderElevation    float32                                 `json:"scrolledUnderElevation,omitempty"`
	BottomOpacity             float32                                 `json:"bottomOpacity,omitempty"`
	ToolbarOpacity            float32                                 `json:"toolbarOpacity,omitempty"`
	ForceMaterialTransparency *bool                                   `json:"forceMaterialTransparency,omitempty"`
	CenterTitle               *bool                                   `json:"centerTitle,omitempty"`
	AutomaticallyImplyLeading *bool                                   `json:"automaticallyImplyLeading,omitempty"`
	ExcludeHeaderSemantics    *bool                                   `json:"excludeHeaderSemantics,omitempty"`
	Primary                   *bool                                   `json:"primary,omitempty"`
}

func (a *AppBarAttributes[TColor, TInsets, TShape]) MarshalJSON() ([]byte, error) {
	if a.AutomaticallyImplyLeading == nil {
		var bPtr = duit_utils.BoolPtr(true)
		a.AutomaticallyImplyLeading = bPtr
	}

	if a.Primary == nil {
		var bPtr = duit_utils.BoolPtr(true)
		a.Primary = bPtr
	}

	if a.ToolbarOpacity < 0 || a.ToolbarOpacity > 1 {
		return nil, fmt.Errorf("toolbarOpacity must be between 0 and 1")
	}

	if a.BottomOpacity < 0 || a.BottomOpacity > 1 {
		return nil, fmt.Errorf("bottomOpacity must be between 0 and 1")
	}
	
	return json.Marshal(*a)
}
