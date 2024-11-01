package duit_decoration

import (
	"github.com/Duit-Foundation/duit_go/pkg/duit_attributes/duit_alignment"
	"github.com/Duit-Foundation/duit_go/pkg/duit_attributes/duit_color"
)

type LinearGradient[T duit_color.Color] struct {
	Colors []T `json:"colors"`
	//Rotation angle in radians
	RotationAngle float32                  `json:"rotationAngle,omitempty"`
	Stops         []float32                `json:"stops,omitempty"`
	Begin         duit_alignment.Alignment `json:"begin,omitempty"`
	End           duit_alignment.Alignment `json:"end,omitempty"`
}
