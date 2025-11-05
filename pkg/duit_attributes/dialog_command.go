package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type DialogCommand struct {
	*OverlayCommandProps `gen:"wrap"`
	BarrierDismissible   duit_utils.Tristate[bool] `json:"barrierDismissible,omitempty"`
	UseSafeArea          duit_utils.Tristate[bool] `json:"useSafeArea,omitempty"`
	UseRootNavigator     duit_utils.Tristate[bool] `json:"useRootNavigator,omitempty"`
	BarrierColor         *duit_props.Color         `json:"barrierColor,omitempty"`
	BarrierLabel         string                    `json:"barrierLabel,omitempty"`
	AnchorPoint          *duit_props.Offset        `json:"anchorPoint,omitempty"`
}

// NewDialogCommand creates a new DialogCommand instance.
func NewDialogCommand() *DialogCommand {
	return &DialogCommand{
		OverlayCommandProps: &OverlayCommandProps{
			Type: "dialog",
		},
	}
}

// SetBarrierDismissible sets the barrier dismissible for the dialog command.
func (r *DialogCommand) SetBarrierDismissible(barrierDismissible duit_utils.Tristate[bool]) *DialogCommand {
	r.BarrierDismissible = barrierDismissible
	return r
}

// SetUseSafeArea sets the use safe area for the dialog command.
func (r *DialogCommand) SetUseSafeArea(useSafeArea duit_utils.Tristate[bool]) *DialogCommand {
	r.UseSafeArea = useSafeArea
	return r
}

// SetUseRootNavigator sets the use root navigator for the dialog command.
func (r *DialogCommand) SetUseRootNavigator(useRootNavigator duit_utils.Tristate[bool]) *DialogCommand {
	r.UseRootNavigator = useRootNavigator
	return r
}

// SetBarrierColor sets the barrier color for the dialog command.
func (r *DialogCommand) SetBarrierColor(barrierColor *duit_props.Color) *DialogCommand {
	r.BarrierColor = barrierColor
	return r
}

// SetBarrierLabel sets the barrier label for the dialog command.
func (r *DialogCommand) SetBarrierLabel(barrierLabel string) *DialogCommand {
	r.BarrierLabel = barrierLabel
	return r
}

// SetAnchorPoint sets the anchor point for the dialog command.
func (r *DialogCommand) SetAnchorPoint(anchorPoint *duit_props.Offset) *DialogCommand {
	r.AnchorPoint = anchorPoint
	return r
}
