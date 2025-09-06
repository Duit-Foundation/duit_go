package duit_props

type LinearGradient struct {
	Colors []Color `json:"colors"`
	//Rotation angle in radians
	RotationAngle float32   `json:"rotationAngle,omitempty"`
	Stops         []float32 `json:"stops,omitempty"`
	Begin         Alignment `json:"begin,omitempty"`
	End           Alignment `json:"end,omitempty"`
}
