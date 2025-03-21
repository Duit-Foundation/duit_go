package duit_attributes

import "github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_clip"

type ElevatedButtonAttributes struct {
	ValueReferenceHolder
	ThemeConsumer
	Autofocus    bool           `json:"autofocus,omitempty"`
	ClipBehavior duit_clip.Clip `json:"clipBehavior,omitempty"`
}
