package duit_attributes

import "github.com/lesleysin/duit_go/pkg/duit_attributes/duit_clip"

type ElevatedButtonAttributes struct {
	Autofocus    bool           `json:"autofocus,omitempty"`
	ClipBehavior duit_clip.Clip `json:"clipBehavior,omitempty"`
}
