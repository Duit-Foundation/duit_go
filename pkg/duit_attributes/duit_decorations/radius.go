package duit_decoration

type Radius struct {
	Radius float32 `json:"radius"`
}

type BorderRadius struct {
	TopLeft     *Radius `json:"topLeft,omitempty"`
	TopRight    *Radius `json:"topRight,omitempty"`
	BottomLeft  *Radius `json:"bottomLeft,omitempty"`
	BottomRight *Radius `json:"bottomRight,omitempty"`
}