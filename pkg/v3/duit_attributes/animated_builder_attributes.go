package duit_attributes

type AnimatedBuilderAttributes struct {
	ValueReferenceHolder
	Tweens       []any  `json:"tweens"`
	PersistentId string `json:"persistentId,omitempty"`
}
