package duit_decoration

import "github.com/lesleysin/duit_go/pkg/duit_attributes/duit_color"

type BoxDecoration[T duit_color.Color] struct {
	Color        T                  `json:"color,omitempty"`
	Boder        *BorderSide[T]     `json:"border,omitempty"`
	BorderRadius float32            `json:"borderRadius,omitempty"`
	Shape        BoxShape           `json:"shape,omitempty"`
	BoxShadow    []BoxShadow[T]     `json:"boxShadow"`
	Gradient     *LinearGradient[T] `json:"gradient,omitempty"`
}
