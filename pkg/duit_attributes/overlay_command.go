package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

type OverlayCommandProps struct {
	Type    string                      `json:"type"`
	Action  OverlayAction               `json:"action"`
	OnClose any                         `json:"onClose,omitempty"`
	Content *duit_core.DuitElementModel `json:"content"`
}
