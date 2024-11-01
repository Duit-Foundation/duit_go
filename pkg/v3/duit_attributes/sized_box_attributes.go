package duit_attributes

type SizedBoxAttributes struct {
	ValueReferenceHolder
	Width  float32 `json:"width,omitempty"`
	Height float32 `json:"height,omitempty"`
}
