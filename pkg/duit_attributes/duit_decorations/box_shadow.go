package duit_decoration

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_flex"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_utility_styles"
)

type BoxShadow[T duit_utility_styles.Color] struct {
	Color        T                             `json:"color,omitempty"`
	BlurRadius   float32                       `json:"blurRadius,omitempty"`
	SpreadRadius float32                       `json:"spreadRadius,omitempty"`
	BlurStyle    duit_utility_styles.BlurStyle `json:"blurStyle,omitempty"`
	Offset       duit_flex.Offset              `json:"offset,omitempty"`
}
