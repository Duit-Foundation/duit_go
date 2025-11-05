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

// SetAction sets the action for the overlay command props.
func (r *OverlayCommandProps) SetAction(a OverlayAction) *OverlayCommandProps {
	r.Action = a
	return r
}

// SetOnClose sets the on close for the overlay command props.
func (r *OverlayCommandProps) SetOnClose(o any) *OverlayCommandProps {
	r.OnClose = o
	return r
}

// SetContent sets the content for the overlay command props.
func (r *OverlayCommandProps) SetContent(c *duit_core.DuitElementModel) *OverlayCommandProps {
	r.Content = c
	return r
}