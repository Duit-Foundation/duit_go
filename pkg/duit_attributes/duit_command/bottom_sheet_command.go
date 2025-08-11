package duit_command

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_clip"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_color"
	duit_decoration "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_flex"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type BottomSheetCommand[TAction duit_core.Action, TColor duit_color.Color, TShape duit_decoration.ShapeBorder[TColor]] struct {
	Type                                string                              `json:"type"`
	Action                              OverlayAction                       `json:"action"`
	OnClose                             *TAction                            `json:"onClose,omitempty"`
	Content                             *duit_core.DuitElementModel         `json:"content"`
	IsDismissible                       duit_utils.Tristate[bool]           `json:"isDismissible,omitempty"`
	IsScrollControlled                  duit_utils.Tristate[bool]           `json:"isScrollControlled,omitempty"`
	UseSafeArea                         duit_utils.Tristate[bool]           `json:"useSafeArea,omitempty"`
	UseRootNavigator                    duit_utils.Tristate[bool]           `json:"useRootNavigator,omitempty"`
	EnableDrag                          duit_utils.Tristate[bool]           `json:"enableDrag,omitempty"`
	ShowDragHandle                      duit_utils.Tristate[bool]           `json:"showDragHandle,omitempty"`
	ScrollControlDisabledMaxHeightRatio duit_utils.Tristate[float32]        `json:"scrollControlDisabledMaxHeightRatio,omitempty"`
	Constraints                         *duit_flex.BoxConstraints            `json:"constraints,omitempty"`
	AnchorPoint                         *duit_flex.Offset                   `json:"anchorPoint,omitempty"`
	BackgroundColor                     TColor                              `json:"backgroundColor,omitempty"`
	BarrierColor                        TColor                              `json:"barrierColor,omitempty"`
	Shape                               *TShape                             `json:"shape,omitempty"`
	ClipBehavior                        duit_utils.Tristate[duit_clip.Clip] `json:"clipBehavior,omitempty"`
}

type BottomSheetUIProps[TColor duit_color.Color, TShape duit_decoration.ShapeBorder[TColor]] struct {
	IsDismissible                       duit_utils.Tristate[bool]           `json:"isDismissible,omitempty"`
	IsScrollControlled                  duit_utils.Tristate[bool]           `json:"isScrollControlled,omitempty"`
	UseSafeArea                         duit_utils.Tristate[bool]           `json:"useSafeArea,omitempty"`
	UseRootNavigator                    duit_utils.Tristate[bool]           `json:"useRootNavigator,omitempty"`
	EnableDrag                          duit_utils.Tristate[bool]           `json:"enableDrag,omitempty"`
	ShowDragHandle                      duit_utils.Tristate[bool]           `json:"showDragHandle,omitempty"`
	ScrollControlDisabledMaxHeightRatio duit_utils.Tristate[float32]        `json:"scrollControlDisabledMaxHeightRatio,omitempty"`
	Constraints                         *duit_flex.BoxConstraints            `json:"constraints,omitempty"`
	AnchorPoint                         *duit_flex.Offset                   `json:"anchorPoint,omitempty"`
	BackgroundColor                     TColor        `json:"backgroundColor,omitempty"`
	BarrierColor                        TColor         `json:"barrierColor,omitempty"`
	Shape                               *TShape                             `json:"shape,omitempty"`
	ClipBehavior                        duit_utils.Tristate[duit_clip.Clip] `json:"clipBehavior,omitempty"`
}

func NewBottomSheetCommand[TAction duit_core.Action, TColor duit_color.Color, TShape duit_decoration.ShapeBorder[TColor]](
	commandProps *OverlayCommandProps[TAction],
	uiProps *BottomSheetUIProps[TColor, TShape],
) *BottomSheetCommand[TAction, TColor, TShape] {
	return &BottomSheetCommand[TAction, TColor, TShape]{
		Type:                                "bottomSheet",
		Content:                             commandProps.Content,
		Action:                              commandProps.Action,
		OnClose:                             commandProps.OnClose,
		IsDismissible:                       uiProps.IsDismissible,
		IsScrollControlled:                  uiProps.IsScrollControlled,
		UseSafeArea:                         uiProps.UseSafeArea,
		UseRootNavigator:                    uiProps.UseRootNavigator,
		EnableDrag:                          uiProps.EnableDrag,
		ShowDragHandle:                      uiProps.ShowDragHandle,
		ScrollControlDisabledMaxHeightRatio: uiProps.ScrollControlDisabledMaxHeightRatio,
		Constraints:                         uiProps.Constraints,
		AnchorPoint:                         uiProps.AnchorPoint,
		BackgroundColor:                     uiProps.BackgroundColor,
		BarrierColor:                        uiProps.BarrierColor,
		Shape:                               uiProps.Shape,
		ClipBehavior:                        uiProps.ClipBehavior,
	}
}
