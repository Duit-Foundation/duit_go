package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type BottomSheetCommand struct {
	Type                                string                               `json:"type"`
	Action                              OverlayAction                        `json:"action"`
	OnClose                             any                                  `json:"onClose,omitempty"`
	Content                             *duit_core.DuitElementModel          `json:"content"`
	IsDismissible                       duit_utils.Tristate[bool]            `json:"isDismissible,omitempty"`
	IsScrollControlled                  duit_utils.Tristate[bool]            `json:"isScrollControlled,omitempty"`
	UseSafeArea                         duit_utils.Tristate[bool]            `json:"useSafeArea,omitempty"`
	UseRootNavigator                    duit_utils.Tristate[bool]            `json:"useRootNavigator,omitempty"`
	EnableDrag                          duit_utils.Tristate[bool]            `json:"enableDrag,omitempty"`
	ShowDragHandle                      duit_utils.Tristate[bool]            `json:"showDragHandle,omitempty"`
	ScrollControlDisabledMaxHeightRatio duit_utils.Tristate[float32]         `json:"scrollControlDisabledMaxHeightRatio,omitempty"`
	Constraints                         *duit_props.BoxConstraints           `json:"constraints,omitempty"`
	AnchorPoint                         *duit_props.Offset                   `json:"anchorPoint,omitempty"`
	BackgroundColor                     *duit_props.Color                    `json:"backgroundColor,omitempty"`
	BarrierColor                        *duit_props.Color                    `json:"barrierColor,omitempty"`
	Shape                               *duit_props.ShapeBorder              `json:"shape,omitempty"`
	ClipBehavior                        duit_utils.Tristate[duit_props.Clip] `json:"clipBehavior,omitempty"`
}

type BottomSheetUIProps struct {
	IsDismissible                       duit_utils.Tristate[bool]            `json:"isDismissible,omitempty"`
	IsScrollControlled                  duit_utils.Tristate[bool]            `json:"isScrollControlled,omitempty"`
	UseSafeArea                         duit_utils.Tristate[bool]            `json:"useSafeArea,omitempty"`
	UseRootNavigator                    duit_utils.Tristate[bool]            `json:"useRootNavigator,omitempty"`
	EnableDrag                          duit_utils.Tristate[bool]            `json:"enableDrag,omitempty"`
	ShowDragHandle                      duit_utils.Tristate[bool]            `json:"showDragHandle,omitempty"`
	ScrollControlDisabledMaxHeightRatio duit_utils.Tristate[float32]         `json:"scrollControlDisabledMaxHeightRatio,omitempty"`
	Constraints                         *duit_props.BoxConstraints           `json:"constraints,omitempty"`
	AnchorPoint                         *duit_props.Offset                   `json:"anchorPoint,omitempty"`
	BackgroundColor                     *duit_props.Color                    `json:"backgroundColor,omitempty"`
	BarrierColor                        *duit_props.Color                    `json:"barrierColor,omitempty"`
	Shape                               *duit_props.ShapeBorder              `json:"shape,omitempty"`
	ClipBehavior                        duit_utils.Tristate[duit_props.Clip] `json:"clipBehavior,omitempty"`
}

func NewBottomSheetCommand(
	commandProps *OverlayCommandProps,
	uiProps *BottomSheetUIProps,
) *BottomSheetCommand {
	if commandProps.OnClose != nil {
		if err := duit_action.CheckActionType(commandProps.OnClose); err != nil {
			panic(err)
		}
	}

	return &BottomSheetCommand{
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
