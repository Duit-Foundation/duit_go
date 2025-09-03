package duit_props

import "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_flex"

type BoxShadow[T Color] struct {
	Color        T         `json:"color,omitempty"`
	BlurRadius   float32   `json:"blurRadius,omitempty"`
	SpreadRadius float32   `json:"spreadRadius,omitempty"`
	BlurStyle    BlurStyle `json:"blurStyle,omitempty"`
	Offset       duit_flex.Offset       `json:"offset,omitempty"`
}
