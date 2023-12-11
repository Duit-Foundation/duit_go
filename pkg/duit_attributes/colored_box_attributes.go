package duit_attributes

import "github.com/lesleysin/duit_go/pkg/duit_attributes/duit_color"

type ColoredBoxAttributes[T duit_color.Color] struct {
	Color T `json:"color,omitempty"`
}
