package duit_decoration

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

type LinearGradient[T duit_props.Color] struct {
	Colors []T `json:"colors"`
	//Rotation angle in radians
	RotationAngle float32                  `json:"rotationAngle,omitempty"`
	Stops         []float32                `json:"stops,omitempty"`
	Begin         duit_props.Alignment `json:"begin,omitempty"`
	End           duit_props.Alignment `json:"end,omitempty"`
}
