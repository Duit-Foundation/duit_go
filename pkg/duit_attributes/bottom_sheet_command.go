package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type BottomSheetCommand struct {
	*OverlayCommandProps                `gen:"wrap"`
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

// NewBottomSheetCommand creates a new BottomSheetCommand instance.
func NewBottomSheetCommand() *BottomSheetCommand {
	return &BottomSheetCommand{
		OverlayCommandProps: &OverlayCommandProps{
			Type: "bottomSheet",
		},
	}
}

// SetIsDismissible sets the is dismissible for the bottom sheet command.
func (r *BottomSheetCommand) SetIsDismissible(isDismissible duit_utils.Tristate[bool]) *BottomSheetCommand {
	r.IsDismissible = isDismissible
	return r
}

// SetIsScrollControlled sets the is scroll controlled for the bottom sheet command.
func (r *BottomSheetCommand) SetIsScrollControlled(isScrollControlled duit_utils.Tristate[bool]) *BottomSheetCommand {
	r.IsScrollControlled = isScrollControlled
	return r
}

// SetUseSafeArea sets the use safe area for the bottom sheet command.
func (r *BottomSheetCommand) SetUseSafeArea(useSafeArea duit_utils.Tristate[bool]) *BottomSheetCommand {
	r.UseSafeArea = useSafeArea
	return r
}

// SetUseRootNavigator sets the use root navigator for the bottom sheet command.
func (r *BottomSheetCommand) SetUseRootNavigator(useRootNavigator duit_utils.Tristate[bool]) *BottomSheetCommand {
	r.UseRootNavigator = useRootNavigator
	return r
}

// SetEnableDrag sets the enable drag for the bottom sheet command.
func (r *BottomSheetCommand) SetEnableDrag(enableDrag duit_utils.Tristate[bool]) *BottomSheetCommand {
	r.EnableDrag = enableDrag
	return r
}

// SetShowDragHandle sets the show drag handle for the bottom sheet command.
func (r *BottomSheetCommand) SetShowDragHandle(showDragHandle duit_utils.Tristate[bool]) *BottomSheetCommand {
	r.ShowDragHandle = showDragHandle
	return r
}

// SetScrollControlDisabledMaxHeightRatio sets the scroll control disabled max height ratio for the bottom sheet command.
func (r *BottomSheetCommand) SetScrollControlDisabledMaxHeightRatio(scrollControlDisabledMaxHeightRatio duit_utils.Tristate[float32]) *BottomSheetCommand {
	r.ScrollControlDisabledMaxHeightRatio = scrollControlDisabledMaxHeightRatio
	return r
}

// SetConstraints sets the constraints for the bottom sheet command.
func (r *BottomSheetCommand) SetConstraints(constraints *duit_props.BoxConstraints) *BottomSheetCommand {
	r.Constraints = constraints
	return r
}

// SetAnchorPoint sets the anchor point for the bottom sheet command.
func (r *BottomSheetCommand) SetAnchorPoint(anchorPoint *duit_props.Offset) *BottomSheetCommand {
	r.AnchorPoint = anchorPoint
	return r
}

// SetBackgroundColor sets the background color for the bottom sheet command.
func (r *BottomSheetCommand) SetBackgroundColor(backgroundColor *duit_props.Color) *BottomSheetCommand {
	r.BackgroundColor = backgroundColor
	return r
}

// SetBarrierColor sets the barrier color for the bottom sheet command.
func (r *BottomSheetCommand) SetBarrierColor(barrierColor *duit_props.Color) *BottomSheetCommand {
	r.BarrierColor = barrierColor
	return r
}

// SetShape sets the shape for the bottom sheet command.
func (r *BottomSheetCommand) SetShape(shape *duit_props.ShapeBorder) *BottomSheetCommand {
	r.Shape = shape
	return r
}

// SetClipBehavior sets the clip behavior for the bottom sheet command.
func (r *BottomSheetCommand) SetClipBehavior(clipBehavior duit_utils.Tristate[duit_props.Clip]) *BottomSheetCommand {
	r.ClipBehavior = clipBehavior
	return r
}
