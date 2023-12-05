package duit_decoration

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_alignment"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_utility_styles"
)

type LinearGradient[T duit_utility_styles.Color] struct {
	Colors []T `json:"colors"`
	//Rotation angle in radians
	RotationAngle float32                  `json:"rotationAngle,omitempty"`
	Stops         []float32                `json:"stops,omitempty"`
	Begin         duit_alignment.Alignment `json:"begin,omitempty"`
	End           duit_alignment.Alignment `json:"end,omitempty"`
}
