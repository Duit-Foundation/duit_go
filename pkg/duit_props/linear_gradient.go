package duit_props

type LinearGradient[T Color] struct {
	Colors []T `json:"colors"`
	//Rotation angle in radians
	RotationAngle float32   `json:"rotationAngle,omitempty"`
	Stops         []float32 `json:"stops,omitempty"`
	Begin         Alignment `json:"begin,omitempty"`
	End           Alignment `json:"end,omitempty"`
}
