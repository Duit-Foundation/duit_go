package duit_props

type BoxShadow[T Color] struct {
	Color        T         `json:"color,omitempty"`
	BlurRadius   float32   `json:"blurRadius,omitempty"`
	SpreadRadius float32   `json:"spreadRadius,omitempty"`
	BlurStyle    BlurStyle `json:"blurStyle,omitempty"`
	Offset       Offset    `json:"offset,omitempty"`
}
