package duit_command

import "github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"

type OverlayCommandProps[TAction duit_core.Action] struct {
	Type    string                      `json:"type"`
	Action  OverlayAction               `json:"action"`
	OnClose *TAction                    `json:"onClose,omitempty"`
	Content *duit_core.DuitElementModel `json:"content"`
}