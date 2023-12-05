package duit_attributes

import "github.com/lesleysin/duit_go/pkg/duit_attributes/duit_utility_styles"

type ColoredBoxAttributes[T duit_utility_styles.Color] struct {
	Color string `json:"color,omitempty"`
}
