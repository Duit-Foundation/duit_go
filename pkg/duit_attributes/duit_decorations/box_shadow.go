package duit_decoration

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_color"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_flex"
)

type BoxShadow[T duit_color.Color] struct {
	Color        T                    `json:"color,omitempty"`
	BlurRadius   float32              `json:"blurRadius,omitempty"`
	SpreadRadius float32              `json:"spreadRadius,omitempty"`
	BlurStyle    duit_color.BlurStyle `json:"blurStyle,omitempty"`
	Offset       duit_flex.Offset     `json:"offset,omitempty"`
}
