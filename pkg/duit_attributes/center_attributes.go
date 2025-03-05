package duit_attributes

type CenterAttributes struct {
	ValueReferenceHolder
	ThemeConsumer
	WidgthFactor float32 `json:"widthFactor,omitempty"`
	HeightFactor float32 `json:"heightFactor,omitempty"`
}
