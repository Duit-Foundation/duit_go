package duit_attributes

type ConstrainedBoxAttributes struct {
	ValueReferenceHolder
	MinWidth  float32 `json:"minWidth"`
	MaxWidth  float32 `json:"maxWidth"`
	MinHeight float32 `json:"minHeight"`
	MaxHeight float32 `json:"maxHeight"`
}
