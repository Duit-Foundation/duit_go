package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type FocusAction string

const (
	FocusActionNextFocus        FocusAction = "nextFocus"
	FocusActionPreviousFocus    FocusAction = "previousFocus"
	FocusActionFocusInDirection FocusAction = "focusInDirection"
	FocusActionRequestFocus     FocusAction = "requestFocus"
	FocusActionUnfocus          FocusAction = "unfocus"
)

type FocusCommand struct {
	CommandType string                                             `json:"type"`
	Action      FocusAction                                        `json:"action"`
	NodeId      duit_utils.TString                                 `json:"nodeId,omitempty"`
	Disposition duit_utils.Tristate[duit_props.UnfocusDisposition] `json:"unfocusDisposition,omitempty"`
	Direction   duit_utils.Tristate[duit_props.TraversalDirection] `json:"traversalDirection,omitempty"`
}

func newFocusCommand() *FocusCommand {
	return &FocusCommand{
		CommandType: "focusNode",
	}
}

func NewRequestFocusCommand(nodeId string) *FocusCommand {
	cmd := newFocusCommand()
	cmd.Action = FocusActionRequestFocus
	cmd.NodeId = duit_utils.StringValue(nodeId)
	return cmd
}

func NewNextFocusCommand() *FocusCommand {
	cmd := newFocusCommand()
	cmd.Action = FocusActionNextFocus
	return cmd
}

func NewPreviousFocusCommand() *FocusCommand {
	cmd := newFocusCommand()
	cmd.Action = FocusActionPreviousFocus
	return cmd
}

func NewFocusInDirectionCommand(dir *duit_props.TraversalDirection) *FocusCommand {
	cmd := newFocusCommand()
	cmd.Action = FocusActionFocusInDirection
	cmd.Direction = duit_utils.TristateFrom[duit_props.TraversalDirection](dir)
	return cmd
}

func NewUnfocusCommand(disposition *duit_props.UnfocusDisposition) *FocusCommand {
	cmd := newFocusCommand()
	cmd.Action = FocusActionUnfocus
	cmd.Disposition = duit_utils.TristateFrom[duit_props.UnfocusDisposition](disposition)
	return cmd
}
